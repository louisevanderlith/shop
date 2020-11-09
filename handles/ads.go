package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/stock/api"
	"html/template"
	"log"
	"net/http"
)

func GetAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/results.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllCategories(clnt, Endpoints["stock"], "A10")

		if err != nil {
			log.Println("Fetch Clothing Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println(err)
		}
	}
}

func GetCategoryAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/results.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		category := drx.FindParam(r, "category")
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchCategoryItems(clnt, Endpoints["stock"], category, "A10")

		if err != nil {
			log.Println("Fetch Clothing Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchAds(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/results.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		category := drx.FindParam(r, "category")
		pagesize := drx.FindParam(r, "pagesize")
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchCategoryItems(clnt, Endpoints["stock"], category, pagesize)

		if err != nil {
			log.Println("Fetch Items Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewAd(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Ads", tmpl, "./views/adview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		category := drx.FindParam(r, "category")
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchStockItem(clnt, Endpoints["stock"], category, key)

		if err != nil {
			log.Println("Fetch Item Error", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
