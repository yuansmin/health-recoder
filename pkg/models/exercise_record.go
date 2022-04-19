package models

import (
	"time"

	"gorm.io/gorm"
)

type ExerciseRecord struct {
	BaseModel

	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	UserID    uint      `json:"user_id"`
	Category  `gorm:"foreignKey:" json:"category"`
	// all groups count
	Count uint `json:"count"`
}

type RecordGroup struct {
	BaseModel

	RecordID  uint      `json:"record_id"`
	StartedAt time.Time `json:"started_at"`
	EndedAt   time.Time `json:"ended_at"`
	Count     uint      `json:"count"`
}

func (r *ExerciseRecord) IsPrimaryKeySet() bool {
	if r.ID == 0 {
		return false
	}

	return true
}

func (r *ExerciseRecord) GetElapse() time.Duration {
	return r.EndedAt.Sub(r.StartedAt)
}

func CreateExerciseRecord(record *ExerciseRecord) error {
	db, err := getConnectedDB()
	if err != nil {
		return err
	}
	if err := db.Create(record).Error; err != nil {
		return err
	}
	return nil
}

func GetExerciseRecord(r *ExerciseRecord) error {
	if !r.IsPrimaryKeySet() {
		return gorm.ErrorPrimaryKeyRequired
	}
	db, err := getConnectedDB()
	if err != nil {
		return err
	}
	if err := db.Find(r).Error; err != nil {
		return err
	}
	return nil

}

func AppendRecordToExercise(r *Record, exercise *ExerciseRecord) error {
	// todo: implement
	return nil
}

func DeleteExerciseRecord(record *ExerciseRecord) error {
	db, err := getConnectedDB()
	if err != nil {
		return err
	}

	if err = db.Delete(&record).Error; err != nil {
		return err
	}
	return nil
}
