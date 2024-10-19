package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
)

type CursoController struct {
	CursoService service.CursoService
}

func NewCursoController(service service.CursoService) *CursoController {
	return &CursoController{CursoService: service}
}

// @BasePath /curso
// @Summary Devuelve todos los carritos de la base de datos
// @Description get cursos
// @Tags curso
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /curso [get]
func (controller *CursoController) FindAll(ctx *gin.Context) {
	data, err := controller.CursoService.FindAll()

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

func (controller *CursoController) FindById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := controller.CursoService.FindById(id)
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

func (controller *CursoController) CreateCurso(ctx *gin.Context) {
	panic("falta")
}

func (controller *CursoController) CreateManyCurso(ctx *gin.Context) {
	panic("faltaaa")
}
