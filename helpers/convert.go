package pgghelpers

import (
	desc_gogo "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	plugin_gogo "github.com/gogo/protobuf/protoc-gen-gogo/plugin"
	desc_go "github.com/golang/protobuf/protoc-gen-go/descriptor"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
)

//
// convert gogo -> go
//

// convertGoGoCodeGeneratorRequest translates a code generator request
// from gogo plugin to grpc go plugin types.
func ConvertGoGoCodeGeneratorRequest(req *plugin_gogo.CodeGeneratorRequest) *plugin_go.CodeGeneratorRequest {
	return &plugin_go.CodeGeneratorRequest{
		FileToGenerate:   req.FileToGenerate,
		Parameter:        req.Parameter,
		ProtoFile:        ConvertGoGoFileDescriptor(req.ProtoFile),
		XXX_unrecognized: req.XXX_unrecognized,
	}
}

func ConvertGoGoFileDescriptor(fds []*desc_gogo.FileDescriptorProto) []*desc_go.FileDescriptorProto {
	results := []*desc_go.FileDescriptorProto{}
	for _, f := range fds {
		results = append(
			results,
			&desc_go.FileDescriptorProto{
				Name:             f.Name,
				Package:          f.Package,
				Dependency:       f.Dependency,
				PublicDependency: f.PublicDependency,
				WeakDependency:   f.WeakDependency,
				MessageType:      convertGoGoDescriptor(f.MessageType),
				EnumType:         convertGoGoEnumDescriptor(f.EnumType),
				Service:          convertGoGoServiceDescriptor(f.Service),
				Extension:        convertGoGoField(f.Extension),
				Options:          convertGoGoFileOptions(f.Options),
				SourceCodeInfo:   convertGoGoSourceCodeInfo(f.SourceCodeInfo),
				Syntax:           f.Syntax,
				XXX_unrecognized: f.XXX_unrecognized,
			},
		)
	}
	return results
}

func convertGoGoServiceDescriptor(sds []*desc_gogo.ServiceDescriptorProto) []*desc_go.ServiceDescriptorProto {
	results := []*desc_go.ServiceDescriptorProto{}
	for _, s := range sds {
		desc := &desc_go.ServiceDescriptorProto{
			Name:             s.Name,
			Method:           convertGoGoMethodDescriptor(s.Method),
			XXX_unrecognized: s.XXX_unrecognized,
		}
		if s.Options != nil {
			desc.Options = convertGoGoServiceOptions(s.Options)
		}
		results = append(results, desc)
	}
	return results
}

func convertGoGoMethodDescriptor(mds []*desc_gogo.MethodDescriptorProto) []*desc_go.MethodDescriptorProto {
	results := []*desc_go.MethodDescriptorProto{}
	for _, m := range mds {
		desc := &desc_go.MethodDescriptorProto{
			Name:             m.Name,
			InputType:        m.InputType,
			OutputType:       m.OutputType,
			ClientStreaming:  m.ClientStreaming,
			ServerStreaming:  m.ServerStreaming,
			XXX_unrecognized: m.XXX_unrecognized,
		}
		if m.Options != nil {
			desc.Options = convertGoGoMethodOptions(m.Options)
		}
		results = append(results, desc)
	}
	return results
}

func convertGoGoMethodOptions(opt *desc_gogo.MethodOptions) *desc_go.MethodOptions {
	opts := &desc_go.MethodOptions{
		Deprecated:       opt.Deprecated,
		XXX_unrecognized: opt.XXX_unrecognized,
		// IdempotencyLevel: ?
		// proto.XXX_InternalExtensions: ?
	}
	if opt.UninterpretedOption != nil {
		opts.UninterpretedOption = convertGoGoUninterpretedOption(opt.UninterpretedOption)
	}
	return opts
}

