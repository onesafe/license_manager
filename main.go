package main

import (
	"fmt"

	"github.com/onesafe/license_manager/config/env"
	"github.com/onesafe/license_manager/db"
	"github.com/onesafe/license_manager/manager"
	"github.com/onesafe/license_manager/router"

	_ "github.com/onesafe/license_manager/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title License Manager API
// @version 0.1
// @description This is a doc describe API of License Manager
// @BasePath /license-manager/v1
// @contact.email onesafe@163.com
func main() {

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

	manager.GetLicenseManager().RegisterPath()

	// doc
	Router.Register("GET", "/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run API Router
	Router.Run()
}
