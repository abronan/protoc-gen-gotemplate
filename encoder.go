package main

import (
	"bytes"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	gogoplugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	godesc "github.com/golang/protobuf/protoc-gen-go/descriptor"

	pgghelpers "github.com/abronan/protoc-gen-gotemplate/helpers"
)

type GenericTemplateBasedEncoder struct {
	templateDir    string
	service        *godesc.ServiceDescriptorProto
	file           *godesc.FileDescriptorProto
	enum           []*godesc.EnumDescriptorProto
	debug          bool
	destinationDir string
}

type Ast struct {
	BuildDate      time.Time                      `json:"build-date"`
	BuildHostname  string                         `json:"build-hostname"`
	BuildUser      string                         `json:"build-user"`
	GoPWD          string                         `json:"go-pwd,omitempty"`
	PWD            string                         `json:"pwd"`
	Debug          bool                           `json:"debug"`
	DestinationDir string                         `json:"destination-dir"`
	File           *godesc.FileDescriptorProto    `json:"file"`
	RawFilename    string                         `json:"raw-filename"`
	Filename       string                         `json:"filename"`
	TemplateDir    string                         `json:"template-dir"`
	Service        *godesc.ServiceDescriptorProto `json:"service"`
	Enum           []*godesc.EnumDescriptorProto  `json:"enum"`
}

func NewGenericServiceTemplateBasedEncoder(templateDir string, service *godesc.ServiceDescriptorProto, file *godesc.FileDescriptorProto, debug bool, destinationDir string) (e *GenericTemplateBasedEncoder) {
	e = &GenericTemplateBasedEncoder{
		service:        service,
		file:           file,
		templateDir:    templateDir,
		debug:          debug,
		destinationDir: destinationDir,
		enum:           file.GetEnumType(),
	}
	if debug {
		log.Printf("new encoder: file=%q service=%q template-dir=%q", file.GetName(), service.GetName(), templateDir)
	}

	return
}

func NewGenericTemplateBasedEncoder(templateDir string, file *godesc.FileDescriptorProto, debug bool, destinationDir string) (e *GenericTemplateBasedEncoder) {
	e = &GenericTemplateBasedEncoder{
		service:        nil,
		file:           file,
		templateDir:    templateDir,
		enum:           file.GetEnumType(),
		debug:          debug,
		destinationDir: destinationDir,
	}
	if debug {
		log.Printf("new encoder: file=%q template-dir=%q", file.GetName(), templateDir)
	}

	return
}

func (e *GenericTemplateBasedEncoder) templates() ([]string, error) {
	filenames := []string{}

	err := filepath.Walk(e.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".tmpl" {
			return nil
		}
		rel, err := filepath.Rel(e.templateDir, path)
		if err != nil {
			return err
		}
		if e.debug {
			log.Printf("new template: %q", rel)
		}
		filenames = append(filenames, rel)
		return nil
	})
	return filenames, err
}

func (e *GenericTemplateBasedEncoder) genAst(templateFilename string) (*Ast, error) {
	// prepare the ast passed to the template engine
	hostname, _ := os.Hostname()
	pwd, _ := os.Getwd()
	goPwd := ""
	if os.Getenv("GOPATH") != "" {
		goPwd, _ = filepath.Rel(os.Getenv("GOPATH")+"/src", pwd)
		if strings.Contains(goPwd, "../") {
			goPwd = ""
		}
	}
	ast := Ast{
		BuildDate:      time.Now(),
		BuildHostname:  hostname,
		BuildUser:      os.Getenv("USER"),
		PWD:            pwd,
		GoPWD:          goPwd,
		File:           e.file,
		TemplateDir:    e.templateDir,
		DestinationDir: e.destinationDir,
		RawFilename:    templateFilename,
		Filename:       "",
		Service:        e.service,
		Enum:           e.enum,
	}
	buffer := new(bytes.Buffer)
	tmpl, err := template.New("").Funcs(pgghelpers.ProtoHelpersFuncMap).Parse(templateFilename)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buffer, ast); err != nil {
		return nil, err
	}
	ast.Filename = buffer.String()
	return &ast, nil
}

func (e *GenericTemplateBasedEncoder) buildContent(templateFilename string) (string, string, error) {
	// initialize template engine
	fullPath := filepath.Join(e.templateDir, templateFilename)
	templateName := filepath.Base(fullPath)
	tmpl, err := template.New(templateName).Funcs(pgghelpers.ProtoHelpersFuncMap).ParseFiles(fullPath)
	if err != nil {
		return "", "", err
	}

	ast, err := e.genAst(templateFilename)
	if err != nil {
		return "", "", err
	}

	// generate the content
	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, ast); err != nil {
		return "", "", err
	}

	return buffer.String(), ast.Filename, nil
}

func (e *GenericTemplateBasedEncoder) Files() []*gogoplugin.CodeGeneratorResponse_File {
	templates, err := e.templates()
	if err != nil {
		log.Fatalf("cannot get templates from %q: %v", e.templateDir, err)
	}

	length := len(templates)
	files := make([]*gogoplugin.CodeGeneratorResponse_File, 0, length)
	errChan := make(chan error, length)
	resultChan := make(chan *gogoplugin.CodeGeneratorResponse_File, length)
	for _, templateFilename := range templates {
		go func(tmpl string) {
			content, translatedFilename, err := e.buildContent(tmpl)
			if err != nil {
				errChan <- err
				return
			}
			filename := translatedFilename[:len(translatedFilename)-len(".tmpl")]

			resultChan <- &gogoplugin.CodeGeneratorResponse_File{
				Content: &content,
				Name:    &filename,
			}
		}(templateFilename)
	}
	for i := 0; i < length; i++ {
		select {
		case f := <-resultChan:
			files = append(files, f)
		case err = <-errChan:
		}
	}
	if err != nil {
		panic(err)
	}
	return files
}
