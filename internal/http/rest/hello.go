package rest

import "net/http"

func Hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("Hello, world - now http!\n"))
	w.WriteHeader(http.StatusOK)
}
