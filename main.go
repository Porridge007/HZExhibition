package main

import (
	_ "HZExhibition/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.BConfig.WebConfig.StaticDir["/static/img"] = "img"
	beego.Run()
}

