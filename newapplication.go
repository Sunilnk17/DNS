package main

import (
	"drone-navigation-service/app/api"
	"drone-navigation-service/app/config"
	"drone-navigation-service/app/initializer"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	initializer.Initialize()
	serveRequest(api.GetRoutes())
}

func serveRequest(configuredRoutes http.Handler) {
	log.Print("########## SERVER STARTED ##########")
	error := http.ListenAndServe(
		config.GetConfig().Server.Port,
		handlers.CORS(
			handlers.AllowedMethods([]string{"GET"}),
			handlers.AllowedHeaders([]string{"User-Agent", "If-Modified-Since", "Cache-Control",
				"Content-Type"}),
			handlers.MaxAge(600),
		)(configuredRoutes))
	if error != nil {
		log.Print("Error during server startup: ", error)
	}
	os.Exit(1)
}
