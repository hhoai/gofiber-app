package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	RoleID    int    `json:"role" gorm:"default:1"`
	SessionID string
}

type Account struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
	Address  string `form:"address"`
	Phone    string `form:"phone"`
	RoleID   int    `form:"role_id"`
}

type UserWithRowNumber struct {
	RowNumber int
	ID        int
	Name      string
	Email     string
	Address   string
	Phone     string
	RoleName  string
}

type PermissionWithRowNumber struct {
	RowNumber  int
	ID         int
	Permission string
}

type RolePermissionWithRowNumber struct {
	RowNumber      int
	ID             int
	RoleName       string
	PermissionName string
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
