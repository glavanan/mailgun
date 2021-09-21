package main

import (
	"github.com/glavanan/mailgun/dbconnector"
	"github.com/glavanan/mailgun/router"
	log "github.com/sirupsen/logrus"
)

func main() {
	//The dsn should have been imported from a config file
	dbconnector.InitDB("host=0.0.0.0 user=postgres password=secretpass dbname=domaindb port=5432")
	r := router.GetRouter()
	//Port should have been set with a config file also
	err := r.Run(":" + "8080")
	if err != nil {
		log.Errorf("Error when running the router")
	}
}
