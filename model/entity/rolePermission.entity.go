package entity

type Role struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Role string `gorm:"unique;not null"`
}

type Permission struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Permission string `gorm:"unique;not null"`
}

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`
}
