package nogo

import (
	"web"
)

// TODO what about POST/PUT/DELETE?
func Get(route string, handler interface{}) {
	web.Get(route, handler)
}

// TODO func Restful()

