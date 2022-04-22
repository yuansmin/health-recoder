package dao

import (
	"github.com/yuansmin/health-recoder/pkg/models"
	"gorm.io/gorm"
)

type exerciseRecord struct {
	db *gorm.DB
}

// List ExerciseRecord
func (e *exerciseRecord) List(userID uint, offset, limit int) ([]*models.ExerciseRecord, error) {
	var recordList []*models.ExerciseRecord
	var err error
	if err = e.db.Where(&models.ExerciseRecord{UserID: userID}).
		Offset(offset).
		Limit(limit).
		Find(&recordList).Error; err != nil {
		return nil, err
	}
	return recordList, nil
}

// Get if userID == 0, will not filter by userID
func (e *exerciseRecord) Get(id uint, userID uint) (*models.ExerciseRecord, error) {
	var er models.ExerciseRecord
	err := e.db.Where(&models.ExerciseRecord{BaseModel: models.BaseModel{ID: id}, UserID: userID}).
		Take(&er).Error
	if err != nil {
		return nil, err
	}

	return &er, nil
}

func (e *exerciseRecord) Create(er *models.ExerciseRecord) error {
	if err := e.db.Create(er).Error; err != nil {
		return err
	}

	return nil
}
