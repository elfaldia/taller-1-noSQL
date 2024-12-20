package main

import (
	"github.com/elfaldia/taller-noSQL/internal/controller"
	"github.com/gin-gonic/gin"
)

func CursoRouter(service *gin.Engine, cursoController *controller.CursoController, unidadController *controller.UnidadController, comentarioController *controller.ComentarioController) {
	router := service.Group("/curso")

	router.GET("", cursoController.FindAll)
	router.GET("/:curso_id", cursoController.FindById)
	router.POST("", cursoController.CreateCurso)
	// router.GET("/:curso_id/comentarios", cursoController.GetComentariosByCursoId) // *
	// router.POST("/:curso_id/comentarios", cursoController.AddComentarioCurso) // *
	router.GET("/:curso_id/unidades", unidadController.FindByIdCurso)

	router.POST("/comentarios", comentarioController.AddComentario)        // *
	router.GET("/:curso_id/comentarios", comentarioController.GetComentariosByCurso) // *
}

func UnidadRouter(service *gin.Engine, unidadController *controller.UnidadController, claseController *controller.ClaseController) {
	router := service.Group("/unidad")

	router.GET("", unidadController.FindAll)
	router.GET("/:unidad_id/clases", claseController.FindAllByIdUnidad)

}

func ClaseRouter(service *gin.Engine, claseController *controller.ClaseController, comentarioClaseController *controller.ComentarioClaseController) {
	router := service.Group("/clase")

	router.GET("/:clase_id", claseController.FindById)
	router.POST("/comentarios", comentarioClaseController.CreateComentarioClase)
	router.GET("/:clase_id/comentarios", comentarioClaseController.FindAllByIdClase)
	router.GET("/comentarios/:comentario_id", comentarioClaseController.FindById)
}

func UserCursoRouter(service *gin.Engine, userCursoController *controller.UserCursoController) {
	router := service.Group("/user_curso")

	router.GET("", userCursoController.FindAll)
	router.GET("/:user_id", userCursoController.FindByIdUser)
	router.POST("", userCursoController.CreateOne)
	router.DELETE("/:user_id/:curso_name", userCursoController.DeleteOne)
	router.PATCH("", userCursoController.UpdateOne)
	router.POST("/:user_id/:curso_name", userCursoController.AddCourseRating)
	router.GET("/avg/:curso_name", userCursoController.GetCourseRating)
}

func UserRouter(service *gin.Engine, userController *controller.UserController) {
	router := service.Group("/user")

	router.GET("", userController.FindAll)
	router.GET("/:user_id", userController.FindById)
	router.POST("", userController.CreateUser)
	router.POST("/login", userController.Login)
	router.DELETE("/:user_id", userController.DeleteUser)

}
