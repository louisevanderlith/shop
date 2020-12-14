package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/open"
	folio "github.com/louisevanderlith/folio/api"
	"github.com/louisevanderlith/theme/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
)

var (
	AuthConfig *oauth2.Config
	credConfig *clientcredentials.Config
	Endpoints  map[string]string
)

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	AuthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess},
	}

	credConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "theme", "folio"},
	}

	err = api.UpdateTemplate(credConfig.Client(ctx), endpoints["theme"])

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

	lock := open.NewHybridLock(provider, credConfig, AuthConfig)

	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)
	r.HandleFunc("/logout", lock.Logout).Methods(http.MethodGet)
	r.HandleFunc("/refresh", lock.Refresh).Methods(http.MethodGet)

	fact := mix.NewPageFactory(tmpl)
	fact.AddMenu(FullMenu())
	fact.AddModifier(mix.EndpointMod(Endpoints))
	fact.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	fact.AddModifier(ThemeContentMod())

	r.Handle("/", lock.Protect(Index(fact))).Methods(http.MethodGet)
	r.Handle("/cart", lock.Protect(lock.Lock(Cart(fact)))).Methods(http.MethodGet)

	rcat := r.PathPrefix("/{category:[a-zA-Z]+}").Subrouter()
	rcat.Handle("", GetCategoryAds(fact)).Methods(http.MethodGet)
	rcat.Handle("/create", lock.Lock(Create(fact))).Methods(http.MethodGet)
	rcat.Handle("/{pagesize:[A-Z][0-9]+}", SearchAds(fact)).Methods(http.MethodGet)
	rcat.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchAds(fact)).Methods(http.MethodGet)
	rcat.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewAd(fact)).Methods(http.MethodGet)
	rcat.Use(lock.Protect)

	return r
}

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("b", "/cart", "Cart", nil))
	m.AddItem(menu.NewItem("e", "/clients", "Clients", nil))
	m.AddItem(menu.NewItem("a", "/orders", "Orders", nil))

	//TODO: Add categories as children
	m.AddItem(menu.NewItem("b", "/stock", "Sell", nil))
	m.AddItem(menu.NewItem("g", "/ads", "Buy", nil))

	return m
}

func ThemeContentMod() mix.ModFunc {
	return func(b mix.Bag, r *http.Request) {
		clnt := credConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		b.SetValue("Folio", content)
	}
}
