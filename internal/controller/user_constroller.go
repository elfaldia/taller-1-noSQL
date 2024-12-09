package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}


func NewUserController(service service.UserService) *UserController {
	return &UserController{UserService: service}
}




func (controller * UserController) FindAll(ctx *gin.Context) {

	data, err := controller.UserService.FindAll()
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}
	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)

}


func (controller *UserController) FindById(ctx *gin.Context) {
	id := ctx.Param("user_id")

	data, err := controller.UserService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}


func (controller *UserController) CreateUser(ctx *gin.Context) {
	var req *request.RegisterUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		if msg == "EOF" {
			msg = "DEBE SER EN FORMATO JSON EL BODY"
		}
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: msg,
		})
		return
	}
	err := controller.UserService.RegisterUser(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   nil,
	}
	ctx.JSON(http.StatusCreated, res)
}


func (controller *UserController) DeleteUser(ctx *gin.Context) {

	userId := ctx.Param("user_id")

	err := controller.UserService.DeleteUsuario(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, res)
}


func (controller *UserController) Login(ctx *gin.Context) {


	var req *request.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		msg := err.Error()
		if msg == "EOF" {
			msg = "DEBE SER EN FORMATO JSON EL BODY"
		}
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: msg,
		})
		return
	}
	data, err := controller.UserService.LoginUser(req)

	if err != nil {
		ctx.JSON(http.StatusNotFound, response.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)




}



