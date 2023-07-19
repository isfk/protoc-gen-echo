{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}


type {{.ServiceType}}EchoServer interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(v4.Context) error
{{- end}}
}