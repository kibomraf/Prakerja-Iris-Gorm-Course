package handler

import (
	"final-project/helper"
	"final-project/students"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type handler struct {
	service students.Service
}

func StudentHandler(service students.Service) *handler {
	return &handler{service}
}

// handler Register User
func (s *handler) RegisterStudent(ctx iris.Context) {
	//json parameter
	var input students.Input
	err := ctx.ReadJSON(&input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		APIResponse := helper.APIresponse(iris.StatusBadRequest, "falied to create account", "failed", nil)
		ctx.JSON(APIResponse)
		return
	}
	//validator parameter json
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		APIResponse := helper.APIresponse(iris.StatusBadRequest, "falied to create account", "failed", err)
		ctx.JSON(APIResponse)
		return
	}
	//call bussines logic
	//busines logic check email
	checkEmail, err := s.service.CheckEmailAvailibity(input.Email)
	if err != nil || !checkEmail {
		ctx.StatusCode(iris.StatusFound)
		APIResponse := helper.APIresponse(iris.StatusFound, "email has been created", "failed", nil)
		ctx.JSON(APIResponse)
		return
	}
	//bussiness logic save user
	newStudent, err := s.service.CreateUser(input)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		APIResponse := helper.APIresponse(iris.StatusInternalServerError, "falied to create account 3", "failed", nil)
		ctx.JSON(APIResponse)
		return
	}
	formatter := students.FormatStudents(newStudent, "")
	APIResponse := helper.APIresponse(iris.StatusOK, "account has been created", "successfully", formatter)
	ctx.JSON(APIResponse)
	ctx.StatusCode(iris.StatusOK)
}
