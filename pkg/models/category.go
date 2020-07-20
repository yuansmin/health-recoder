package models

type CategoryType uint16

const (
	CategoryTypeExercise = 0
	CategoryTypeHealth   = 1
)

type Category struct {
	BaseModel
	Type        CategoryType `json:"type"`
	Description string       `json:"description"`
}

func ListCategories() ([]Category, error) {
	//todo: implement
	return nil, nil
}
