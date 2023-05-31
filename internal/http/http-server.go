package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mirfaiziev/golang-grpc-sandbox/internal/http/rest"
)

func StartHttpServer() {
	router := chi.NewRouter()

	router.Get("/", rest.Hello)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("cannot start http server %v", err)
	} else {
		log.Println("start http service")
	}
}
