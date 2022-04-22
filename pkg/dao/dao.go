// Package dao is for database access
package dao

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// todo: refine logger
func New(url string) (*Dao, error) {
	var db *gorm.DB
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	if db, err = gorm.Open(sqlite.Open(url), &gorm.Config{Logger: newLogger}); err != nil {
		return nil, err
	}

	return &Dao{
		url:            url,
		db:             db,
		exerciseRecord: &exerciseRecord{db: db},
	}, nil
}

type Dao struct {
	url            string
	db             *gorm.DB
	exerciseRecord *exerciseRecord
}

func (d *Dao) ExerciseRecord() *exerciseRecord {
	return d.exerciseRecord
}
