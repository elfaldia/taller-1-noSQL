package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var req request.CreateCursoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err := controller.CursoService.CreateCurso(req)
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

func (controller *CursoController) CreateManyCurso(ctx *gin.Context) {

	var req request.CreateManyCursoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err := controller.CursoService.CreateManyCursos(req)
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
func (controller *CursoController) AddComentarioCurso(ctx *gin.Context) {
	cursoID := ctx.Param("curso_id") // Extraer ID del curso desde la URL

	var comentario model.ComentarioCurso
	if err := ctx.ShouldBindJSON(&comentario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convertir el cursoID de string a ObjectID
	objectIdCurso, err := primitive.ObjectIDFromHex(cursoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid curso ID"})
		return
	}

	comentario.IdCurso = objectIdCurso // Asociar comentario al curso

	// Lógica para guardar el comentario en la base de datos
	if err := controller.CursoService.AddComentarioCurso(comentario); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusCreated, gin.H{
		"message":    "Comentario añadido con éxito",
		"comentario": comentario,
	})
}

func (controller *CursoController) GetComentariosByCursoId(ctx *gin.Context) {
	cursoID := ctx.Param("curso_id") // Extraer el ID del curso desde la URL

	// Convertir el cursoID de string a ObjectID
	objectIdCurso, err := primitive.ObjectIDFromHex(cursoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid curso ID"})
		return
	}

	// Obtener los comentarios del curso desde el servicio
	comentarios, err := controller.CursoService.GetComentariosByCursoId(objectIdCurso)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respuesta exitosa
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   comentarios,
	})
}

// CAMBIAR A OTRO FILE DE CLASE_CONTROLLER.GO
func (controller *CursoController) CreateClase(ctx *gin.Context) {
	var clase model.Clase

	if err := ctx.ShouldBindJSON(&clase); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.CursoService.AddClase(clase)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Clase añadida con éxito",
		"clase":   clase,
	})
}

func (controller *CursoController) GetClaseById(ctx *gin.Context) {
	claseID := ctx.Param("id")

	objectIdClase, err := primitive.ObjectIDFromHex(claseID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid clase ID"})
		return
	}

	clase, err := controller.CursoService.GetClaseById(objectIdClase)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Clase obtenida con éxito",
		"clase":   clase,
	})
}
