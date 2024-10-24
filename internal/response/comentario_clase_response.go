package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type ComentarioClaseResponse struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Nombre   string             `json:"nombre" bson:"nombre"`
	Fecha    string             `json:"fecha" bson:"fecha"`
	Titulo   string             `json:"titulo" bson:"titulo"`
	Detalle  string             `json:"detalle" bson:"detalle"`
	Likes    int                `json:"likes" bson:"likes"`
	Dislikes int                `json:"dislikes" bson:"dislikes"`
	IdClase  primitive.ObjectID `json:"id_clase" bson:"id_clase"`
}
