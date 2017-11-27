package pgghelpers

import (
	gogo "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	gogoplugin "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	golang "github.com/golang/protobuf/protoc-gen-go/descriptor"
	golangplugin "github.com/golang/protobuf/protoc-gen-go/plugin"
)

//
// Transform Go -> Gogo
//

// TransformCodeGeneratorRequestGo translates a code generator request
// from gogo plugin to grpc golang plugin types.
func TransformCodeGeneratorRequestGo(req *gogoplugin.CodeGeneratorRequest) *golangplugin.CodeGeneratorRequest {
	return &golangplugin.CodeGeneratorRequest{
		FileToGenerate:   req.FileToGenerate,
		Parameter:        req.Parameter,
		ProtoFile:        TransformFileDescriptorGo(req.ProtoFile),
		XXX_unrecognized: req.XXX_unrecognized,
	}
}

func TransformFileDescriptorGo(fds []*gogo.FileDescriptorProto) []*golang.FileDescriptorProto {
	results := []*golang.FileDescriptorProto{}
	for _, f := range fds {
		results = append(
			results,
			&golang.FileDescriptorProto{
				Name:             f.Name,
				Package:          f.Package,
				Dependency:       f.Dependency,
				PublicDependency: f.PublicDependency,
				WeakDependency:   f.WeakDependency,
				MessageType:      transformDescriptorGo(f.MessageType),
				EnumType:         transformEnumDescriptorGo(f.EnumType),
				Service:          transformServiceDescriptorGo(f.Service),
				Extension:        transformFieldGo(f.Extension),
				Options:          transformFileOptionsGo(f.Options),
				SourceCodeInfo:   transformSourceCodeInfoGo(f.SourceCodeInfo),
				Syntax:           f.Syntax,
				XXX_unrecognized: f.XXX_unrecognized,
			},
		)
	}
	return results
}

func transformServiceDescriptorGo(sds []*gogo.ServiceDescriptorProto) []*golang.ServiceDescriptorProto {
	results := []*golang.ServiceDescriptorProto{}
	for _, s := range sds {
		desc := &golang.ServiceDescriptorProto{
			Name:             s.Name,
			Method:           transformMethodDescriptorGo(s.Method),
			XXX_unrecognized: s.XXX_unrecognized,
		}
		if s.Options != nil {
			desc.Options = transformServiceOptionsGo(s.Options)
		}
		results = append(results, desc)
	}
	return results
}

func transformMethodDescriptorGo(mds []*gogo.MethodDescriptorProto) []*golang.MethodDescriptorProto {
	results := []*golang.MethodDescriptorProto{}
	for _, m := range mds {
		desc := &golang.MethodDescriptorProto{
			Name:             m.Name,
			InputType:        m.InputType,
			OutputType:       m.OutputType,
			ClientStreaming:  m.ClientStreaming,
			ServerStreaming:  m.ServerStreaming,
			XXX_unrecognized: m.XXX_unrecognized,
		}
		if m.Options != nil {
			desc.Options = transformMethodOptionsGo(m.Options)
		}
		results = append(results, desc)
	}
	return results
}

func transformMethodOptionsGo(opt *gogo.MethodOptions) *golang.MethodOptions {
	opts := &golang.MethodOptions{
		Deprecated:       opt.Deprecated,
		XXX_unrecognized: opt.XXX_unrecognized,
		// IdempotencyLevel: ?
		// proto.XXX_InternalExtensions: ?
	}
	if opt.UninterpretedOption != nil {
		opts.UninterpretedOption = transformUninterpretedOptionGo(opt.UninterpretedOption)
	}
	return opts
}

