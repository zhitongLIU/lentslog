package routes

import (
	"github.com/gorilla/mux"
	"github.com/zhitongLIU/lentslog/middleware"
	"github.com/zhitongLIU/lentslog/routes/borrow"
	"github.com/zhitongLIU/lentslog/routes/lents"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/lents", middleware.CertifyApp(lents.PostHandler)).Methods("POST")
	r.HandleFunc("/borrow", middleware.CertifyApp(borrow.PostHandler)).Methods("POST")

	return r
}
