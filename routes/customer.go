package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/henriqueholanda/cache-failover/controllers"
)

const BASE_ROUTE = "/customer"

func LoadCustomerRoutes(router *httprouter.Router)  {
	router.GET(BASE_ROUTE, controllers.GetIndex)
}
