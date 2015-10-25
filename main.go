package main

import (
	"github.com/astaxie/beego"
	_ "github.com/everfore/upload/routers"
)

func main() {
	beego.Run()
}
