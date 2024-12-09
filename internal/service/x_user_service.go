package service

import (
	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(*request.RegisterUserRequest) (err error)
	FindAll() ([]response.UserResponse, error)
	FindById(string) (response.UserResponse, error)
	DeleteUsuario(string) (error)
	UpdateUser(*request.RegisterUserRequest, string) error
	LoginUser(*request.LoginRequest) (response.LoginResponse, error)
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}


func NewUserServiceImpl(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

// DeleteUsuario implements UserService.
func (u *UserServiceImpl) DeleteUsuario(userId string) error{
	return u.UserRepository.DeleteOne(userId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() (users []response.UserResponse, err error) {
	results, err := u.UserRepository.FindAll()
	if err != nil {
		return []response.UserResponse{}, err 
	}

	for _, value := range results {
		user := response.UserResponse{
			Id:                value.UserId,
			Nombre:            value.Nombre,
			Email: value.Email,
		}
		users = append(users, user)
	}
	return users, nil
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById( userId string) (response.UserResponse, error) {
	result, err := u.UserRepository.FindById((userId))
	if err != nil {
		return response.UserResponse{}, err
	}

	user := response.UserResponse {
		Id: result.UserId,
		Email: result.Email,
		Nombre: result.Nombre,
	}
	return user, nil
}

// LoginUser implements UserService.
func (u *UserServiceImpl) LoginUser(req *request.LoginRequest) (response.LoginResponse, error) {
	
	email := req.Email
	pass := req.Password

	user, err := u.UserRepository.FindById(email)

	if err != nil {
		return response.LoginResponse{}, err 
	}

	valid := u.CheckPassword(user.Clave, pass)

	if valid {
		return response.LoginResponse{
			Token: "a",
			Nombre: user.Nombre,
			Email: user.Email,
			Success: true,
		}, nil
	}

	return response.LoginResponse{
		Success: false,
	}, nil



}

// RegisterUser implements UserService.
func (u *UserServiceImpl) RegisterUser(userReq *request.RegisterUserRequest) (err error) {

	hashedPassword, err := u.HashPassword(userReq.Password)
	if err != nil {
		return err
	} 
	user := model.User{
		Email: userReq.Email,
		Nombre: userReq.Nombre,
		Clave: hashedPassword,
	}
	_, err = u.UserRepository.InsertOne(user)
	if err != nil {
		return err 
	}
	return nil
}

// UpdateUser implements UserService.
func (u *UserServiceImpl) UpdateUser( userReq *request.RegisterUserRequest, userId string) error {

	hashedPassword, err := u.HashPassword(userReq.Password)
	if err != nil {
		return err
	} 
	user := model.User{
		UserId: userId,
		Email: userReq.Email,
		Nombre: userReq.Nombre,
		Clave: hashedPassword,
	}
	_ , err = u.UserRepository.UpdateOne(user)
	if err != nil {
		return err 
	}
	return nil
}

func (u *UserServiceImpl) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *UserServiceImpl) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}