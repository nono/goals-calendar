package main

import (
	"gostache"
	"nogo"
	"web"
)

var m *nogo.Model

func hello(ctx *web.Context, val string) {
	filename := "app/templates/hello.mustache"
	// m.Set("toto", "foobar")
	world := map[string]string{"name": m.Get(val)}
	output, _ := gostache.RenderFile(filename, world)
	ctx.WriteString(output)
}

func main() {
	nogo.Initialize()
	nogo.Get("/(.*)", hello)
	nogo.Start()
	m = nogo.NewModel("calendar")
}

