package request

type AgregarCurso struct {
	UserId     string `json:"user_id"`
	CourseName string `json:"course_name"`
	State      string `json:"state" `
}

type UpdateCurso struct {
	State string `json:"state" `
}
