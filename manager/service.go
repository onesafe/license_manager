package manager

import (
	"log"
	"time"

	"github.com/onesafe/license_manager/db"
	"github.com/onesafe/license_manager/views"
)

func registerLicense(license *views.License, content string) (err error) {
	// check license expired
	valid, err := license.IsExpired()
	if valid == false {
		return
	}

	// create or update license record
	lr := &db.License_record{}
	err = lr.GetByProduct(license.Product)
	if err != nil {
		log.Println("Creating new license record for " + license.Product)
		lr.Product = license.Product
		lr.Content = content
		lr.CreatedBy = "onesafe"
		lr.UpdatedBy = "onesafe"

		err = lr.Create()
		if err != nil {
			log.Println("Create new License record error")
			return
		}
	} else {
		log.Println("Updating existed license record for " + license.Product)
		lr.Content = content
		lr.UpdatedAt = time.Now()
		lr.UpdatedBy = "onesafe"

		lr.Save()
	}

	return
}
