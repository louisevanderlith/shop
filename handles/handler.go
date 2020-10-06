package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong/middle"
	"net/http"
)

func SetupRoutes(clnt, scrt, securityUrl, managerUrl, authorityUrl string) http.Handler {
	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	clntIns := middle.NewClientInspector(clnt, scrt, http.DefaultClient, securityUrl, managerUrl, authorityUrl)
	r.HandleFunc("/callback", clntIns.Callback).Queries("state", "{state}", "token", "{token}").Methods(http.MethodGet)

	r.HandleFunc("/", clntIns.Middleware(Index(tmpl), map[string]bool{"stock.clothing.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}", clntIns.Middleware(SearchAds(tmpl), map[string]bool{"stock.clothing.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", clntIns.Middleware(SearchAds(tmpl), map[string]bool{"stock.clothing.search": true})).Methods(http.MethodGet)
	r.HandleFunc("/{key:[0-9]+\\x60[0-9]+}", clntIns.Middleware(ViewAd(tmpl), map[string]bool{"stock.clothing.view": true})).Methods(http.MethodGet)

	return r
}
