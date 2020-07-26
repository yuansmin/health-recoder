package models

type CategoryType uint16

const (
	// eg: pull-push, sit-up
	CategoryTypeExercise = 0
	// eg: weight, height
	CategoryTypeHealth = 1
)

type Category struct {
	BaseModel
	Type        CategoryType `json:"type"`
	Description string       `json:"description"`
}

func ListCategories(categoryType CategoryType) ([]Category, error) {
	//todo: implement
	return nil, nil
}
