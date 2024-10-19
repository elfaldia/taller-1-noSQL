package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/elfaldia/taller-noSQL/internal/controller"
	"github.com/elfaldia/taller-noSQL/internal/db"
	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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

	cursoRepository := repository.NewCursoRepositoryImpl(cursoCollection)

	cursoService, _ := service.NewCursoServiceImpl(cursoRepository, validate)

	cursoController := controller.NewCursoController(cursoService)

	routes := gin.Default()
	routes.GET("/", func(ctx *gin.Context) {
		res := struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    200,
			Message: "Hola Mundo",
		}
		ctx.JSON(http.StatusOK, res)
	})
	CursoRouter(routes, cursoController)

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
