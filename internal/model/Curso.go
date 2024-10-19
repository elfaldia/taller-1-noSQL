package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Curso struct {
	Id               primitive.ObjectID `json:"_id" bson:"_id"`
	Nombre           string             `json:"nombre" bson:"nombre"`
	Descripcion      string             `json:"descripcion" bson:"descripcion"`
	Valoracion       float64            `json:"valoracion" bson:"valoracion"`
	ImagenMiniatura  string             `json:"imagen_miniatura" bson:"imagen_miniatura"`
	ImagenBanner     string             `json:"imagen_banner" bson:"imagen_banner"`
	CantidadUsuarios int                `json:"cantidad_usuarios" bson:"cantidad_usuarios"`
}
