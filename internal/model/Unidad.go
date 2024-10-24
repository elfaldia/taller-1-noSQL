package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Unidad struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre  string             `json:"nombre" bson:"nombre"`
	Indice  int                `json:"indice" bson:"indice"` // Validar que debe ser unico
	IdCurso primitive.ObjectID `json:"id_curso" bson:"id_curso"`
}
