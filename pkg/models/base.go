// Package modeles
package models

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

//func init() {
//	var err error
//	_, err = getConnectedDB()
//	if err != nil {
//		panic(fmt.Errorf("connect to db err: %s", err))
//	}
//
//	err = AutoMigrate()
//	if err != nil {
//		panic(fmt.Errorf("migrate models table err: %s", err))
//	}
//}

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

type ListOptions struct {
	StartedAt time.Time
	EndedAt   time.Time
	Limit     int
}

// return cached connected db, create if not exists
func getConnectedDB() (*gorm.DB, error) {
	if db == nil {
		var err error
		if db, err = gorm.Open(sqlite.Open("/Users/fancy/go/src/github.com/yuansmin/health-recoder/data.db"), nil); err != nil {
			return nil, err
		}

	}
	return db, nil
}

func genConnectDBError(err error) error {
	return fmt.Errorf("can't connect to db: %s", err)

}

func AutoMigrate(url string) error {
	db, err := gorm.Open(sqlite.Open(url), nil)
	if err != nil {
		return err
	}

	// migrate models
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&ExerciseRecord{})
	db.AutoMigrate(&RecordGroup{})

	// todo
	return nil
}
