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

// Nuevo router
func ClaseRouter(service *gin.Engine, cursoController *controller.CursoController) {
	router := service.Group("/clase")

	router.GET("/:id", cursoController.GetClaseById)
	router.GET("/:id/comentarios", cursoController.GetComentariosByClaseId)
	router.POST("", cursoController.CreateClase)
	router.POST("/:id/comentarios", cursoController.AddComentarioClase)
}
