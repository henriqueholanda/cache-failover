package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/henriqueholanda/cache-failover/config"
	"github.com/henriqueholanda/cache-failover/bootstrap"
	"gopkg.in/mgo.v2"
	"encoding/json"
	"github.com/henriqueholanda/cache-failover/entities"
)

func GetIndex(response http.ResponseWriter, request *http.Request, params httprouter.Params) {

	database := bootstrap.Get("database").(*mgo.Session)

	collection := database.DB(config.GetConfigs().DatabaseName).C(config.GetConfigs().CustomerCollection)

	customers := entities.Customers{}
	err := collection.Find(nil).All(&customers)
	if err != nil {
		fmt.Println(err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	jsonResponse, _ := json.Marshal(customers)
	response.Write([]byte(jsonResponse))
}
