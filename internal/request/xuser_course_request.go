package request

type AgregarCurso struct {
	UserId     string `json:"user_id"`
	CourseName string `json:"course_name"`
	State      string `json:"state" `
}

type UpdateCurso struct {
	UserId       string `json:"user_id"`
	CourseName   string `json:"course_name"`
	State        string `json:"state" `
	ClasesVistas int    `json:"clases_vistas"`
}

type AgregarRating struct {
	Rating int `json:"rating"`
}
