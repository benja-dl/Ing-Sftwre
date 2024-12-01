package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"type:varchar(50);not null"`
	LastName  string `gorm:"type:varchar(50);not null"`
	Mail      string `gorm:"type:varchar(256); not null;"`
	Password  string `gorm:"type:varchar(256); not null;"`
}

type UserRegisterData struct {
	FirstName string `json:"firstName" xml:"firstName" form:"firstName"`
	LastName  string `json:"lastName" xml:"lastName" form:"lastName"`
	Mail      string `json:"mail" xml:"mail" form:"mail"`
	Password  string `json:"password" xml:"password" form:"password"`
}

func CreateNewUser(userForm *UserRegisterData, db *gorm.DB) error {
	// Buscamos si el mail ya se encuentra registrado
	var user User
	result := db.Where("mail = ?", userForm.Mail).First(&user)

	if result.Error == nil {
		return errors.New("user already exists")
	}
	// Error en la base de datos
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	user = User{
		FirstName: userForm.FirstName,
		LastName:  userForm.LastName,
		Mail:      userForm.Mail,
		Password:  userForm.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
