package model

type Curso struct {
	Nombre       string  `json:"nombre_curso"`
	Descripcion  string  `json:"descripcion"`
	Valoracion   float32 `json:"valoracion"`
	Miniatura    string  `json:"miniatura"`
	ImagenBanner string  `json:"imagen_banner"`
	// mas cosas
}
