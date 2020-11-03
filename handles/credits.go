package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func GetCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Credits", tmpl, "./views/xchange/credits.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Credits", tmpl, "./views/xchange/credits.html")
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewCredits(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("CreditsView", tmpl, "./views/xchange/creditview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		//key, err := keys.ParseKey(drx.FindParam(r, "key"))

		//if err != nil {
		//	log.Println("Parse Key Error", err)
		//	http.Error(w, "", http.StatusBadRequest)
		//	return
		//}

		//result, err := api.FetchCredits(key)

		//if err != nil {
		//	log.Println("Fetch Credit Error", err)
		//	http.Error(w, "", http.StatusUnauthorized)
		//	return
		//}

		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
