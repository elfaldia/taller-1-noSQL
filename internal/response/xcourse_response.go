package response

// Imports yatusae

type UpdateCourseStatusResponse struct {
	IdCurso  string `json:"id_curso"`
	Estado   string `json:"estado"`
	Progreso int    `json:"progreso"`
	Message  string `json:"message"`
}
