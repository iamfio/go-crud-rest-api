package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamfio/crud-rest-api/controllers"
)

func main() {
	LoadAppConfig()

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)

	log.Printf("Starting Server on PORT %s\n", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))

}

func RegisterProductRoutes(router *mux.Router) {
	var muxBase = "/api/products"
	router.HandleFunc(muxBase, controllers.GetProducts).Methods("GET")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.GetProductById).Methods("GET")
	router.HandleFunc(muxBase, controllers.CreateProduct).Methods("POST")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc(fmt.Sprintf("%s/{id}", muxBase), controllers.DeleteProduct).Methods("DELETE")
}
