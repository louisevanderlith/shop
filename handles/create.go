package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func Create(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Create", tmpl, "./views/uploads.html")

	return func(w http.ResponseWriter, r *http.Request) {
		pge.ChangeTitle("Create")
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println(err)
		}
	}
}