package main

import (
	"gostache"
	"nono/model"
	"web"
)

var m *model.Model

func hello(ctx *web.Context, val string) {
	filename := "templates/hello.mustache"
	// m.Set("toto", "foobar")
	world := map[string]string{"name": m.Get(val)}
	output, _ := gostache.RenderFile(filename, world)
	ctx.WriteString(output)
}

func main() {
	m = model.NewModel("goals-calendar", 13)
	web.SetStaticDir("public")
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
}

