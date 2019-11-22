package main

import (
	"dujt/data"
	"errors"
	"flag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var (
	addr string
	port string
)

func main() {

	flag.StringVar(&addr, "addr", "0.0.0.0", "服务地址")
	flag.StringVar(&port, "port", "80", "服务端口")
	flag.Parse()

	app := iris.New()

	app.RegisterView(iris.HTML("./web", ".html"))
	app.StaticWeb("/img", "./web/img")

	app.Get("/", func(ctx context.Context) {
		ctx.Gzip(true)
		ctx.ContentType("text/html")

		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("line", data.GetLine())
		ctx.ViewData("host", ctx.Request().Host)
		ctx.ViewData("url", "http://" + ctx.Request().Host)

		// 渲染模板文件： ./views/hello.html
		ctx.View("index.html")
	})

	err := app.Run(
		iris.Addr(addr+":"+port),
		iris.WithoutInterruptHandler,
	)
	if err != nil {
		if !errors.Is(err, iris.ErrServerClosed) {
			panic(err)
		}
	}
}
