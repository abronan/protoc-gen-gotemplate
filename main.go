package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	plugin_gogo "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
	ggdescriptor "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"

	pgghelpers "github.com/moul/protoc-gen-gotemplate/helpers"
)

var (
	registry *ggdescriptor.Registry // some helpers need access to registry
)

func main() {
	g := generator.New()

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	files := g.Request.GetProtoFile()

	for _, opt := range []func(*descriptor.FileDescriptorProto){
		// We do not turn off getters for the generator.
		// vanity.TurnOffGoGettersAll,
		vanity.TurnOffGoStringerAll,
		vanity.TurnOnMarshalerAll,
		vanity.TurnOnStringerAll,
		vanity.TurnOnUnmarshalerAll,
		vanity.TurnOnSizerAll,
		CustomNameID,
	} {
		vanity.ForEachFile(files, opt)
	}

	// Parse parameters
	var (
		templateDir       = "./templates"
		destinationDir    = "."
		debug             = false
		all               = false
		singlePackageMode = false
	)
	if parameter := g.Request.GetParameter(); parameter != "" {
		for _, param := range strings.Split(parameter, ",") {
			parts := strings.Split(param, "=")
			if len(parts) != 2 {
				log.Printf("Err: invalid parameter: %q", param)
				continue
			}
			switch parts[0] {
			case "template_dir":
				templateDir = parts[1]
				break
			case "destination_dir":
				destinationDir = parts[1]
				break
			case "single-package-mode":
				switch strings.ToLower(parts[1]) {
				case "true", "t":
					singlePackageMode = true
				case "false", "f":
				default:
					log.Printf("Err: invalid value for single-package-mode: %q", parts[1])
				}
				break
			case "debug":
				switch strings.ToLower(parts[1]) {
				case "true", "t":
					debug = true
				case "false", "f":
				default:
					log.Printf("Err: invalid value for debug: %q", parts[1])
				}
				break
			case "all":
				switch strings.ToLower(parts[1]) {
				case "true", "t":
					all = true
				case "false", "f":
				default:
					log.Printf("Err: invalid value for debug: %q", parts[1])
				}
				break
			default:
				log.Printf("Err: unknown parameter: %q", param)
			}
		}
	}

	tmplMap := make(map[string]*plugin_gogo.CodeGeneratorResponse_File)
	concatOrAppend := func(file *plugin_gogo.CodeGeneratorResponse_File) {
		if val, ok := tmplMap[file.GetName()]; ok {
			*val.Content += file.GetContent()
		} else {
			tmplMap[file.GetName()] = file
			g.Response.File = append(g.Response.File, file)
		}
	}

	if singlePackageMode {
		registry = ggdescriptor.NewRegistry()
		pgghelpers.SetRegistry(registry)
		goReq := pgghelpers.ConvertGoGoCodeGeneratorRequest(g.Request)
		if err := registry.Load(goReq); err != nil {
			g.Error(err, "registry: failed to load the request")
		}
	}

	// Convert gogo file descriptor to golang descriptor
	gofiles := pgghelpers.ConvertGoGoFileDescriptor(files)

	// Generate the encoders
	for _, file := range gofiles {
		if all {
			if singlePackageMode {
				if _, err := registry.LookupFile(file.GetName()); err != nil {
					g.Error(err, "registry: failed to lookup file %q", file.GetName())
				}
			}
			encoder := NewGenericTemplateBasedEncoder(templateDir, file, debug, destinationDir)
			for _, tmpl := range encoder.Files() {
				concatOrAppend(tmpl)
			}

			continue
		}

		for _, service := range file.GetService() {
			encoder := NewGenericServiceTemplateBasedEncoder(templateDir, service, file, debug, destinationDir)
			for _, tmpl := range encoder.Files() {
				concatOrAppend(tmpl)
			}
		}
	}

	// Generate the protobufs
	resp := command.Generate(g.Request)
	command.Write(resp)
}
