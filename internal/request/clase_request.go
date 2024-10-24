package request

import (
	"github.com/elfaldia/taller-noSQL/internal/model"
)

type CreateClaseRequest struct {
	Nombre            string           `json:"nombre" validate:"required"`
	Indice            int              `json:"indice_clase" validate:"required"`
	Video             string           `json:"video" validate:"required"`
	Descripcion       string           `json:"descripcion" bson:"descripcion"`
	MaterialAdicional []model.Material `json:"material_adicional" bson:"material_adicional"`
	IdUnidad          string           `json:"id_unidad" validate:"required"`
}

type CreateManyClaseRequest struct {
	Data []CreateClaseRequest `json:"data" validate:"required"`
}
