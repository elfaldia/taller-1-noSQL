package response

import (
	"github.com/elfaldia/taller-noSQL/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClaseReponse struct {
	Id                primitive.ObjectID `bson:"_id" json:"_id"`
	Nombre            string             `json:"nombre_clase" bson:"nombre_clase"`
	Indice            int                `json:"indice_clase" bson:"indice_clase"`
	Video             string             `json:"video" bson:"video"`
	Descripcion       string             `json:"descripcion" bson:"descripcion"`
	MaterialAdicional []model.Material    `json:"material_adicional" bson:"material_adicional"`
	IdUnidad          primitive.ObjectID `json:"id_unidad" bson:"id_unidad"`
}
