package main

import (
	"fmt"

	"github.com/onesafe/license_manager/config/env"
	"github.com/onesafe/license_manager/db"
	"github.com/onesafe/license_manager/router"
)

func main() {
	fmt.Println("Staring Lincese Manager")

	DB_init()

	Router_init()

}

func DB_init() {
	db_pass, err := env.GetDBPass()
	if err != nil {
		panic("Failed to get DB password: " + err.Error())
	}
	db_host, err := env.GetDBHost()
	if err != nil {
		panic("Failed to get DB Host: " + err.Error())
	}
	db_host_port := db_host + ":" + env.GetDBPort()
	db_user := env.GetDBUser()
	db_name := env.GetDBName()
	db.SetConfig(db_user, db_pass, db_host_port, db_name)
	err = db.Migrate()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	fmt.Println("Database initialized")
}

func Router_init() {
	fmt.Println("Starting API Router")
	Router := router.GetAPIRouter()

	// Run API Router
	Router.Run()
}
