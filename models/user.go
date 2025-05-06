package models

type User struct {
	ID           int    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	UserName     string `json:"user_name" gorm:"type:varchar(50);not null;column:user_name"`
	UserSurname  string `json:"user_surname" gorm:"type:varchar(50);not null;column:user_surname"`
	UserNickname string `json:"user_nickname" gorm:"type:varchar(75);unique;not null;column:user_nickname"`
	UserEmail    string `json:"user_email" gorm:"type:varchar(75);unique;not null;column:user_email"`
	UserPassword string `json:"user_password" gorm:"type:varchar(255);not null;column:user_password"`
	RoleID       int    `json:"role_id" gorm:"not null;column:role_id"`
}
