package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func GetSubmissions(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, fact.Create(r, "Submissions", "./views/submissions.html", nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
