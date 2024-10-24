package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComentarioCurso struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre   string             `json:"nombre" bson:"nombre"`
	Fecha    string             `json:"fecha" bson:"fecha"`
	Titulo   string             `json:"titulo" bson:"titulo"`
	Detalle  string             `json:"detalle" bson:"detalle"`
	Likes    int                `json:"likes" bson:"likes"`
	Dislikes int                `json:"dislikes" bson:"dislikes"`
	IdCurso  primitive.ObjectID `json:"id_curso" bson:"id_curso"`
}
