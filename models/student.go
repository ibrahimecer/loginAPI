package models

type Student struct {
	ID        int     `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Exam1     int     `json:"exam1" gorm:"not null;column:exam1"`
	Exam2     int     `json:"exam2" gorm:"not null;column:exam2"`
	Exam3     int     `json:"exam3" gorm:"not null;column:exam3"`
	Average   float64 `json:"average" gorm:"column:average"` // DB hesaplar
	StudentID int     `json:"student_id" gorm:"column:student_id"`
}
