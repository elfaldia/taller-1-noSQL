package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
)

type UnidadController struct {
	UnidadService service.UnidadService
}

func NewUnidadController(service service.UnidadService) *UnidadController {
	return &UnidadController{UnidadService: service}
}

func (controller *UnidadController) FindAll(ctx *gin.Context) {
	data, err := controller.UnidadService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponseUnidad{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	res := response.ResponseUnidad{
		Code:   200,
		Status: "OK",
		Data:   data,
	}

	ctx.JSON(http.StatusOK, res)
}

func (controller *UnidadController) FindByIdCurso(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := controller.UnidadService.FindByIdCurso(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponseUnidad{
			Code:    500,
			Message: err.Error(),
		})
	}

	res := response.ResponseUnidad{
		Code:   200,
		Status: "OK",
		Data:   data,
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *UnidadController) CreateOne(ctx *gin.Context) {
	var req request.CrearUnidadRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponseUnidad{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	unidad, err := controller.UnidadService.CreateOne(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponseUnidad{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	res := response.ResponseUnidad{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   unidad,
	}

	ctx.JSON(http.StatusCreated, res)
}
