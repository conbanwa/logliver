## 项目启动

```cmd
go run main.go
```

如果出现端口冲突问题，可以在 main.go 中修改端口

```go
app.Run(iris.Addr("localhost:1234"), iris.WithoutServerError(iris.ErrServerClosed))
```

启动后访问[http://localhost:1234/](http://localhost:1454/index.html)
