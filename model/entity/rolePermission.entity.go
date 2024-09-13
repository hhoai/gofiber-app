package entity

type Role struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Role string `gorm:"unique;not null" form:"role" json:"role"`
}

type Permission struct {
	ID         int    `gorm:"primaryKey;autoIncrement"`
	Permission string `gorm:"unique;not null" form:"permission" json:"permission"`
}

type RolePermission struct {
	RoleID       int `gorm:"primaryKey"`
	PermissionID int `gorm:"primaryKey"`
}

type PermissionsRequest struct {
	Permissions []int `json:"permission" form:"permission"`
}

type SalesData struct {
	ID          uint    `json:"id"`
	Month       string  `json:"month"`
	SalesAmount float64 `json:"sales_amount"`
	CreatedAt   string  `json:"created_at"`
}
