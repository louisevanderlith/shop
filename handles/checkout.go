package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func Checkout(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Checkout")
		err := mix.Write(w, fact.Create(r, "Checkout", "./views/checkout.html", nil))

		if err != nil {
			log.Println(err)
		}
	}
}
