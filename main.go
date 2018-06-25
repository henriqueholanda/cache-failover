package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/henriqueholanda/cache-failover/config"
	"github.com/julienschmidt/httprouter"
	"github.com/henriqueholanda/cache-failover/routes"
	"github.com/henriqueholanda/cache-failover/bootstrap"
)

func main() {
	bootstrap.LoadServices()

	router := httprouter.New()
	routes.LoadCustomerRoutes(router)

	log.Printf("Server starting on port %v\n", config.GetConfigs().Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.GetConfigs().Port), router))

}
