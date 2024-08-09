package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Role     int    `json:"role" gorm:"default:1"`
}

func HashPassword(user *UserEntity, password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func CheckPassword(user *UserEntity, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// BeforeCreate là hook GORM, chạy trước khi tạo một bản ghi mới.
func (user *UserEntity) BeforeCreate(tx *gorm.DB) (err error) {
	if err := HashPassword(user, user.Password); err != nil {
		return err
	}
	return nil
}

// BeforeUpdate là hook GORM, chạy trước khi cập nhật một bản ghi.
func (user *UserEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := HashPassword(user, user.Password); err != nil {
		return err
	}
	return nil
}
