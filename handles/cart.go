package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func Cart(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Cart", tmpl, "./views/cart.html")

	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Cart")
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println(err)
		}
	}
}