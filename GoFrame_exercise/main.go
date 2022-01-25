package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)
import "github.com/gogf/gf"

func main() {
	fmt.Println("GoFrame version:", gf.VERSION)
	//create server
	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("Welcome GoFrame")
	})

	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World")
	})

	s.SetPort(80)
	s.Run()
}
