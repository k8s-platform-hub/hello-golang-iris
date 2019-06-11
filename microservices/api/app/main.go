package main

import (
	"github.com/kataras/iris"
)

// Hello response structure
type Hello struct {
	Message string
}

func main() {
	app := iris.Default()
	app.Get("/", index)
	app.Run(iris.Addr(":8080"))
}

func index(ctx iris.Context) {
	m := Hello{"Welcome to Hasura"}
	ctx.JSON(m)
}
