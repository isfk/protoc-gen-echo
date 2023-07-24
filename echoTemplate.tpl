{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

// type myHandler {
//     {{$svrType}}_EchoClientHandlerImpl
//	   Log *slog.Logger	
// }
// 
// func NewMyHandler(log *slog.Logger) *myHandler {
// 	   return &myHandler{Log: log}
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
//     h.Log.Info("打印参数", slog.Any("args", args))
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
	Handler {{$svrType}}_EchoClientHandler
}

func New{{$svrType}}_EchoServerHandler(handler {{$svrType}}_EchoClientHandler) {{$svrType}}_EchoServerHandler {
	return &{{$svrType}}_EchoServerHandlerImpl{
		Handler: handler,
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

	resp, err := s.Handler.{{.Name}}(&args)
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

// 下面方法仅供参考, 具体需要自己实现

type {{$svrType}}_EchoClientHandlerImpl struct {
}

func New{{$svrType}}_EchoClientHandler() {{$svrType}}_EchoClientHandler {
	return &{{$svrType}}_EchoClientHandlerImpl{}
}

{{- range .Methods}}
{{- if ne .Comment ""}}
{{.Comment}}
{{- end}}
func ({{$svrType}}_EchoClientHandlerImpl) {{.Name}}(args *{{.Request}}) (*{{.Reply}}, error) {
	return &{{.Reply}}{}, nil
}
{{- end}}