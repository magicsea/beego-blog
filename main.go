package main

import (
	"github.com/astaxie/beego"
	"github.com/magicsea/beego-blog/g"
	_ "github.com/magicsea/beego-blog/routers"
)

func main() {
	g.InitEnv()
	beego.Run()
}
