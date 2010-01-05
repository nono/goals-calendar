package main

import (
    "web"
)

func hello(val string) string { return "hello " + val }

func main() {
    web.SetStaticDir("public")
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}

