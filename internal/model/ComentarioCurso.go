package model

type ComentarioCurso struct {
	ComentarioID string `bson:"_id,omitempty" json:"id,omitempty"`
	IdCurso      string `json:"id_curso" bson:"id_curso"`
	IdUsuario    string `json:"id_usuario" bson:"id_usuario"` //Nuevo
	Nombre       string `json:"nombre" bson:"nombre"`
	Fecha        string `json:"fecha" bson:"fecha"`
	Titulo       string `json:"titulo" bson:"titulo"`
	Detalle      string `json:"detalle" bson:"detalle"`
	Likes        int    `json:"likes" bson:"likes"`
	Dislikes     int    `json:"dislikes" bson:"dislikes"`
}
