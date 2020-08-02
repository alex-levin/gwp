package main

import (
	_ "github.com/alex-levin/ws-b/docs"
	_ "github.com/alex-levin/ws-b/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
