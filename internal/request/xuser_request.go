package request

// IMPORT A REDIS....

type RegisterUserRequest struct {
	Nombre   string `json:"nombre" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type GetUserCoursesRequest struct {
	IdUsuario string `json:"id_usuario" validate:"required"`
}
