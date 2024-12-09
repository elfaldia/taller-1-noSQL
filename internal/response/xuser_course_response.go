package response

type UserCourseResponse struct {
	IdUsuario  string `json:"id_usuario"`
	CourseName string `json:"course_name"`
	State      string `json:"state"`
}
