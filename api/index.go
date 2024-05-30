package api

import (
	controllers "billhub/controllers"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "Token"},
	})
	handler := corsMiddleware.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			Home(w, r)
			return
		case "/billboard":
			controllers.KSI_Billboard(w, r)
			return
		case "/login":
			controllers.KSI_Login(w, r)
			return
		case "/register":
			controllers.KSI_Register(w, r)
			return
		case "/sewa":
			controllers.KSI_Sewa(w, r)
			return
		case "/user":
			controllers.KSI_User(w, r)
			return
		default:
			http.Error(w, "404 Not Found", http.StatusNotFound)
		}
	}))

	handler.ServeHTTP(w, r)
}

func Home(w http.ResponseWriter, r *http.Request) {
	type entryPoint map[string]any
	data, err := json.Marshal(entryPoint{
		"message": "Welcome to KSI Billboard API",
		"status": 200,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s\n", data);
}