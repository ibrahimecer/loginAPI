package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ibrahimecer/loginAPI/config"
	"github.com/ibrahimecer/loginAPI/models"
	"github.com/ibrahimecer/loginAPI/utils"
)

func StudentRegister(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri formatı"})
		return
	}

	if input.RoleID != 1 {
		c.JSON(http.StatusForbidden, gin.H{"error": "Sadece öğrenciler kayıt olabilir"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.UserPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Şifre hashlenemedi"})
		return
	}
	input.UserPassword = hashedPassword

	// Kullanıcıyı kaydet
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kullanıcı oluşturulamadı: " + err.Error()})
		return
	}

	// Debug amaçlı log
	fmt.Println("Yeni kullanıcı ID:", input.ID)

	student := models.Student{
		Exam1:     0,
		Exam2:     0,
		Exam3:     0,
		StudentID: input.ID,
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Öğrenci tablosuna eklenemedi: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Öğrenci kaydı başarılı"})
}

func handleLogin(c *gin.Context, expectedRoleID int) {
	type LoginInput struct {
		EmailOrNickname string `json:"email_or_nickname"`
		Password        string `json:"password"`
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz giriş verisi"})
		return
	}

	var user models.User
	if err := config.DB.Where("user_email = ? OR user_nickname = ?", input.EmailOrNickname, input.EmailOrNickname).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kullanıcı bulunamadı"})
		return
	}

	if user.RoleID != expectedRoleID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Yetkiniz yok"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.UserPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Şifre hatalı"})
		return
	}
	if user.RoleID == 1 {
		var student models.Student
		config.DB.Where("student_id = ?", user.ID).First(&student)
		c.JSON(http.StatusOK, gin.H{
			"message":    "Giriş başarılı",
			"user_id":    user.ID,
			"student_id": student.StudentID,
			"user_name":  user.UserName,
			"exam1":      student.Exam1,
			"exam2":      student.Exam2,
			"exam3":      student.Exam3,
			"average":    student.Average,
			"role_id":    user.RoleID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Giriş başarılı",
			"user_id":   user.ID,
			"user_name": user.UserName,
			"role_id":   user.RoleID,
		})
	}
}

func GetAllStudents(c *gin.Context) {
	var results []map[string]interface{}

	query := `
		SELECT 
			s.student_id, 
			u.user_name, 
			u.user_surname,
			s.exam1, 
			s.exam2, 
			s.exam3, 
			s.average
		FROM students s
		JOIN users u ON s.student_id = u.id
	`

	if err := config.DB.Raw(query).Scan(&results).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Veri alınamadı: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}

func GetStudentGrades(c *gin.Context) {
	studentID := c.Param("id")

	var result map[string]interface{}

	query := `
		SELECT 
			s.student_id,
			u.user_name,
			u.user_surname,
			u.user_email,
			u.user_nickname,
			s.exam1,
			s.exam2,
			s.exam3,
			s.average
		FROM students s
		JOIN users u ON s.student_id = u.id
		WHERE s.student_id = ?
	`

	if err := config.DB.Raw(query, studentID).Scan(&result).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Öğrenci bilgisi bulunamadı: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAllStudentGrades(c *gin.Context) {
	var students []models.Student
	if err := config.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Öğrenci notları alınamadı"})
		return
	}

	c.JSON(http.StatusOK, students)
}
func UpdateStudentGrades(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Exam1 int `json:"exam1"`
		Exam2 int `json:"exam2"`
		Exam3 int `json:"exam3"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz veri"})
		return
	}

	average := float64(input.Exam1+input.Exam2+input.Exam3) / 3.0

	if err := config.DB.Model(&models.Student{}).
		Where("student_id = ?", id).
		Updates(models.Student{
			Exam1:   input.Exam1,
			Exam2:   input.Exam2,
			Exam3:   input.Exam3,
			Average: average,
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Notlar güncellenemedi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notlar ve ortalama güncellendi"})
}

func StudentLogin(c *gin.Context) {
	handleLogin(c, 1) // 1 = öğrenci
}

func TeacherLogin(c *gin.Context) {
	handleLogin(c, 2) // 2 = öğretmen
}
