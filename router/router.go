package router

import (
	"go-back/handler"
	"net/http"
)

func SetUpRouter() *http.ServeMux {

	router := http.NewServeMux()

	router.HandleFunc("/posts", handler.HandlePosts)

	router.HandleFunc("/health", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(`{"status":"ok"}`))
	})

	return router
}
