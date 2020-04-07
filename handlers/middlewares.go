package handlers

import (
	"fmt"
	"net/http"

	"gitlab.com/jorgechavezrnd/go_rest/utils"
)

type customeHandler func(w http.ResponseWriter, r *http.Request)

// Authentication ...
func Authentication(function customeHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsAuthenticated(r) {
			http.Redirect(w, r, "/users/login", http.StatusSeeOther)
			return
		}

		function(w, r)
	})
}

// MiddlewareTwo ...
func MiddlewareTwo(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logica aqui!
		fmt.Println("Este es el segundo wrap!")
		handler.ServeHTTP(w, r)
	})
}
