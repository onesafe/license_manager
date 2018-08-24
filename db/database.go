package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB         *gorm.DB
	DB_INIT    bool
	DB_MIGRATE bool

	db_user   string
	db_passwd string
	db_host   string
	db_name   string
)

func SetConfig(user, passwd, host, dbname string) {
	db_user = user
	db_passwd = passwd
	db_host = host
	db_name = dbname
}

func GetDB() (*gorm.DB, error) {
	if DB_INIT {
		return DB, nil
	}

	config := db_user + ":" + db_passwd
	config += fmt.Sprintf("@tcp(%s)/", db_host)
	config += db_name
	config += "?charset=utf8mb4,utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open("mysql", config)
	if err != nil {
		return DB, err
	}

	DB_INIT = true
	return DB, err
}

func Migrate() error {
	if !DB_INIT {
		_, err := GetDB()
		if err != nil {
			return err
		}
	}

	if DB_MIGRATE {
		return nil
	}

	DB.AutoMigrate(&License_record{})

	DB_MIGRATE = true
	return nil
}

func Close() {
	DB.Close()
}
