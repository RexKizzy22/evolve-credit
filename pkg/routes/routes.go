package routes

import (
	"evolve-credit/pkg/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Routes() http.Handler {
	router := httprouter.New()
	
	router.Handler(http.MethodGet, "/favicon.ico", http.NotFoundHandler())
	router.HandlerFunc(http.MethodGet, "/", handlers.Index)
	router.HandlerFunc(http.MethodGet, "/v1/users", handlers.GetUsers)
	router.HandlerFunc(http.MethodGet, "/v1/user/:email", handlers.GetUserByEmail)
	return router
}
