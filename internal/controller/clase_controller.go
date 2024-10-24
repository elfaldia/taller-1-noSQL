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
// @Tags clase
// @Param id_unidad path string true "UNIDAD OBJECT ID"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /unidad/{id_unidad}/clase [get]
func (controller *ClaseController) FindAllByIdUnidad(ctx *gin.Context) {

	idUnidad := ctx.Param("id_unidad")

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

// @BasePath /unidad
// @Summary get clase por Object ID
// @Description Devuelve una clase
// @Tags clase
// @Param id path string true "CLASE ID"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /unidad/{id_unidad}/clase [get]
func (controller *ClaseController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")
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

// @BasePath /unidad
// @Summary Crea una clase
// @Description Agrega una clase a la coleccion Clase
// @Tags clase
// @Param clase body request.CreateClaseRequest true "Carrito a crear"
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /unidad/{id_unidad}/clase [post]
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
