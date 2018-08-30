package manager

import (
	"github.com/onesafe/license_manager/db"
	"github.com/onesafe/license_manager/router"

	"github.com/jinzhu/gorm"
)

type Manager struct {
	apiRouter *router.APIRouter
	db        *gorm.DB
}

func NewManager() Manager {
	router := router.GetAPIRouter()

	database, err := db.GetDB()
	if err != nil {
		panic("Failed to get database for deploy manager: " + err.Error())
	}

	return Manager{
		apiRouter: router,
		db:        database,
	}
}
