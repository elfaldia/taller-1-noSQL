package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComentarioCurso struct {
	ComentarioID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IdCurso      primitive.ObjectID `json:"id_curso" bson:"id_curso"`
	IdUsuario    primitive.ObjectID `json:"id_usuario" bson:"id_usuario"` //Nuevo
	Nombre       string             `json:"nombre" bson:"nombre"`
	Fecha        string             `json:"fecha" bson:"fecha"`
	Titulo       string             `json:"titulo" bson:"titulo"`
	Detalle      string             `json:"detalle" bson:"detalle"`
	Likes        int                `json:"likes" bson:"likes"`
	Dislikes     int                `json:"dislikes" bson:"dislikes"`
}
