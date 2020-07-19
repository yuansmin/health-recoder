// Package modeles
package models

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	_, err := getGlobalDB()
	if err != nil {
		log.Errorf("init DB err: %s", err)
		os.Exit(1)
	}

	err = AutoMigrate()
	if err != nil {
		log.Errorf("migrate models table err: %s", err)
		os.Exit(1)
	}
}

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

var db *gorm.DB

// return cached db, create if not exists
func getGlobalDB() (*gorm.DB, error) {
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

func AutoMigrate() error {
	db, err := getGlobalDB()
	if err != nil {
		return err
	}

	// migrate models
	db.AutoMigrate(&User{})

	// todo
	return nil
}
