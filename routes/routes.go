package routes

import (
	"github.com/gorilla/mux"
	"github.com/zhitongLIU/lentslog/routes/lents"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/lents", lents.LentsPostHandler).Methods("POST")

	return r
}
