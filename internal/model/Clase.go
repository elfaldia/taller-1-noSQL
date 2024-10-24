package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Clase struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Nombre            string             `json:"nombre_clase" bson:"nombre_clase"`
	Indice            int                `json:"indice_clase" bson:"indice_clase"`
	Video             string             `json:"video" bson:"video"`
	Descripcion       string             `json:"descripcion" bson:"descripcion"`
	MaterialAdicional []Material         `json:"material_adicional" bson:"material_adicional"`
	IdUnidad          primitive.ObjectID `json:"id_unidad" bson:"id_unidad"`
}

type Material struct {
	URL    string `json:"url" bson:"url"`
	Nombre string `json:"nombre" bson:"nombre"`
	Tipo   string `json:"tipo" bson:"tipo"` // Puede ser ENUM
}
