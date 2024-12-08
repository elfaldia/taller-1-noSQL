package response

// Imports yatusae

type AddCommentResponse struct {
	IdComentario string `json:"id_comentario"`
	IdCurso      string `json:"id_curso"`
	IdUsuario    string `json:"id_usuario"`
	Titulo       string `json:"titulo"`
	Detalle      string `json:"detalle"`
	Valoracion   int    `json:"valoracion"`
	Message      string `json:"message"`
}
