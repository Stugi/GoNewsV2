package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	db *storage.DB
	r  *mux.Router
}

// Регистрация методов API в маршрутизаторе запросов.
func (api *API) endpoints() {
	// получить n последних новостей
	api.r.HandleFunc("/news/{n}", api.posts).Methods(http.MethodGet, http.MethodOptions)
	// веб-приложение
	api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}
