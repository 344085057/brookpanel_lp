package main

import (
	_ "myBrookWeb/routers"
	_ "myBrookWeb/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()

}
