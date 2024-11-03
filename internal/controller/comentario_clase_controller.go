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


// @BasePath /clase
// @Summary Devuelve todos los comentarios de una clase
// @Description todos los comentarios de una clase 
// @Tags clase
// @Param clase_id path int true
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /clase/{clase_id} [get]
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


// @BasePath /clase
// @Summary Devuelve una comentario clase
// @Description comentario clase
// @Tags clase
// @Param clase_id path int true
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /curso/{clase_id} [get]
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

// @BasePath /clase
// @Summary Crea un comentario
// @Description crea un comentario para una clase
// @Tags clase
// @Param  cometario body request.CreateComentarioClase true "json del comentario"
// @Accept json
// @Produce json
// @Success 201 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /clase/comentarios [post]
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
