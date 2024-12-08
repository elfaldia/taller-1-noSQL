package request

// IMPORT A REDIS....

type UpdateCourseStatusRequest struct {
	IdCurso  string `json:"id_curso" validate:"required"`
	Estado   string `json:"estado" validate:"required,oneof=INICIADO EN_CURSO COMPLETADO"`
	Progreso int    `json:"progreso" validate:"required,min=0,max=100"`
}
