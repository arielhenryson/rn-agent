package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type AppRouter struct {
	*mux.Router
}

func NewRouter() *AppRouter {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)

	return &AppRouter{
		Router: r,
	}
}

func StartApiServer() {
	http.Handle("/", NewRouter())

	log.Println("Serving on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
