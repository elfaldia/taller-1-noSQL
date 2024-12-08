package response

// Imports yatusae

type LoginResponse struct {
	Token   string `json:"token"`
	Nombre  string `json:"nombre"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
