package controller

import "github.com/elfaldia/taller-noSQL/internal/service"

type UserController struct {
	UserService service.UserService
}


func NewUserController(service service.UserService) *UserController {
	return &UserController{UserService: service}
}



