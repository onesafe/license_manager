package db

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type License_record struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT"`
	Product   string    `gorm:"not null;index:IDX_PRODUCT;type:varchar(32)"`
	Content   string    `gorm:"column:content;type:text"`
	CreatedAt time.Time `gorm:"not null"`
	CreatedBy string    `gorm:"not null;type:varchar(128)"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"type:varchar(128)"`
}

func (License_record) TableName() string {
	return "license_record"
}

func (l *License_record) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("create_at", time.Now())
	scope.SetColumn("update_at", time.Now())
	return nil
}

func (l *License_record) Save() *gorm.DB {
	return DB.Save(l)
}

func (l *License_record) GetLicenses(pageNum int, pageSize int, maps interface{}) {
	DB.Offset(pageNum).Limit(pageSize).Find(&l)
	return
}

func (l *License_record) GetLicense(maps interface{}) error {
	err := DB.Where(maps).First(&l).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *License_record) GetByProduct(product string) error {
	l.Product = product
	err := DB.First(l).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *License_record) Create() error {
	err := DB.Create(&l).Error
	if err != nil {
		return err
	}
	return nil
}
