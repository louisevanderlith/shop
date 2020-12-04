package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func Create(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//pge.ChangeTitle("Create")
		err := mix.Write(w, fact.Create(r, "Create", "./views/uploads.html", nil))

		if err != nil {
			log.Println(err)
		}
	}
}
