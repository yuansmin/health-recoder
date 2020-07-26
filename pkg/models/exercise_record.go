package models

import (
	"time"

	"gorm.io/gorm"
)

type ExerciseRecord struct {
	BaseModel

	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	UserID     uint      `json:"user_id"`
	Category   `gorm:"foreignKey:ID" json:"category"`
	RecordList []Record `json:"record_list" gorm:"foreignKey:ID"`
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

func ListUserExerciseRecord(user *User, opts ...ListOptions) ([]ExerciseRecord, error) {
	// todo use ListOptions filter result
	db, err := getConnectedDB()
	if err != nil {
		return nil, err
	}

	var recordList []ExerciseRecord
	if err = db.Where("user_id=?", user.ID).Find(&recordList).Error; err != nil {
		return nil, err
	}
	return recordList, nil
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
