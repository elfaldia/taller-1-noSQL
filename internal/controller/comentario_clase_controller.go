package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
)

type ComentarioClaseController struct {
	ComentarioClaseService service.ComentarioClaseService
}

func NewComentarioClaseController(service service.ComentarioClaseService) *ComentarioClaseController {
	return &ComentarioClaseController{ComentarioClaseService: service}
}

func (controller *ComentarioClaseController) FindAllByIdClase(ctx *gin.Context) {

	idClase := ctx.Param("clase_id")

	data, err := controller.ComentarioClaseService.FindAllByIdClase(idClase)

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

func (controller *ComentarioClaseController) FindById(ctx *gin.Context) {
	id := ctx.Param("comentario_id")
	data, err := controller.ComentarioClaseService.FindById(id)
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

func (controller *ComentarioClaseController) CreateComentarioClase(ctx *gin.Context) {
	var req request.CreateComentarioClase

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	_, err := controller.ComentarioClaseService.CreateComentarioClase(req)
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
