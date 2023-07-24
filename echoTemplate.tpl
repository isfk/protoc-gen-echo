{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

// type myHandler {
//     {{$svrType}}_EchoClientHandlerImpl
//	   log *slog.Logger	
// }
// 
// func NewMyHandler(log *slog.Logger) *myHandler {
// 	   return &myHandler{log: log}
// }
// 
// func main () {
//     e := echo.New()
//     log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
//     handler := example.NewExampleService_EchoServerHandler(NewMyHandler(log))
//     e.Get("/", e.handler.Hello)
//     if err := e.Start(":1111"); err != nil {
//         panic(err)
//     }
// }
// 
// func (h myHandler) Hello(args *example.HelloRequest) (*example.HelloResponse, error) {
//     h.log.Info("打印参数", slog.Any("args", args))
//     return &example.HelloResponse{Msg: args.Name}, nil
// }



type {{$svrType}}_EchoServerHandler interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(v4.Context) error
{{- end}}
}

type {{$svrType}}_EchoServerHandlerImpl struct {
	handler {{$svrType}}_EchoClientHandler
}

func New{{$svrType}}_EchoServerHandler(handler {{$svrType}}_EchoClientHandler) {{$svrType}}_EchoServerHandler {
	return &{{$svrType}}_EchoServerHandlerImpl{
		handler: handler,
	}
}

{{- range .Methods}}
{{- if ne .Comment ""}}
{{.Comment}}
{{- end}}
func (s {{$svrType}}_EchoServerHandlerImpl) {{.Name}}(c v4.Context) error {
	var args {{.Request}}
	if err := c.Bind(&args); err != nil {
		return err
	}

	resp, err := s.handler.{{.Name}}(&args)
	if err != nil {
		return err
	}
	
	return c.JSON(http.StatusOK, resp)
}
{{- end}}

type {{$svrType}}_EchoClientHandler interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(*{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

type {{$svrType}}_EchoClientHandlerImpl struct {
}

func New{{$svrType}}_EchoClientHandler() {{$svrType}}_EchoClientHandler {
	return &{{$svrType}}_EchoClientHandlerImpl{}
}

{{- range .Methods}}
{{- if ne .Comment ""}}
{{.Comment}} 方法需要自己实现
{{- end}}
func ({{$svrType}}_EchoClientHandlerImpl) {{.Name}}(args *{{.Request}}) (*{{.Reply}}, error) {
	return &{{.Reply}}{}, nil
}
{{- end}}