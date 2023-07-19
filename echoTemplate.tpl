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