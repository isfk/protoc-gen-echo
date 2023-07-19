# isfk/protoc-gen-echo

## 安装

```sh
go install git.isfk.cn/isfk/protoc-gen-echo@latest
```

## 测试

```sh
go build .; protoc --plugin protoc-gen-echo --echo_out=. --echo_opt=paths=source_relative ./example/example.proto
```

## 调试

使用 `fmt.Fprintf`

```go
fmt.Fprintf(os.Stderr, "%v \n", method.Desc.Name())
fmt.Fprintf(os.Stderr, "%v \n", method.Desc.FullName())
```
