package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResponseUnidad struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponseUnidad struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ObtenerUnidadResponde struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre  string             `json:"nombre" bson:"nombre"`
	Indice  int                `json:"indice" bson:"indice"`
	IdCurso primitive.ObjectID `json:"id_curso" bson:"id_curso"`
}
