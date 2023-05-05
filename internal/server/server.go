package server

import (
	"os"
	"time"

	"github.com/bradford-hamilton/bt-data-server/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type API struct {
	baseURL string
	db      storage.SQLDatabase
	Mux     *chi.Mux
}

func New(db storage.SQLDatabase) *API {
	r := chi.NewRouter()
	r.Use(
		corsMiddleware().Handler,
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.StripSlashes,            // strip slashes to no slash URL versions
		middleware.Recoverer,               // recover from panics without crashing server
		middleware.Timeout(30*time.Second), // start with a pretty standard timeout
	)

	baseURL := "http://localhost:4000"
	if os.Getenv("BT_DATA_SERVER_ENVIRONMENT") == "production" {
		baseURL = "TODO amazon load balancer URL when ready"
	}

	api := &API{db: db, Mux: r, baseURL: baseURL}
	api.initializeRoutes()

	return api
}

func (a *API) initializeRoutes() {
	a.Mux.Get("/ping", a.ping)
	a.Mux.Post("/dd/new", a.createDataDump)
}
