package main

import (
	"github.com/elfaldia/taller-noSQL/internal/controller"
	"github.com/gin-gonic/gin"
)

func CursoRouter(service *gin.Engine, cursoController *controller.CursoController) {
	router := service.Group("/curso")

	router.GET("", cursoController.FindAll)
	router.GET("/:id", cursoController.FindById)
	router.POST("", cursoController.CreateCurso)
	router.POST("ruta-para-insertar-muchos-cursos", cursoController.CreateManyCurso)
}
