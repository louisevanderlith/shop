package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func GetCredits(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Credits", "./views/xchange/credits.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchCredits(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := mix.Write(w, fact.Create(r, "Credits", "./views/xchange/credits.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewCredits(fact mix.MixerFactory) http.HandlerFunc {
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

		err := mix.Write(w, fact.Create(r, "CreditsView", "./views/xchange/creditview.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
