package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CursoController struct {
	CursoService           service.CursoService
	ComentarioClaseService service.ComentarioClaseService
	ClaseService           service.ClaseService
}

func NewCursoController(
	service service.CursoService,
	comentarioClaseService service.ComentarioClaseService,
	claseService service.ClaseService,
) *CursoController {
	return &CursoController{
		CursoService:           service,
		ComentarioClaseService: comentarioClaseService,
		ClaseService:           claseService,
	}
}

// @BasePath /curso
// @Summary Devuelve todos los cursos de la base de datos
// @Description get cursos
// @Tags curso
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
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

// @BasePath /curso
// @Summary Devuelve un curso
// @Description get curso a partir del ID
// @Tags curso
// @Param curso_id path int true
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Router /curso/{curso_id} [get]
func (controller *CursoController) FindById(ctx *gin.Context) {
	id := ctx.Param("curso_id")

	data, err := controller.CursoService.FindById(id)
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

// @BasePath /curso
// @Summary Crea un curso
// @Description crea un curso con todos sus componentes
// @Tags curso
// @Param  curso body request.CreateCursoRequest true "json del curso"
// @Accept json
// @Produce json
// @Success 201 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /curso [post]
func (controller *CursoController) CreateCurso(ctx *gin.Context) {
	var req *request.CreateCursoRequest

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
	_, err := controller.CursoService.CreateCurso(req)
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

// @BasePath /curso
// @Summary Agrega comentario a un curso
// @Param curso_id path string true "id del curso"
// @Description add comentarios
// @Tags curso
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /curso/{curso_id}/comentarios [post]
func (controller *CursoController) AddComentarioCurso(ctx *gin.Context) {
	cursoID := ctx.Param("curso_id")

	var comentario model.ComentarioCurso
	if err := ctx.ShouldBindJSON(&comentario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	objectIdCurso, err := primitive.ObjectIDFromHex(cursoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid curso ID"})
		return
	}

	comentario.IdCurso = objectIdCurso

	if err := controller.CursoService.AddComentarioCurso(comentario); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":    "Comentario añadido con éxito",
		"comentario": comentario,
	})
}

// @BasePath /curso
// @Summary Obtiene comentarios de un curso
// @Param curso_id path string true "id del curso"
// @Description get comentarios
// @Tags curso
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /curso/{curso_id}/comentarios [get]
func (controller *CursoController) GetComentariosByCursoId(ctx *gin.Context) {
	cursoID := ctx.Param("curso_id") // Extraer el ID del curso desde la URL

	// Convertir el cursoID de string a ObjectID
	objectIdCurso, err := primitive.ObjectIDFromHex(cursoID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid curso ID"})
		return
	}

	comentarios, err := controller.CursoService.GetComentariosByCursoId(objectIdCurso)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Code:   200,
		Status: "success",
		Data:   comentarios,
	})
}


// @BasePath
// @Summary Crea cursos con todos sus componentes (rellena la base)
// @Description crear cursos
// @Tags curso
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.ErrorResponse
// @Router /ruta-para-rellenar-base [get]
func (c *CursoController) RellenarBase(ctx *gin.Context) {

	cursosFile, err := os.Open("data_cursos.json")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	defer cursosFile.Close()
	cursosBytes, _ := io.ReadAll(cursosFile)

	var cursos []request.CreateCursoRequest
	if err := json.Unmarshal(cursosBytes, &cursos); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	for _, cursoReq := range cursos {
		_, err = c.CursoService.CreateCurso(&cursoReq)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}

	comentariosCursoFile, err := os.Open("data_comentarios_cursos.json")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	defer comentariosCursoFile.Close()
	comentarioCursoByte, _ := io.ReadAll(comentariosCursoFile)

	var comentarios []request.CreateComentarioRequest
	if err := json.Unmarshal(comentarioCursoByte, &comentarios); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	for _, comentarioCursoReq := range comentarios {
		idCursoRandom, err := c.CursoService.GetRandomId()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}

		comentarioCurso := model.ComentarioCurso{
			IdCurso:  idCursoRandom,
			Nombre:   comentarioCursoReq.Nombre,
			Likes:    comentarioCursoReq.Likes,
			Dislikes: comentarioCursoReq.Dislikes,
			Fecha:    comentarioCursoReq.Fecha,
			Titulo:   comentarioCursoReq.Titulo,
			Detalle:  comentarioCursoReq.Detalle,
		}

		err = c.CursoService.AddComentarioCurso(comentarioCurso)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
	}

	comentariosClaseFile, err := os.Open("data_com_clases.json")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	defer comentariosCursoFile.Close()
	comentarioClaseByte, _ := io.ReadAll(comentariosClaseFile)

	var comentariosClase []request.CreateComentarioClase
	if err := json.Unmarshal(comentarioClaseByte, &comentariosClase); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	for _, comentarioClaseReq := range comentariosClase {
		idClase, err := c.ClaseService.GetRandomId()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
			return
		}
		comentarioClaseReq.IdClase = idClase.Hex()
		_, err = c.ComentarioClaseService.CreateComentarioClase(comentarioClaseReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid error al ingresar un comentario clase"})
			return
		}
	}

	res := response.Response{
		Code:   http.StatusCreated,
		Status: "Created",
		Data:   nil,
	}
	ctx.JSON(http.StatusCreated, res)

}
