package request

type CreateComentarioClase struct {
	Nombre   string `json:"nombre" bson:"nombre" validate:"required"`
	Fecha    string `json:"fecha" bson:"fecha" validate:"required"`
	Titulo   string `json:"titulo" bson:"titulo" validate:"required"`
	Detalle  string `json:"detalle" bson:"detalle" validate:"required"`
	IdClase  string `json:"id_clase" bson:"id_clase" validate:"required"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}
