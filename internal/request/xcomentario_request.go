package request

// IMPORT A REDIS....

type AddCommentRequest struct {
	IdCurso    string `json:"id_curso" validate:"required"`
	IdUsuario  string `json:"id_usuario" validate:"required"`
	Titulo     string `json:"titulo" validate:"required"`
	Detalle    string `json:"detalle" validate:"required"`
	Valoracion int    `json:"valoracion" validate:"required,min=1,max=5"`
}
