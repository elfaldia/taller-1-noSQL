package request

type CrearUnidadRequest struct {
	Nombre  string `json:"nombre" bson:"nombre"`
	Indice  int    `json:"indice" bson:"indice"`
	IdCurso string `json:"id_curso"`
}

type CrearUnidadesRequest struct {
	Data []CrearUnidadRequest `json:"data" validate:"required"`
}
