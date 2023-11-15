package db

import (
	"fmt"
	"time"

	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BaseModel struct {
	ID           uint32 `gorm:"primary_key" json:"id"`
	CreatedDate  uint32 `json:"created_date"`
	ModifiedDate uint32 `json:"modified_date"`
	DeletedDate  uint32 `json:"deleted_date"`
}

func NewDBEngine() (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db, err := gorm.Open("mysql", fmt.Sprintf(s,
		"root",
		"123456",
		"127.0.0.1:3306",
		"blog_service",
		"utf8",
		true,
	))
	if err != nil {
		return nil, err
	}

	db.LogMode(true)
	db.SingularTable((true))
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)

	otgorm.AddGormCallbacks(db)
	return db, nil
}

func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedDate"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedDate"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}
