package model

type ComentarioCurso struct {
	ComentarioID string `bson:"_id,omitempty" json:"id,omitempty"`
	IdCurso      string `json:"id_curso" bson:"id_curso" binding:"required"`
	IdUsuario    string `json:"id_usuario" bson:"id_usuario" binding:"required"` //Nuevo
	Nombre       string `json:"nombre" bson:"nombre" binding:"required"`
	Fecha        string `json:"fecha" bson:"fecha" binding:"required"`
	Titulo       string `json:"titulo" bson:"titulo" binding:"required"`
	Detalle      string `json:"detalle" bson:"detalle" binding:"required"`
	Likes        int    `json:"likes" bson:"likes"`
	Dislikes     int    `json:"dislikes" bson:"dislikes"`
}
