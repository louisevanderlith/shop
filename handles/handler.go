package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/open"
	"github.com/louisevanderlith/theme/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
)

var (
	CredConfig *clientcredentials.Config
	Endpoints map[string]string
	//StockURL   string
)

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	authConfig := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID},
	}

	CredConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "theme"},
	}

	err = api.UpdateTemplate(CredConfig.Client(ctx), endpoints["theme"])

	if err != nil {
		panic(err)
	}

	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	lock := open.NewUILock(authConfig)
	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)

	oidcConfig := &oidc.Config{
		ClientID: clientId,
	}
	v := provider.Verifier(oidcConfig)

	r.HandleFunc("/", open.LoginMiddleware(v, Index(tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}", open.LoginMiddleware(v, SearchAds(tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", open.LoginMiddleware(v, SearchAds(tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/{key:[0-9]+\\x60[0-9]+}", open.LoginMiddleware(v, ViewAd(tmpl))).Methods(http.MethodGet)

	r.HandleFunc("/cart", clntIns.Middleware(Cart(tmpl), map[string]bool{"stock.clothing.search": true})).Methods(http.MethodGet)

	r.HandleFunc("/create", clntIns.Middleware(Create(tmpl), map[string]bool{"stock.clothing.search": true})).Methods(http.MethodGet)

	return r
}
