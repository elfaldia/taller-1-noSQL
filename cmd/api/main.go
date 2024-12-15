package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/elfaldia/taller-noSQL/docs"
	"github.com/elfaldia/taller-noSQL/internal/controller"
	"github.com/elfaldia/taller-noSQL/internal/db"
	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	err1 := godotenv.Load(".envrc")
	if err1 != nil {
		log.Fatalf("Error cargando archivo .envrc: %v", err1)
	}

	driverNeo4j := db.ConnectNeo4jDB()
	log.Println("Cliente Neo4j inicializado correctamente: ", driverNeo4j)

	clientDB := db.ConnectDynamoDB()
	log.Println("Cliente DynamoDB inicializado:", clientDB)

	client, err := db.ConnectToDataBase()
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error desconectando de MongoDB: %v", err)
		}
	}()

	validate := validator.New()
	db := client.Database("taller-nosql")
	cursoCollection := db.Collection("curso")
	unidadCollection := db.Collection("unidad")
	claseCollection := db.Collection("clases")
	comentarioClaseCollection := db.Collection("comentarios_clase")

	userRepository := repository.NewUserRepositoryImpl(clientDB)
	userService := service.NewUserServiceImpl(userRepository)
	userController := controller.NewUserController(userService)

	cursoRepository := repository.NewCursoRepositoryImpl(cursoCollection)
	unidadRepository := repository.NewUnidadRepositoryImpl(unidadCollection)
	claseRepository := repository.NewClaseRepositoryImpl(claseCollection)
	comentarioClaseRepository := repository.NewComentarioClaseRepositoryImpl(comentarioClaseCollection)

	unidadService, _ := service.NewUnidadServiceImpl(unidadRepository, validate)
	unidadController := controller.NewUnidadController(unidadService)

	ComentarioRepository := repository.NewComentarioRepositoryImpl(*driverNeo4j) //añadir Neoj4

	claseService, _ := service.NewClaseServiceImpl(claseRepository, validate)
	claseController := controller.NewClaseController(claseService)

	comentarioClaseService, _ := service.NewComentarioClaseServiceImpl(comentarioClaseRepository, claseService, validate)
	comentarioClaseController := controller.NewComentarioClaseController(comentarioClaseService)

	cursoService, _ := service.NewCursoServiceImpl(cursoRepository, ComentarioRepository, validate, db, unidadService, claseService)
	cursoController := controller.NewCursoController(cursoService, comentarioClaseService, claseService)
	comentarioController := controller.NewComentarioController(cursoService)

	cursoUsuarioRepositorio := repository.NewCursoUsuarioRepositoryImpl(clientDB, driverNeo4j)
	cursoUsuarioService := service.NewXUserCourseServiceImpl(cursoUsuarioRepositorio, userService, cursoService)
	cursoUsuarioController := controller.NewUserCursoController(cursoUsuarioService)

	routes := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	routes.GET("ruta-para-rellenar-base", cursoController.RellenarBase)
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	CursoRouter(routes, cursoController, unidadController, comentarioController)
	UnidadRouter(routes, unidadController, claseController)
	ClaseRouter(routes, claseController, comentarioClaseController)
	UserCursoRouter(routes, cursoUsuarioController)
	UserRouter(routes, userController)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
