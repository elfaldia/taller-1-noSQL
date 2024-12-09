package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
)

type UserCursoController struct {
	UserCursoService service.XUserCourseService
}

func NewUserCursoController(service service.XUserCourseService) *UserCursoController {
	return &UserCursoController{UserCursoService: service}
}

func (controller *UserCursoController) FindAll(ctx *gin.Context) {
	data, err := controller.UserCursoService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponseUnidad{
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

func (controller *UserCursoController) FindByIdUser(ctx *gin.Context) {
	id := ctx.Param("user_id")

	data, err := controller.UserCursoService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponseUnidad{
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

func (controller *UserCursoController) CreateOne(ctx *gin.Context) {
	var request request.AgregarCurso
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponseUnidad{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	err := controller.UserCursoService.AgregarCurso(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponseUnidad{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   "Curso agregado exitosamente",
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *UserCursoController) UpdateOne(ctx *gin.Context) {
	var request request.UpdateCurso
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponseUnidad{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	err := controller.UserCursoService.UpdateCurso(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponseUnidad{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   "Curso actualizado exitosamente",
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *UserCursoController) DeleteOne(ctx *gin.Context) {
	id := ctx.Param("user_id")
	cursoName := ctx.Param("curso_name")

	err := controller.UserCursoService.DeleteCurso(id, cursoName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponseUnidad{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.Response{
		Code:   200,
		Status: "OK",
		Data:   "Curso eliminado exitosamente",
	}

	ctx.JSON(http.StatusOK, res)
}
