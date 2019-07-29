package routers

import (
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/shop/controllers"
)

func Setup(e *droxolite.Epoxy) {
	deftCtrl := &controllers.DefaultController{}
	deftGroup := droxolite.NewRouteGroup("", deftCtrl)
	deftGroup.AddRoute("/", "GET", roletype.Unknown, deftCtrl.Get)
}
