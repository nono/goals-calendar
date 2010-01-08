package nogo

import (
	"web"
)

func Initialize() {
	ReadConfig("config.json")
}

func Start() {
	addr := GetConfig("interface") + ":" + GetConfig("port")
	web.SetStaticDir(GetConfig("docroot"))
	web.Run(addr)
}

