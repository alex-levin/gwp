package main

import (
	_ "github.com/alex-levin/mosaic-b/routers"
	"github.com/alex-levin/mosaic-b/mosaic"
	"github.com/astaxie/beego"
)

func main() {
	go mosaic.TilesDB()
	beego.Run()
}

