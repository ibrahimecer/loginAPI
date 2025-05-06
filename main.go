package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimecer/loginAPI/config"
	"github.com/ibrahimecer/loginAPI/routes"
)

func main() {
	router := gin.New()
	config.ConnectDB()
	routes.SetupRoutes(router)
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/student_login.html", "./static/student_login.html")
	router.StaticFile("/teacher_login.html", "./static/teacher_login.html")
	router.StaticFile("/student.html", "./static/student.html")
	router.StaticFile("/student_register.html", "./static/student_register.html")
	router.StaticFile("/teacher.html", "./static/teacher.html")
	router.StaticFile("/detail.html", "./static/detail.html")
	router.Run(":4569")
}
