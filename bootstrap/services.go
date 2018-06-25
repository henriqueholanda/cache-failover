package bootstrap

import (
	"github.com/henriqueholanda/cache-failover/services"
)

func LoadServices() {
	database()
}

func database() {
	serviceContainer := Container{
		name: "database",
		instance: services.GetDatabaseConnection(),
	}
	Container.Set(serviceContainer)
}
