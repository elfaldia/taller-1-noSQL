package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Unidad struct {
	Id     primitive.ObjectID `json:"_id" bson:"_id"`
	Nombre string             `json:"nombre" bson:"nombre"`
	Indice int                `json:"indice" bson:"indice"` // Validar que debe ser unico
	Clases []Clase            `json:"clases" bson:"clases"`
}
