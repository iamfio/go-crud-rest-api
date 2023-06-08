package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamfio/crud-rest-api/entities"
	"github.com/iamfio/crud-rest-api/repos"
	"github.com/iamfio/crud-rest-api/routers"
)

func main() {
	LoadAppConfig()

	router := mux.NewRouter().StrictSlash(true)

	RegisterProductRoutes(router)
	RegisterBrandRoutes(router)

	log.Printf("Starting Server on PORT %s\n", AppConfig.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	var productRepo repos.GenericRepository[entities.Product] = repos.NewProductRepository()
	routers.NewGenericRouter[entities.Product, *repos.ProductRepository]("/api/products", router, &productRepo)
}

func RegisterBrandRoutes(router *mux.Router) {
	var brandRepo repos.GenericRepository[entities.Brand] = repos.NewBrandRepository()
	routers.NewGenericRouter[entities.Brand, *repos.BrandRepository]("/api/brands", router, &brandRepo)
}
