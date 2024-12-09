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


// @BasePath /user
// @Summary Devuelve todos los users de la base de datos
// @Description get users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /user [get]
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


// @BasePath /curso
// @Summary Devuelve un user
// @Description get user a partir del ID
// @Tags user
// @Param user_id path string true "email"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Router /user/{user_id} [get]
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


// @BasePath /user
// @Summary Crea un user
// @Description crea un user con todos sus componentes
// @Tags user
// @Param  curso body request.RegisterUserRequest true "json del curso"
// @Accept json
// @Produce json
// @Success 201 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /user [post]
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

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes a user by their user ID.
// @Tags user
// @Param user_id path string true "User ID"
// @Produce json
// @Success 200 {object} response.Response "Success Response"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /user/{user_id} [delete]
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

// Login godoc
// @Summary User login
// @Description Authenticates a user with their credentials.
// @Tags user
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login Request"
// @Success 200 {object} response.Response "Success Response"
// @Failure 400 {object} response.ErrorResponse "Bad Request"
// @Failure 500 {object} response.ErrorResponse "Internal Server Error"
// @Router /user/login [post]
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



