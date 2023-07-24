# isfk/protoc-gen-echo

## 安装

```sh
go install github.com/isfk/protoc-gen-echo@latest
```

## 生成

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --echo_out=. --echo_opt=paths=source_relative \
    demo/demo.proto

# buf generate example; buf generate example --template=buf.gen.tag.yaml
```

## 测试代码

```go
package main

import (
	"os"

	"github.com/isfk/protoc-gen-echo/example"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

type myHandler struct {
	example.ExampleService_EchoClientHandlerImpl
	log *slog.Logger
}

func NewMyHandler(log *slog.Logger) *myHandler {
	return &myHandler{log: log}
}

func main() {
	e := echo.New()

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	handler := example.NewExampleService_EchoServerHandler(NewMyHandler(log))

	e.GET("/hello", handler.Hello)
	e.GET("/say", handler.Say)
	e.Start(":1111")
}

func (h myHandler) Hello(args *example.HelloRequest) (*example.HelloResponse, error) {
	h.log.Info("打印参数", slog.Any("args", args))
	return &example.HelloResponse{Msg: args.Name}, nil
}

func (h myHandler) Say(args *example.SayRequest) (*example.SayResponse, error) {
	h.log.Info("打印参数", slog.Any("args", args))
	return &example.SayResponse{Msg: args.Name}, nil
}
```

https://github.com/isfk/protoc-gen-echo-test

## 调试

使用 `fmt.Fprintf`

```go
fmt.Fprintf(os.Stderr, "%v \n", method.Desc.Name())
fmt.Fprintf(os.Stderr, "%v \n", method.Desc.FullName())
```
