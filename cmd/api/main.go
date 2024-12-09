package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
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

	fmt.Printf("aws_access_key_id: %s\n", os.Getenv("aws_access_key_id"))
	fmt.Printf("aws_secret_access_key: %s\n", os.Getenv("aws_secret_access_key"))
	fmt.Printf("region: %s\n", os.Getenv("region"))

	clientDB := db.ConnectDynamoDB()
	log.Println("Cliente DynamoDB inicializado:", clientDB)

	userRepository := repository.NewUserRepositoryImpl(clientDB)
	_ = service.NewUserServiceImpl(userRepository)

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

	cursoRepository := repository.NewCursoRepositoryImpl(cursoCollection)
	unidadRepository := repository.NewUnidadRepositoryImpl(unidadCollection)
	claseRepository := repository.NewClaseRepositoryImpl(claseCollection)
	comentarioClaseRepository := repository.NewComentarioClaseRepositoryImpl(comentarioClaseCollection)

	unidadService, _ := service.NewUnidadServiceImpl(unidadRepository, validate)
	unidadController := controller.NewUnidadController(unidadService)

	claseService, _ := service.NewClaseServiceImpl(claseRepository, validate)
	claseController := controller.NewClaseController(claseService)

	comentarioClaseService, _ := service.NewComentarioClaseServiceImpl(comentarioClaseRepository, claseService, validate)
	comentarioClaseController := controller.NewComentarioClaseController(comentarioClaseService)

	cursoService, _ := service.NewCursoServiceImpl(cursoRepository, validate, db, unidadService, claseService)
	cursoController := controller.NewCursoController(cursoService, comentarioClaseService, claseService)

	cursoUsuarioRepositorio := repository.NewCursoUsuarioRepositoryImpl(clientDB)
	cursoUsuarioService := service.NewXUserCourseServiceImpl(cursoUsuarioRepositorio)
	cursoUsuarioController := controller.NewUserCursoController(cursoUsuarioService)

	routes := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	routes.GET("ruta-para-rellenar-base", cursoController.RellenarBase)
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	CursoRouter(routes, cursoController, unidadController)
	UnidadRouter(routes, unidadController, claseController)
	ClaseRouter(routes, claseController, comentarioClaseController)
	UserCursoRouter(routes, cursoUsuarioController)

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
