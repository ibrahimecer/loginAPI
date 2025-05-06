package models

type Role struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	RoleName string `json:"role_name" gorm:"type:varchar(50);unique;not null;column:role_name"`
}
