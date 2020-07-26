package models

type HealthRecord struct {
	BaseModel
	Record
}

func CreateHealthRecord(record *HealthRecord) error {
	// todo: implement
	return nil
}

func ListUserHealthRecord(user *User, opt ...ListOptions) ([]HealthRecord, error) {
	// todo: implement
	return nil, nil
}

func DeleteHealthRecord(record *HealthRecord) error {
	// todo: implement
	return nil
}
