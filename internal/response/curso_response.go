package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CursoReponse struct {
	Id               primitive.ObjectID `json:"_id" bson:"_id"`
	Nombre           string             `json:"nombre"`
	Descripcion      string             `json:"descripcion"`
	Valoracion       float64            `json:"valoracion"`
	ImagenMiniatura  string             `json:"imagen_miniatura"`
	ImagenBanner     string             `json:"imagen_banner"`
	CantidadUsuarios int                `json:"cantidad_usuarios"`
}
