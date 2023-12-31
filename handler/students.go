package handler

import (
	"final-project/auth"
	"final-project/helper"
	"final-project/students"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type handler struct {
	service     students.Service
	authservice auth.Service
}

func StudentHandler(service students.Service, auth auth.Service) *handler {
	return &handler{service, auth}
}

// handler Register User
func (h *handler) RegisterStudent(ctx iris.Context) {
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
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		APIResponse := helper.APIresponse(iris.StatusUnprocessableEntity, "falied to create account", "failed", err)
		ctx.JSON(APIResponse)
		return
	}
	//call bussines logic
	//busines logic check email
	checkEmail, err := h.service.CheckEmailAvailibity(input.Email)
	if err != nil || !checkEmail {
		ctx.StatusCode(iris.StatusFound)
		APIResponse := helper.APIresponse(iris.StatusFound, "email has been created", "failed", nil)
		ctx.JSON(APIResponse)
		return
	}
	//bussiness logic save user
	newStudent, err := h.service.CreateStudent(input)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		APIResponse := helper.APIresponse(iris.StatusInternalServerError, "falied to create account", "failed", nil)
		ctx.JSON(APIResponse)
		return
	}
	token, err := h.authservice.GenerateToken(newStudent.Id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		APIResponse := helper.APIresponse(iris.StatusInternalServerError, "failed to create token", "failed", nil)
		ctx.JSON(APIResponse)
		return
	}
	formatter := students.FormatStudents(newStudent, token)
	APIResponse := helper.APIresponse(iris.StatusOK, "account has been created", "successfully", formatter)
	ctx.JSON(APIResponse)
	ctx.StatusCode(iris.StatusOK)
}
func (h *handler) LoginStudent(ctx iris.Context) {
	//mapping input user
	var input students.Login
	err := ctx.ReadJSON(&input)
	if err != nil {
		APIResponse := helper.APIresponse(iris.StatusBadRequest, "falied to login", "failed", nil)
		ctx.JSON(APIResponse)
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	//validate struct with validator
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		APIResponse := helper.APIresponse(iris.StatusUnprocessableEntity, "falied to login", "failed", err)
		ctx.JSON(APIResponse)
		return
	}
	//call bussiness logic
	loginStudent, err := h.service.LoginStudent(input)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		APIResponse := helper.APIresponse(iris.StatusInternalServerError, "email or password is wrong", "failed", err)
		ctx.JSON(APIResponse)
		return
	}
	token, err := h.authservice.GenerateToken(loginStudent.Id)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		APIResponse := helper.APIresponse(iris.StatusInternalServerError, "failed to create token", "failed", nil)
		ctx.JSON(APIResponse)
		return
	}
	formatter := students.FormatStudents(loginStudent, token)
	APIResponse := helper.APIresponse(iris.StatusOK, "login successfull", "success", formatter)
	ctx.JSON(APIResponse)
	ctx.StatusCode(iris.StatusOK)
}
