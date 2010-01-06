package main

import (
	"fmt"
	"gostache"
	"log"
	"redis"
	"web"
)

func redis_name(key string) string {
	spec := redis.DefaultSpec().Db(13); /* WTF 13 means? */
	client, e := redis.NewSynchClientWithSpec (spec)
	if e != nil { log.Stderr ("failed to create the client", e); return "failed" }
	// client.Set("foo", strings.Bytes("bar"))
	value, e := client.Get(key)
	if e!= nil { log.Stderr ("error on Get", e); return "failed 2" }
	return fmt.Sprintf("%s", value)
}

func hello(ctx *web.Context, val string) {
	filename := "templates/hello.mustache"
	world := map[string]string{"name": redis_name(val)}
	output, _ := gostache.RenderFile(filename, world)
	ctx.WriteString(output)
}

func main() {
	web.SetStaticDir("public")
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
}

