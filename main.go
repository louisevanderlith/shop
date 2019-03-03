package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"
	"github.com/louisevanderlith/shop/routers"
)

func main() {
	// Register with router
	mode := os.Getenv("RUNMODE")
	name := beego.BConfig.AppName
	srv := mango.NewService(mode, name, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		routers.Setup(srv)
		beego.Run()
	}
}
