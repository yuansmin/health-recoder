// Package models
package models

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func init() {
	err := InitUserData()
	if err != nil {
		log.Errorf("init DB err, %s", err)
		os.Exit(1)
	}
}

// for api expose
type User struct {
	BaseModel
	Name string `binding:"required"`
	// for third platform user id
	OpenID string
	Email  string
	Tel    string
}

type InternalUser struct {
	User
	PasswordHash string `json:"password_hash"`
}

// hide secret fields
func (u *InternalUser) ToSafeUser() User {
	return u.User
}

func ListUsers() ([]User, error) {
	db, err := getConnectedDB()
	if err != nil {
		return nil, err
	}

	users := []User{}
	db.Find(&users)
	return users, nil
}

func GetUser(user *User) error {
	if user.ID == 0 {
		return gorm.ErrorPrimaryKeyRequired
	}
	db, err := getConnectedDB()
	if err != nil {
		return genConnectDBError(err)
	}

	if err := db.First(user).Error; err != nil {
		return fmt.Errorf("query user err: %s", err)
	}
	return nil
}

func CreateUser(user *User) error {
	db, err := getConnectedDB()
	if err != nil {
		return genConnectDBError(err)
	}

	if err := db.Create(user).Error; err != nil {
		return fmt.Errorf("create user err: %s", err)
	}
	return nil
}

func DeleteUser(user *User) error {
	if user.ID == 0 {
		return gorm.ErrorPrimaryKeyRequired
	}
	db, err := getConnectedDB()
	if err != nil {
		return genConnectDBError(err)
	}
	if err := db.Delete(user).Error; err != nil {
		return err
	}

	return nil
}

func InitUserData() error {

	// user := User{Name: "purple"}
	// db.Create(&user)

	return nil
}
