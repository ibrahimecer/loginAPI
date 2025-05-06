package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimecer/loginAPI/controller"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/student/register", controller.StudentRegister)
	router.POST("/student/login", controller.StudentLogin)
	router.GET("/student/grades/:id", controller.GetStudentGrades)
	router.POST("/teacher/login", controller.TeacherLogin)
	router.GET("/teacher/detail/:id", controller.GetStudentGrades)
	router.PUT("/teacher/grades/:id", controller.UpdateStudentGrades)
	router.GET("/teacher/grades", controller.GetAllStudents)
}
