package models

import "time"

type ExerciseRecord struct {
	BaseModel

	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
	UserID     uint      `json:"user_id"`
	Category   `gorm:"foreignKey:ID" json:"category"`
	RecordList []Record `json:"record_list" gorm:"foreignKey:ID"`
}

func (r *ExerciseRecord) GetElapse() time.Duration {
	return r.EndedAt.Sub(r.StartedAt)
}

func CreateExerciseRecord(record *ExerciseRecord) error {
	// todo: implement
	return nil
}

func ListUserExerciseRecord(user *User) ([]ExerciseRecord, error) {
	// todo: implement
	return nil, nil
}

func DeleteExerciseRecord(record *ExerciseRecord) error {
	// todo: implement
	return nil
}
