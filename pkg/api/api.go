package api

import (
	"gonews/v2/pkg/storage"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	db *storage.Storage
	r  *mux.Router
}

func New(db *storage.Storage) *API {
	api := API{
		db: db,
		r:  mux.NewRouter(),
	}
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.r
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	// получить n последних новостей
	api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet, http.MethodOptions)
	// веб-приложение
	api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

func (api *API) posts(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// n := vars["n"]
	// count, err := strconv.Atoi(n)
	// if err != nil {
	// 	count = 10
	// }
	// _, err := api.db.GetPosts(1, count)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
}