func convertGoGoServiceOptions(opt *desc_gogo.ServiceOptions) *desc_go.ServiceOptions {
	opts := &desc_go.ServiceOptions{
		Deprecated:       opt.Deprecated,
		XXX_unrecognized: opt.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if opt.UninterpretedOption != nil {
		opts.UninterpretedOption = convertGoGoUninterpretedOption(opt.UninterpretedOption)
	}
	return opts
}

func convertGoGoFileOptions(opt *desc_gogo.FileOptions) *desc_go.FileOptions {
	opts := &desc_go.FileOptions{
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
		opts.OptimizeFor = desc_go.FileOptions_OptimizeMode(*opt.OptimizeFor).Enum()
	}
	if opt.UninterpretedOption != nil {
		opts.UninterpretedOption = convertGoGoUninterpretedOption(opt.UninterpretedOption)
	}
	return opts
}

func convertGoGoSourceCodeInfo(s *desc_gogo.SourceCodeInfo) *desc_go.SourceCodeInfo {
	return &desc_go.SourceCodeInfo{
		Location:         convertGoGoSourceCodeInfoLocation(s.Location),
		XXX_unrecognized: s.XXX_unrecognized,
	}
}

func convertGoGoSourceCodeInfoLocation(locs []*desc_gogo.SourceCodeInfo_Location) []*desc_go.SourceCodeInfo_Location {
	results := []*desc_go.SourceCodeInfo_Location{}
	for _, l := range locs {
		results = append(
			results,
			&desc_go.SourceCodeInfo_Location{
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

func convertGoGoField(fields []*desc_gogo.FieldDescriptorProto) []*desc_go.FieldDescriptorProto {
	results := []*desc_go.FieldDescriptorProto{}
	for _, f := range fields {
		fieldDesc := &desc_go.FieldDescriptorProto{
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
			fieldDesc.Options = convertGoGoOptions(f.Options)
		}
		if f.Label != nil {
			fieldDesc.Label = desc_go.FieldDescriptorProto_Label(*f.Label).Enum()
		}
		if f.Type != nil {
			fieldDesc.Type = desc_go.FieldDescriptorProto_Type(*f.Type).Enum()
		}
		results = append(results, fieldDesc)
	}
	return results
}

func convertGoGoOptions(opts *desc_gogo.FieldOptions) *desc_go.FieldOptions {
	fieldOpts := &desc_go.FieldOptions{
		Packed:           opts.Packed,
		Lazy:             opts.Lazy,
		Deprecated:       opts.Deprecated,
		Weak:             opts.Weak,
		XXX_unrecognized: opts.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if opts.Ctype != nil {
		fieldOpts.Ctype = desc_go.FieldOptions_CType(*opts.Ctype).Enum()
	}
	if opts.Jstype != nil {
		fieldOpts.Jstype = desc_go.FieldOptions_JSType(*opts.Jstype).Enum()
	}
	if opts.UninterpretedOption != nil {
		fieldOpts.UninterpretedOption = convertGoGoUninterpretedOption(opts.UninterpretedOption)
	}
	return fieldOpts
}

func convertGoGoUninterpretedOption(uOpts []*desc_gogo.UninterpretedOption) []*desc_go.UninterpretedOption {
	results := []*desc_go.UninterpretedOption{}
	for _, uOpt := range uOpts {
		results = append(
			results,
			&desc_go.UninterpretedOption{
				Name:             convertGoGoUninterpretedOptionNamePart(uOpt.Name),
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

func convertGoGoUninterpretedOptionNamePart(optNps []*desc_gogo.UninterpretedOption_NamePart) []*desc_go.UninterpretedOption_NamePart {
	results := []*desc_go.UninterpretedOption_NamePart{}
	for _, optNp := range optNps {
		results = append(
			results,
			&desc_go.UninterpretedOption_NamePart{
				NamePart:         optNp.NamePart,
				IsExtension:      optNp.IsExtension,
				XXX_unrecognized: optNp.XXX_unrecognized,
			},
		)
	}
	return results
}

func convertGoGoDescriptor(descs []*desc_gogo.DescriptorProto) []*desc_go.DescriptorProto {
	results := []*desc_go.DescriptorProto{}
	for _, d := range descs {
		desc := &desc_go.DescriptorProto{
			Name:             d.Name,
			Field:            convertGoGoField(d.Field),
			Extension:        convertGoGoField(d.Extension),
			NestedType:       convertGoGoDescriptor(d.NestedType),
			EnumType:         convertGoGoEnumDescriptor(d.EnumType),
			ExtensionRange:   convertGoGoExtensionRange(d.ExtensionRange),
			OneofDecl:        convertGoGoOneofDescriptor(d.OneofDecl),
			ReservedRange:    convertGoGoReservedRange(d.ReservedRange),
			ReservedName:     d.ReservedName,
			XXX_unrecognized: d.XXX_unrecognized,
		}
		if d.Options != nil {
			desc.Options = convertGoGoMessageOptions(d.Options)
		}
		results = append(results, desc)
	}
	return results
}

func convertGoGoEnumDescriptor(enums []*desc_gogo.EnumDescriptorProto) []*desc_go.EnumDescriptorProto {
	results := []*desc_go.EnumDescriptorProto{}
	for _, e := range enums {
		desc := &desc_go.EnumDescriptorProto{
			Name:             e.Name,
			Value:            convertGoGoEnumValueDescriptor(e.Value),
			XXX_unrecognized: e.XXX_unrecognized,
			// EnumReservedRange: ?
		}
		if e.Options != nil {
			desc.Options = convertGoGoEnumOptions(e.Options)
		}
		results = append(results, desc)
	}
	return results
}

func convertGoGoEnumValueDescriptor(evds []*desc_gogo.EnumValueDescriptorProto) []*desc_go.EnumValueDescriptorProto {
	results := []*desc_go.EnumValueDescriptorProto{}
	for _, e := range evds {
		desc := &desc_go.EnumValueDescriptorProto{
			Name:             e.Name,
			Number:           e.Number,
			XXX_unrecognized: e.XXX_unrecognized,
		}
		if e.Options != nil {
			desc.Options = convertGoGoEnumValueOptions(e.Options)
		}
		results = append(results, desc)
	}
	return results
}

func convertGoGoEnumValueOptions(e *desc_gogo.EnumValueOptions) *desc_go.EnumValueOptions {
	opts := &desc_go.EnumValueOptions{
		Deprecated:       e.Deprecated,
		XXX_unrecognized: e.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if e.UninterpretedOption != nil {
		opts.UninterpretedOption = convertGoGoUninterpretedOption(e.UninterpretedOption)
	}
	return opts
}

func convertGoGoEnumOptions(e *desc_gogo.EnumOptions) *desc_go.EnumOptions {
	opts := &desc_go.EnumOptions{
		AllowAlias:       e.AllowAlias,
		Deprecated:       e.Deprecated,
		XXX_unrecognized: e.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if e.UninterpretedOption != nil {
		opts.UninterpretedOption = convertGoGoUninterpretedOption(e.UninterpretedOption)
	}
	return opts
}

func convertGoGoExtensionRange(exrs []*desc_gogo.DescriptorProto_ExtensionRange) []*desc_go.DescriptorProto_ExtensionRange {
	results := []*desc_go.DescriptorProto_ExtensionRange{}
	for _, e := range exrs {
		results = append(
			results,
			&desc_go.DescriptorProto_ExtensionRange{
				Start:            e.Start,
				End:              e.End,
				XXX_unrecognized: e.XXX_unrecognized,
				// ExtensionRangeOptions: ?
			},
		)
	}
	return results
}

func convertGoGoMessageOptions(m *desc_gogo.MessageOptions) *desc_go.MessageOptions {
	opts := &desc_go.MessageOptions{
		MessageSetWireFormat:         m.MessageSetWireFormat,
		NoStandardDescriptorAccessor: m.NoStandardDescriptorAccessor,
		Deprecated:                   m.Deprecated,
		MapEntry:                     m.MapEntry,
		XXX_unrecognized:             m.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if m.UninterpretedOption != nil {
		opts.UninterpretedOption = convertGoGoUninterpretedOption(m.UninterpretedOption)
	}
	return opts
}

func convertGoGoOneofDescriptor(oods []*desc_gogo.OneofDescriptorProto) []*desc_go.OneofDescriptorProto {
	results := []*desc_go.OneofDescriptorProto{}
	for _, o := range oods {
		desc := &desc_go.OneofDescriptorProto{
			Name:             o.Name,
			XXX_unrecognized: o.XXX_unrecognized,
		}
		if o.Options != nil {
			desc.Options = convertGoGoOneofOptions(o.Options)
		}
		results = append(results, desc)
	}
	return results
}

func convertGoGoOneofOptions(m *desc_gogo.OneofOptions) *desc_go.OneofOptions {
	opts := &desc_go.OneofOptions{
		XXX_unrecognized: m.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if m.UninterpretedOption != nil {
		opts.UninterpretedOption = convertGoGoUninterpretedOption(m.UninterpretedOption)
	}
	return opts
}

// ReservedRange

func convertGoGoReservedRange(rrs []*desc_gogo.DescriptorProto_ReservedRange) []*desc_go.DescriptorProto_ReservedRange {
	results := []*desc_go.DescriptorProto_ReservedRange{}
	for _, r := range rrs {
		results = append(
			results,
			&desc_go.DescriptorProto_ReservedRange{
				Start:            r.Start,
				End:              r.End,
				XXX_unrecognized: r.XXX_unrecognized,
			},
		)
	}
	return results
}

//
// convert go -> gogo
//

func convertGoField(fields []*desc_go.FieldDescriptorProto) []*desc_gogo.FieldDescriptorProto {
	results := []*desc_gogo.FieldDescriptorProto{}
	for _, f := range fields {
		fieldDesc := &desc_gogo.FieldDescriptorProto{
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
			fieldDesc.Options = convertGoOptions(f.Options)
		}
		if f.Label != nil {
			fieldDesc.Label = desc_gogo.FieldDescriptorProto_Label(*f.Label).Enum()
		}
		if f.Type != nil {
			fieldDesc.Type = desc_gogo.FieldDescriptorProto_Type(*f.Type).Enum()
		}
		results = append(results, fieldDesc)
	}
	return results
}

func convertGoOptions(f *desc_go.FieldOptions) *desc_gogo.FieldOptions {
	fieldOpts := &desc_gogo.FieldOptions{
		Packed:           f.Packed,
		Lazy:             f.Lazy,
		Deprecated:       f.Deprecated,
		Weak:             f.Weak,
		XXX_unrecognized: f.XXX_unrecognized,
		// proto.XXX_InternalExtensions: ?
	}
	if f.Ctype != nil {
		fieldOpts.Ctype = desc_gogo.FieldOptions_CType(*f.Ctype).Enum()
	}
	if f.Jstype != nil {
		fieldOpts.Jstype = desc_gogo.FieldOptions_JSType(*f.Jstype).Enum()
	}
	if f.UninterpretedOption != nil {
		fieldOpts.UninterpretedOption = convertGoUninterpretedOption(f.UninterpretedOption)
	}
	return fieldOpts
}

func convertGoUninterpretedOption(uOpts []*desc_go.UninterpretedOption) []*desc_gogo.UninterpretedOption {
	results := []*desc_gogo.UninterpretedOption{}
	for _, uOpt := range uOpts {
		results = append(
			results,
			&desc_gogo.UninterpretedOption{
				Name:             convertGoUninterpretedOptionNamePart(uOpt.Name),
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

func convertGoUninterpretedOptionNamePart(optNps []*desc_go.UninterpretedOption_NamePart) []*desc_gogo.UninterpretedOption_NamePart {
	results := []*desc_gogo.UninterpretedOption_NamePart{}
	for _, optNp := range optNps {
		results = append(
			results,
			&desc_gogo.UninterpretedOption_NamePart{
				NamePart:         optNp.NamePart,
				IsExtension:      optNp.IsExtension,
				XXX_unrecognized: optNp.XXX_unrecognized,
			},
		)
	}
	return results
}

func ConvertGoCodeGeneratorResponse(resp *plugin_go.CodeGeneratorResponse_File) *plugin_gogo.CodeGeneratorResponse_File {
	return &plugin_gogo.CodeGeneratorResponse_File{
		Name:             resp.Name,
		InsertionPoint:   resp.InsertionPoint,
		Content:          resp.Content,
		XXX_unrecognized: resp.XXX_unrecognized,
	}
}
