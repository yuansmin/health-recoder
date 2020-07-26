package models

import "time"

type Record struct {
	BaseModel

	ExerciseID uint      `json:"-"`
	Count      int       `json:"count"`
	StartedAt  time.Time `json:"started_at"`
	EndedAt    time.Time `json:"ended_at"`
}

func (r *Record) GetElapse() time.Duration {
	return r.EndedAt.Sub(r.StartedAt)
}

func CreateRecord(r *Record) error {
	// todo: implement
	return nil
}

func DeleteRecord(r *Record) error {
	// todo: implement
	return nil
}

func UpdateRecord(r *Record) error {
	// todo: implement
	return nil
}
