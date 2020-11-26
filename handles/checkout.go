package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func Checkout(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Checkout", tmpl, "./views/checkout.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Checkout")
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println(err)
		}
	}
}