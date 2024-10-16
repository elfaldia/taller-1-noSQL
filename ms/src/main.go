package main

import (
	"context"
	"log"

	"github.com/ms/src/connection"
)

func main() {
	client, err := connection.ConnectToDataBase()
	if err != nil {
		log.Fatalf("Error conectando a MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error desconectando de MongoDB: %v", err)
		}
	}()
}
