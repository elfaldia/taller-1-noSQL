package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/elfaldia/taller-noSQL/internal/db"
	"github.com/gin-gonic/gin"
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
