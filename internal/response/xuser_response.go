package response

// Imports yatusae

type RegisterUserResponse struct {
	Id      string `json:"id"`
	Nombre  string `json:"nombre"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
