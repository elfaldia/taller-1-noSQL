package request

// IMPORT A REDIS....

type RegisterUserRequest struct {
	Nombre   string `json:"nombre" validate:"required" binding:"required"`
	Email    string `json:"email" validate:"required,email" binding:"required"`
	Password string `json:"password" validate:"required,min=5" binding:"required"`
}

type GetUserCoursesRequest struct {
	IdUsuario string `json:"id_usuario" validate:"required"`
}
