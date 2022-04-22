package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/yuansmin/health-recoder/pkg/apis"

	"github.com/stretchr/testify/require"
	"github.com/yuansmin/health-recoder/pkg/models"
)

func TestListExercises(t *testing.T) {
	router, dao, clean := setupRouter()
	defer clean()

	r := require.New(t)
	now := time.Now()
	r.NoError(dao.ExerciseRecord().Create(&models.ExerciseRecord{
		BaseModel: models.BaseModel{},
		StartedAt: now,
		EndedAt:   now.Add(2 * time.Minute),
		UserID:    1,
		Count:     5,
	}))
	r.NoError(dao.ExerciseRecord().Create(&models.ExerciseRecord{
		BaseModel: models.BaseModel{},
		StartedAt: now.Add(2 * time.Hour),
		EndedAt:   now.Add(3 * time.Hour),
		UserID:    1,
		Count:     50,
	}))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/exercises", nil)
	router.ServeHTTP(w, req)

	r.Equal(200, w.Code, w.Body.String())
	var resp apis.ListExerciseRecordResponse
	r.NoError(json.Unmarshal(w.Body.Bytes(), &resp))
	r.Equal(2, resp.Total)
	r.Equal(now.Unix(), resp.Items[0].CreatedAt.Unix())
}
