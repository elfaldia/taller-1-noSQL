package main

import (
	"github.com/elfaldia/taller-noSQL/internal/controller"
	"github.com/gin-gonic/gin"
)

func CursoRouter(service *gin.Engine, cursoController *controller.CursoController, unidadController *controller.UnidadController) {
	router := service.Group("/curso")

	router.GET("", cursoController.FindAll)
	router.GET("/:curso_id", cursoController.FindById)
	router.POST("", cursoController.CreateCurso)

	router.GET("/:curso_id/comentarios", cursoController.GetComentariosByCursoId)
	router.POST("/:curso_id/comentarios", cursoController.AddComentarioCurso)

	router.GET("/:curso_id/unidad", unidadController.FindByIdCurso)

	router.POST("ruta-para-insertar-muchos-cursos", cursoController.CreateManyCurso)

}

func UnidadRouter(service *gin.Engine, unidadController *controller.UnidadController, claseController *controller.ClaseController) {
	router := service.Group("/unidad")

	router.GET("", unidadController.FindAll)
	router.POST("", unidadController.CreateOne)

	router.GET("/:unidad_id/clase", claseController.FindAllByIdUnidad)

}

func ClaseRouter(service *gin.Engine, claseController *controller.ClaseController, comentarioClaseController *controller.ComentarioClaseController) {
	router := service.Group("/clase")

	router.GET("/:clase_id", claseController.FindById)
	router.POST("", claseController.CreateClase)
	router.POST("/post-many", claseController.CreateManyClase)

	router.POST("/comentario", comentarioClaseController.CreateComentarioClase)
	router.GET("/:clase_id/comentario", comentarioClaseController.FindAllByIdClase)
	router.GET("/comentario/:comentario_id", comentarioClaseController.FindById)
}
