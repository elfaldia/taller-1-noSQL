package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
)

type ClaseController struct {
	ClaseService service.ClaseService
}

func NewClaseController(service service.ClaseService) *ClaseController {
	return &ClaseController{ClaseService: service}
}

// @BasePath /unidad
// @Summary get clases por unidad
// @Description Devuelve todas las clases que tiene una unidad
// @Tags unidad
// @Param unidad_id path string true "UNIDAD OBJECT ID"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /unidad/{unidad_id}/clases [get]
func (controller *ClaseController) FindAllByIdUnidad(ctx *gin.Context) {

	idUnidad := ctx.Param("unidad_id")

	if idUnidad == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "unidad_id es obligatorio",
		})
	}

	data, err := controller.ClaseService.FindAllByIdUnidad(idUnidad)

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

// @BasePath /clase
// @Summary get clase por Object ID
// @Description Devuelve una clase
// @Tags clase
// @Param id path string true "CLASE ID"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /clase/{id} [get]
func (controller *ClaseController) FindById(ctx *gin.Context) {
	id := ctx.Param("clase_id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    400,
			Message: "clase_id es obligatorio",
		})
	}

	data, err := controller.ClaseService.FindById(id)
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

func (controller *ClaseController) CreateClase(ctx *gin.Context) {
	var req request.CreateClaseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	_, err := controller.ClaseService.CreateClase(req)
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

func (controller *ClaseController) CreateManyClase(ctx *gin.Context) {

	var req request.CreateManyClaseRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err := controller.ClaseService.CreateManyClase(req.Data)
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
