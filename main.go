package main

import (
	"github.com/kataras/iris/v12"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := iris.New()
	dsn := "host=localhost user=root password=123qweasdzxc dbname=course port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		app.Logger()
	}

	app.Listen(":8080")

}
