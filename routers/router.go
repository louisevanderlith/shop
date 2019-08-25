package routers

import (
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
	"github.com/louisevanderlith/shop/controllers"
)

func Setup(e resins.Epoxi) {
	servGroup := routing.NewInterfaceBundle("Home", roletype.Unknown, &controllers.Home{})
	e.AddGroup(servGroup)
}
