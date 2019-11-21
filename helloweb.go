package main

import (
	"github.com/hoisie/web"
)

func main() {
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
}

//func hello(say string) string {
//	return "hello " + say
//}

func hello(ctx *web.Context, val string)  {
	for k,v := range ctx.Params {
		println(k, v)
	}
}
