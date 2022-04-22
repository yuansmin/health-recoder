package apis

import "github.com/yuansmin/health-recoder/pkg/models"

type ListExerciseRecordRequest struct {
	Offset int `json:"offset" binding:"-"`
	Limit  int `json:"limit" binding:"-"`
}

type ListExerciseRecordResponse struct {
	Total int                      `json:"total"`
	Items []*models.ExerciseRecord `json:"items"`
}
