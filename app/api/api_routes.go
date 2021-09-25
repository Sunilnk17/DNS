package api

import (
	"drone-navigation-service-master/app/api/controller"
	"drone-navigation-service-master/dns_middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	getLocation = "GetLocation"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	v1UnAuthenticatedRouter := r.PathPrefix("/internal/v1").Subrouter()
	v1UnAuthenticatedRouter.HandleFunc("/health_check", controller.GetHeartBeat).Methods(http.MethodGet).Name("GetHeartBeat")

	v1Router := r.PathPrefix("/api/v1").Subrouter()

	//GET -- READ
	v1Router.HandleFunc("/sectors/{sectorId}/drones", controller.GetLocation).Methods(http.MethodGet).Name(getLocation)

	addMiddlewares(v1Router)

	return r
}

func addMiddlewares(routes *mux.Router) {
	log.Print("Adding Middlewares")
	routes.Use(dns_middleware.GenerateRequestIdHandler)

}
