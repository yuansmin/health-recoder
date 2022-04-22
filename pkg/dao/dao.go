// Package dao is for database access
package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(url string) (*Dao, error) {
	var db *gorm.DB
	var err error
	if db, err = gorm.Open(sqlite.Open(url), nil); err != nil {
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
