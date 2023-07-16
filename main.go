package main

import (
	"final-project/auth"
	"final-project/handler"
	"final-project/students"

	"github.com/kataras/iris/v12"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//iris init
	app := iris.New()
	//connect database use gorm
	dsn := "host=localhost user=root password=123qweasdzxc dbname=course port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		app.Logger()
	}
	//service auth
	authService := auth.AuthService("Prakerja")
	//student repository
	studentRepository := students.StudentRepository(db)
	//student Service
	studentService := students.StudentService(studentRepository)
	//student handler
	studentHandler := handler.StudentHandler(studentService, authService)
	//route  group
	v1 := app.Party("/v1")
	{
		v1.Post("/sign-up", studentHandler.RegisterStudent)
		v1.Post("/login", studentHandler.LoginStudent)
	}
	app.Listen(":8080")

}
