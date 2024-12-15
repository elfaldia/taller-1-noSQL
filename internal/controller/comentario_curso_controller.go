package controller

import (
	"net/http"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
)

type ComentarioController struct {
	ComentarioService service.CursoService
}

func NewComentarioController(service service.CursoService) *ComentarioController {
	return &ComentarioController{ComentarioService: service}
}

func (c *ComentarioController) AddComentario(ctx *gin.Context) {
	var comentario model.ComentarioCurso
	if err := ctx.ShouldBindJSON(&comentario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.ComentarioService.AddComentarioCurso(comentario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Comentario creado"})
}

func (c *ComentarioController) GetComentariosByCurso(ctx *gin.Context) {
	cursoID := ctx.Param("curso_id")
	comentarios, err := c.ComentarioService.GetComentariosByCursoId(cursoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comentarios)
}
