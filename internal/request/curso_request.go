package request

import (
	"github.com/elfaldia/taller-noSQL/internal/model"
)

type CreateCursoRequest struct {
	Nombre          string          `json:"nombre" validate:"required" binding:"required"`
	Descripcion     string          `json:"descripcion" validate:"required" binding:"required"`
	ImagenMiniatura string          `json:"imagen_miniatura" validate:"required" binding:"required"`
	ImagenBanner    string          `json:"imagen_banner" validate:"required" binding:"required"`
	Unidades        []UnidadRequest `json:"unidades" validate:"required"`
	Valoracion      float64         `json:"valoracion" validate:"gte=0,lte=5.0"`
	CantidadUsuario int             `json:"cantidad_usuarios" validate:"gte=0"`
}

type UnidadRequest struct {
	NombreUnidad string         `json:"nombre_unidad" validate:"required" binding:"required"`
	IndiceUnidad int            `json:"indice_unidad" validate:"required" binding:"required"`
	Clases       []ClaseRequest `json:"clases" validate:"required" binding:"required"`
}

type ClaseRequest struct {
	NombreClase       string           `json:"nombre_clase" validate:"required" binding:"required"`
	IndiceClase       int              `json:"indice_clase" validate:"required" binding:"required"`
	Video             string           `json:"video" validate:"required" binding:"required"`
	Descripcion       string           `json:"descripcion" validate:"required" binding:"required"`
	MaterialAdicional []model.Material `json:"material_adicional"`
}

type CreateManyCursoRequest struct {
	Data []CreateCursoRequest `json:"data" validate:"required"`
}

type CreateComentarioRequest struct {
	Nombre   string `json:"nombre" validate:"required"`
	Fecha    string `json:"fecha" validate:"required"`
	Titulo   string `json:"titulo" validate:"required"`
	Detalle  string `json:"detalle" validate:"required"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes" `
	IdCurso  string `json:"id_curso" validate:"required"`
}
