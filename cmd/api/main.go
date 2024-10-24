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
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
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

	cursoRepository := repository.NewCursoRepositoryImpl(cursoCollection)
	unidadRepository := repository.NewUnidadRepositoryImpl(unidadCollection)

	cursoService, _ := service.NewCursoServiceImpl(cursoRepository, validate)
	unidadService, _ := service.NewUnidadServiceImpl(unidadRepository, validate)

	cursoController := controller.NewCursoController(cursoService)
	unidadController := controller.NewUnidadController(unidadService)

	routes := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	CursoRouter(routes, cursoController)
	UnidadRouter(routes, unidadController)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
