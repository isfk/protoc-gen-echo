package main

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	httpPackage        = protogen.GoImportPath("net/http")
	echoPackage        = protogen.GoImportPath("github.com/labstack/echo/v4")
	deprecationComment = "// Deprecated: Do not use."
)

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_echo.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-echo. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-echo ", version)
	g.P("// - protoc          ", getProtocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	g.P("var _ = new(", echoPackage.Ident("Context"), ")")
	g.P("var _ = new(", httpPackage.Ident("Client"), ")")

	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}

	sd := &serviceDesc{
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		Metadata:    file.Desc.Path(),
	}

	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		comment := method.Comments.Leading.String() + method.Comments.Trailing.String()
		if comment != "" {
			comment = "// " + method.GoName + strings.TrimPrefix(strings.TrimSuffix(comment, "\n"), "//")
		}

		comments := []string{}
		for _, v := range method.Comments.LeadingDetached {
			comments = append(comments, v.String())
		}

		sd.Methods = append(sd.Methods, &methodDesc{
			Name:         method.GoName,
			OriginalName: string(method.Desc.Name()),
			Num:          0,
			Request:      g.QualifiedGoIdent(method.Input.GoIdent),
			Reply:        g.QualifiedGoIdent(method.Output.GoIdent),
			Comment:      comment,
			Path:         "",
			Method:       "",
			HasVars:      false,
			HasBody:      false,
			Body:         "",
			ResponseBody: "",
			Swag:         strings.TrimSuffix(strings.Join(comments, "\n"), "\n"),
		})
	}

	if len(service.Methods) > 0 {
		g.P(sd.execute())
	}
}

func getProtocVersion(gen *protogen.Plugin) string {
	return fmt.Sprintf("%v.%v.%v", gen.Request.GetCompilerVersion().GetMajor(), gen.Request.GetCompilerVersion().GetMinor(), gen.Request.GetCompilerVersion().GetPatch())
}
