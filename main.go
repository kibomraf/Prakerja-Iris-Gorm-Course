package main

import (
	"final-project/handler"
	"final-project/students"

	"github.com/kataras/iris/v12"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := iris.New()
	dsn := "host=localhost user=root password=123qweasdzxc dbname=course port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		app.Logger()
	}
	//student repository
	studentRepository := students.StudentRepository(db)
	//student Service
	studentService := students.StudentService(studentRepository)
	//student handler
	studentHandler := handler.StudentHandler(studentService)
	//route  group
	v1 := app.Party("/v1")
	{
		v1.Post("/sign-up", studentHandler.RegisterStudent)
	}
	app.Listen(":8080")

}
