package main

import (
	"github.com/elfaldia/taller-noSQL/internal/controller"
	"github.com/gin-gonic/gin"
)

func CursoRouter(service *gin.Engine, cursoController *controller.CursoController) {
	router := service.Group("/curso")

	router.GET("", cursoController.FindAll)
	// NECESITO QUE ALGUIEN VEA ESTA TONTERA JAJAJAJAJA (curso_id != id)
	router.GET("/:curso_id", cursoController.FindById)
	router.GET("/:curso_id/comentarios", cursoController.GetComentariosByCursoId)
	router.POST("", cursoController.CreateCurso)
	router.POST("ruta-para-insertar-muchos-cursos", cursoController.CreateManyCurso)
	router.POST("/:curso_id/comentarios", cursoController.AddComentarioCurso)

}

func UnidadRouter(service *gin.Engine, unidadController *controller.UnidadController) {
	router := service.Group("/unidad")

	router.GET("", unidadController.FindAll)
	router.GET("/:id", unidadController.FindByIdCurso)
	router.POST("", unidadController.CreateOne)

}
