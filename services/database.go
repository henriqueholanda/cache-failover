package services

import (
	"gopkg.in/mgo.v2"
	"github.com/henriqueholanda/cache-failover/config"
)

func GetDatabaseConnection() *mgo.Session {
	session, err := mgo.Dial(config.GetConfigs().DatabaseUrl)
	if err != nil {
		panic(err)
	}
	return session
}
