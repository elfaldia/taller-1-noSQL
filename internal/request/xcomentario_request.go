package request

// IMPORT A REDIS....

type AddCommentRequest struct {
	IdCurso    string `json:"id_curso" validate:"required" binding:"required"`
	IdUsuario  string `json:"id_usuario" validate:"required" binding:"required"`
	Titulo     string `json:"titulo" validate:"required" binding:"required"`
	Detalle    string `json:"detalle" validate:"required" binding:"required"`
	Valoracion int    `json:"valoracion" validate:"required,min=1,max=5" binding:"required"`
}
