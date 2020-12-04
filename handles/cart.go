package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func Cart(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Cart")
		err := mix.Write(w, fact.Create(r, "Cart", "./views/cart.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
