package request

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateCursoRequest struct {
	Nombre           string  `json:"nombre" validate:"required"`
	Descripcion      string  `json:"descripcion" validate:"required"`
	Valoracion       float64 `json:"valoracion" validate:"required"`
	ImagenMiniatura  string  `json:"imagen_miniatura" validate:"required"`
	ImagenBanner     string  `json:"imagen_banner" validate:"required"`
	CantidadUsuarios int     `json:"cantidad_usuarios" validate:"required"`
}

type CreateManyCursoRequest struct {
	Data []CreateCursoRequest `json:"data" validate:"required"`
}

type CreateComentarioRequest struct {
	Nombre   string             `json:"nombre" validate:"required"`
	Fecha    string             `json:"fecha" validate:"required"`
	Titulo   string             `json:"titulo" validate:"required"`
	Detalle  string             `json:"detalle" validate:"required"`
	Likes    int                `json:"likes" validate:"required"`
	Dislikes int                `json:"dislikes" validate:"required"`
	IdCurso  primitive.ObjectID `json:"id_curso" validate:"required"`
}
