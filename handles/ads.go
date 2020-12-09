package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"github.com/louisevanderlith/stock/api"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GetAds(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchAllCategories(clnt, Endpoints["stock"], "A10")

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

		err = mix.Write(w, fact.Create(r, "Ads", "./views/results.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println(err)
		}
	}
}

func GetCategoryAds(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := drx.FindParam(r, "category")
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchCategoryItems(clnt, Endpoints["stock"], category, "A10")

		if err != nil {
			log.Println("Fetch Clothing Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Ads", "./views/results.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchAds(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := drx.FindParam(r, "category")
		pagesize := drx.FindParam(r, "pagesize")
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchCategoryItems(clnt, Endpoints["stock"], category, pagesize)

		if err != nil {
			log.Println("Fetch Items Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		err = mix.Write(w, fact.Create(r, "Ads", "./views/results.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewAd(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		category := drx.FindParam(r, "category")
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchStockItem(clnt, Endpoints["stock"], category, key)

		if err != nil {
			log.Println("Fetch Item Error", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		err = mix.Write(w, fact.Create(r, "Ad View", "./views/adview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
