package request

type CreateCursoRequest struct {
	Nombre           string `json:"nombre" validate:"required"`
	Descripcion      string `json:"descripcion" validate:"required"`
	ImagenMiniatura  string `json:"imagen_miniatura" validate:"required"`
	ImagenBanner     string `json:"imagen_banner" validate:"required"`
	CantidadUsuarios int    `json:"cantidad_usuarios" validate:"required"`
}

type CreateManyCursoRequest struct {
	Data []CreateCursoRequest `json:"data" validate:"required"`
}
