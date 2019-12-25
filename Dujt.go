package main

import (
	"dujt/data"
	"errors"
	"flag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"log"
)

var (
	addr string
	port string
	target string
)

func main() {
	flag.StringVar(&addr, "addr", "0.0.0.0", "服务地址")
	flag.StringVar(&port, "port", "80", "服务端口")
	flag.StringVar(&target, "target", "statement.txt", "毒鸡汤文件路径")
	flag.Parse()

	if err := data.Init(target); err != nil {
		log.Panicln("毒鸡汤导入失败: " + err.Error())
		return
	}

	app := iris.New()
	app.RegisterView(iris.HTML("./", ".html"))
	app.StaticWeb("/img", "./img")
	app.Get("/", func(ctx context.Context) {
		ctx.Gzip(true)
		ctx.ContentType("text/html")

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
