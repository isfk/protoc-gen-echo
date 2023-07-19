{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}


type {{.ServiceType}}EchoHandler interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(v4.Context) error
{{- end}}
}

{{- range .Methods}}
func {{$svrType}}_{{.Name}}Binder(c v4.Context) (*{{.Request}}, error) {
	var args *{{.Request}}
	if err := c.Bind(&args); err != nil {
		return &{{.Request}}{}, err
	}
	return args, nil
}
{{- end}}