func transformServiceOptionsGo(opt *gogo.ServiceOptions) *golang.ServiceOptions {
	opts := &golang.ServiceOptions{
		Deprecated:       opt.Deprecated,
		XXX_unrecognized: opt.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if opt.UninterpretedOption != nil {
		opts.UninterpretedOption = transformUninterpretedOptionGo(opt.UninterpretedOption)
	}
	return opts
}

func transformFileOptionsGo(opt *gogo.FileOptions) *golang.FileOptions {
	opts := &golang.FileOptions{
		JavaPackage:               opt.JavaPackage,
		JavaOuterClassname:        opt.JavaOuterClassname,
		JavaMultipleFiles:         opt.JavaMultipleFiles,
		JavaGenerateEqualsAndHash: opt.JavaGenerateEqualsAndHash,
		JavaStringCheckUtf8:       opt.JavaStringCheckUtf8,
		GoPackage:                 opt.GoPackage,
		CcGenericServices:         opt.CcGenericServices,
		JavaGenericServices:       opt.JavaGenericServices,
		PyGenericServices:         opt.PyGenericServices,
		Deprecated:                opt.Deprecated,
		CcEnableArenas:            opt.CcEnableArenas,
		ObjcClassPrefix:           opt.ObjcClassPrefix,
		CsharpNamespace:           opt.CsharpNamespace,
		XXX_unrecognized:          opt.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if opt.OptimizeFor != nil {
		opts.OptimizeFor = golang.FileOptions_OptimizeMode(*opt.OptimizeFor).Enum()
	}
	if opt.UninterpretedOption != nil {
		opts.UninterpretedOption = transformUninterpretedOptionGo(opt.UninterpretedOption)
	}
	return opts
}

func transformSourceCodeInfoGo(s *gogo.SourceCodeInfo) *golang.SourceCodeInfo {
	return &golang.SourceCodeInfo{
		Location:         transformSourceCodeInfoLocationGo(s.Location),
		XXX_unrecognized: s.XXX_unrecognized,
	}
}

func transformSourceCodeInfoLocationGo(locs []*gogo.SourceCodeInfo_Location) []*golang.SourceCodeInfo_Location {
	results := []*golang.SourceCodeInfo_Location{}
	for _, l := range locs {
		results = append(
			results,
			&golang.SourceCodeInfo_Location{
				Path:                    l.Path,
				Span:                    l.Span,
				LeadingComments:         l.LeadingComments,
				TrailingComments:        l.TrailingComments,
				LeadingDetachedComments: l.LeadingDetachedComments,
				XXX_unrecognized:        l.XXX_unrecognized,
			},
		)
	}
	return results
}

func transformFieldGo(fields []*gogo.FieldDescriptorProto) []*golang.FieldDescriptorProto {
	results := []*golang.FieldDescriptorProto{}
	for _, f := range fields {
		fieldDesc := &golang.FieldDescriptorProto{
			Name:             f.Name,
			Number:           f.Number,
			TypeName:         f.TypeName,
			Extendee:         f.Extendee,
			DefaultValue:     f.DefaultValue,
			OneofIndex:       f.OneofIndex,
			JsonName:         f.JsonName,
			XXX_unrecognized: f.XXX_unrecognized,
		}
		if f.Options != nil {
			fieldDesc.Options = transformOptionsGo(f.Options)
		}
		if f.Label != nil {
			fieldDesc.Label = golang.FieldDescriptorProto_Label(*f.Label).Enum()
		}
		if f.Type != nil {
			fieldDesc.Type = golang.FieldDescriptorProto_Type(*f.Type).Enum()
		}
		results = append(results, fieldDesc)
	}
	return results
}

func transformOptionsGo(opts *gogo.FieldOptions) *golang.FieldOptions {
	fieldOpts := &golang.FieldOptions{
		Packed:           opts.Packed,
		Lazy:             opts.Lazy,
		Deprecated:       opts.Deprecated,
		Weak:             opts.Weak,
		XXX_unrecognized: opts.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if opts.Ctype != nil {
		fieldOpts.Ctype = golang.FieldOptions_CType(*opts.Ctype).Enum()
	}
	if opts.Jstype != nil {
		fieldOpts.Jstype = golang.FieldOptions_JSType(*opts.Jstype).Enum()
	}
	if opts.UninterpretedOption != nil {
		fieldOpts.UninterpretedOption = transformUninterpretedOptionGo(opts.UninterpretedOption)
	}
	return fieldOpts
}

func transformUninterpretedOptionGo(uOpts []*gogo.UninterpretedOption) []*golang.UninterpretedOption {
	results := []*golang.UninterpretedOption{}
	for _, uOpt := range uOpts {
		results = append(
			results,
			&golang.UninterpretedOption{
				Name:             transformUninterpretedOptionNamePartGo(uOpt.Name),
				IdentifierValue:  uOpt.IdentifierValue,
				PositiveIntValue: uOpt.PositiveIntValue,
				NegativeIntValue: uOpt.NegativeIntValue,
				DoubleValue:      uOpt.DoubleValue,
				StringValue:      uOpt.StringValue,
				AggregateValue:   uOpt.AggregateValue,
				XXX_unrecognized: uOpt.XXX_unrecognized,
			})
	}
	return results
}

func transformUninterpretedOptionNamePartGo(optNps []*gogo.UninterpretedOption_NamePart) []*golang.UninterpretedOption_NamePart {
	results := []*golang.UninterpretedOption_NamePart{}
	for _, optNp := range optNps {
		results = append(
			results,
			&golang.UninterpretedOption_NamePart{
				NamePart:         optNp.NamePart,
				IsExtension:      optNp.IsExtension,
				XXX_unrecognized: optNp.XXX_unrecognized,
			},
		)
	}
	return results
}

func transformDescriptorGo(descs []*gogo.DescriptorProto) []*golang.DescriptorProto {
	results := []*golang.DescriptorProto{}
	for _, d := range descs {
		desc := &golang.DescriptorProto{
			Name:             d.Name,
			Field:            transformFieldGo(d.Field),
			Extension:        transformFieldGo(d.Extension),
			NestedType:       transformDescriptorGo(d.NestedType),
			EnumType:         transformEnumDescriptorGo(d.EnumType),
			ExtensionRange:   transformExtensionRangeGo(d.ExtensionRange),
			OneofDecl:        transformOneofDescriptorGo(d.OneofDecl),
			ReservedRange:    transformReservedRangeGo(d.ReservedRange),
			ReservedName:     d.ReservedName,
			XXX_unrecognized: d.XXX_unrecognized,
		}
		if d.Options != nil {
			desc.Options = transformMessageOptionsGo(d.Options)
		}
		results = append(results, desc)
	}
	return results
}

func transformEnumDescriptorGo(enums []*gogo.EnumDescriptorProto) []*golang.EnumDescriptorProto {
	results := []*golang.EnumDescriptorProto{}
	for _, e := range enums {
		desc := &golang.EnumDescriptorProto{
			Name:             e.Name,
			Value:            transformEnumValueDescriptorGo(e.Value),
			XXX_unrecognized: e.XXX_unrecognized,
			// EnumReservedRange: ?
		}
		if e.Options != nil {
			desc.Options = transformEnumOptionsGo(e.Options)
		}
		results = append(results, desc)
	}
	return results
}

func transformEnumValueDescriptorGo(evds []*gogo.EnumValueDescriptorProto) []*golang.EnumValueDescriptorProto {
	results := []*golang.EnumValueDescriptorProto{}
	for _, e := range evds {
		desc := &golang.EnumValueDescriptorProto{
			Name:             e.Name,
			Number:           e.Number,
			XXX_unrecognized: e.XXX_unrecognized,
		}
		if e.Options != nil {
			desc.Options = transformEnumValueOptionsGo(e.Options)
		}
		results = append(results, desc)
	}
	return results
}

func transformEnumValueOptionsGo(e *gogo.EnumValueOptions) *golang.EnumValueOptions {
	opts := &golang.EnumValueOptions{
		Deprecated:       e.Deprecated,
		XXX_unrecognized: e.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if e.UninterpretedOption != nil {
		opts.UninterpretedOption = transformUninterpretedOptionGo(e.UninterpretedOption)
	}
	return opts
}

func transformEnumOptionsGo(e *gogo.EnumOptions) *golang.EnumOptions {
	opts := &golang.EnumOptions{
		AllowAlias:       e.AllowAlias,
		Deprecated:       e.Deprecated,
		XXX_unrecognized: e.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if e.UninterpretedOption != nil {
		opts.UninterpretedOption = transformUninterpretedOptionGo(e.UninterpretedOption)
	}
	return opts
}

func transformExtensionRangeGo(exrs []*gogo.DescriptorProto_ExtensionRange) []*golang.DescriptorProto_ExtensionRange {
	results := []*golang.DescriptorProto_ExtensionRange{}
	for _, e := range exrs {
		results = append(
			results,
			&golang.DescriptorProto_ExtensionRange{
				Start:            e.Start,
				End:              e.End,
				XXX_unrecognized: e.XXX_unrecognized,
				// ExtensionRangeOptions: ?
			},
		)
	}
	return results
}

func transformMessageOptionsGo(m *gogo.MessageOptions) *golang.MessageOptions {
	opts := &golang.MessageOptions{
		MessageSetWireFormat:         m.MessageSetWireFormat,
		NoStandardDescriptorAccessor: m.NoStandardDescriptorAccessor,
		Deprecated:                   m.Deprecated,
		MapEntry:                     m.MapEntry,
		XXX_unrecognized:             m.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if m.UninterpretedOption != nil {
		opts.UninterpretedOption = transformUninterpretedOptionGo(m.UninterpretedOption)
	}
	return opts
}

func transformOneofDescriptorGo(oods []*gogo.OneofDescriptorProto) []*golang.OneofDescriptorProto {
	results := []*golang.OneofDescriptorProto{}
	for _, o := range oods {
		desc := &golang.OneofDescriptorProto{
			Name:             o.Name,
			XXX_unrecognized: o.XXX_unrecognized,
		}
		if o.Options != nil {
			desc.Options = transformOneofOptionsGo(o.Options)
		}
		results = append(results, desc)
	}
	return results
}

func transformOneofOptionsGo(m *gogo.OneofOptions) *golang.OneofOptions {
	opts := &golang.OneofOptions{
		XXX_unrecognized: m.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if m.UninterpretedOption != nil {
		opts.UninterpretedOption = transformUninterpretedOptionGo(m.UninterpretedOption)
	}
	return opts
}

// ReservedRange

func transformReservedRangeGo(rrs []*gogo.DescriptorProto_ReservedRange) []*golang.DescriptorProto_ReservedRange {
	results := []*golang.DescriptorProto_ReservedRange{}
	for _, r := range rrs {
		results = append(
			results,
			&golang.DescriptorProto_ReservedRange{
				Start:            r.Start,
				End:              r.End,
				XXX_unrecognized: r.XXX_unrecognized,
			},
		)
	}
	return results
}

//
// Transform Gogo -> Go
//

func transformFieldGoGo(fields []*golang.FieldDescriptorProto) []*gogo.FieldDescriptorProto {
	results := []*gogo.FieldDescriptorProto{}
	for _, f := range fields {
		fieldDesc := &gogo.FieldDescriptorProto{
			Name:             f.Name,
			Number:           f.Number,
			TypeName:         f.TypeName,
			Extendee:         f.Extendee,
			DefaultValue:     f.DefaultValue,
			OneofIndex:       f.OneofIndex,
			JsonName:         f.JsonName,
			XXX_unrecognized: f.XXX_unrecognized,
		}
		if f.Options != nil {
			fieldDesc.Options = transformOptionsGoGo(f.Options)
		}
		if f.Label != nil {
			fieldDesc.Label = gogo.FieldDescriptorProto_Label(*f.Label).Enum()
		}
		if f.Type != nil {
			fieldDesc.Type = gogo.FieldDescriptorProto_Type(*f.Type).Enum()
		}
		results = append(results, fieldDesc)
	}
	return results
}

func transformOptionsGoGo(f *golang.FieldOptions) *gogo.FieldOptions {
	fieldOpts := &gogo.FieldOptions{
		Packed:           f.Packed,
		Lazy:             f.Lazy,
		Deprecated:       f.Deprecated,
		Weak:             f.Weak,
		XXX_unrecognized: f.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if f.Ctype != nil {
		fieldOpts.Ctype = gogo.FieldOptions_CType(*f.Ctype).Enum()
	}
	if f.Jstype != nil {
		fieldOpts.Jstype = gogo.FieldOptions_JSType(*f.Jstype).Enum()
	}
	if f.UninterpretedOption != nil {
		fieldOpts.UninterpretedOption = transformUninterpretedOptionGoGo(f.UninterpretedOption)
	}
	return fieldOpts
}

func transformUninterpretedOptionGoGo(uOpts []*golang.UninterpretedOption) []*gogo.UninterpretedOption {
	results := []*gogo.UninterpretedOption{}
	for _, uOpt := range uOpts {
		results = append(
			results,
			&gogo.UninterpretedOption{
				Name:             transformUninterpretedOptionNamePartGoGo(uOpt.Name),
				IdentifierValue:  uOpt.IdentifierValue,
				PositiveIntValue: uOpt.PositiveIntValue,
				NegativeIntValue: uOpt.NegativeIntValue,
				DoubleValue:      uOpt.DoubleValue,
				StringValue:      uOpt.StringValue,
				AggregateValue:   uOpt.AggregateValue,
				XXX_unrecognized: uOpt.XXX_unrecognized,
			})
	}
	return results
}

func transformUninterpretedOptionNamePartGoGo(optNps []*golang.UninterpretedOption_NamePart) []*gogo.UninterpretedOption_NamePart {
	results := []*gogo.UninterpretedOption_NamePart{}
	for _, optNp := range optNps {
		results = append(
			results,
			&gogo.UninterpretedOption_NamePart{
				NamePart:         optNp.NamePart,
				IsExtension:      optNp.IsExtension,
				XXX_unrecognized: optNp.XXX_unrecognized,
			},
		)
	}
	return results
}

func TransformCodeGeneratorResponseGoGo(resp *golangplugin.CodeGeneratorResponse_File) *gogoplugin.CodeGeneratorResponse_File {
	return &gogoplugin.CodeGeneratorResponse_File{
		Name:             resp.Name,
		InsertionPoint:   resp.InsertionPoint,
		Content:          resp.Content,
		XXX_unrecognized: resp.XXX_unrecognized,
	}
}
