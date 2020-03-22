package main

import (
	_ "mbook/routers"
	_ "mbook/sysinit"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

