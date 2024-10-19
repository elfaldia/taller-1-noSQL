package model


type Clase struct {
	Nombre            string             `json:"nombre_clase" bson:"nombre_clase"`
	Indice            int                `json:"indice_clase" bson:"indice_clase"`
	Video             string             `json:"video" bson:"video"`
	Descripcion       string             `json:"descripcion" bson:"descripcion"`
	MaterialAdicional []Material         `json:"material_adicional" bson:"material_adicional"`
}

type Material struct {
	URL    string `json:"url" bson:"url"`
	Nombre string `json:"nombre" bson:"nombre"`
	Tipo   string `json:"tipo" bson:"tipo"` // Puede ser ENUM
}
