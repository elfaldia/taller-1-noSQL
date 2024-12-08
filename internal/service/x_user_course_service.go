package service

import (
	"fmt"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
)

type XUserCourseService interface {
	AgregarCurso(*request.AgregarCurso) error
	UpdateCurso(*request.UpdateCurso) error
	FindAll() ([]response.UserCourseResponse, error)
	FindById(string) ([]response.UserCourseResponse, error)
	DeleteCurso(string) error
}

type XUserCourseServiceImpl struct {
	XUserCourseRepository repository.CursoUsuarioRepository
}

func NewXUserCourseServiceImpl(xUserCourseRepository repository.CursoUsuarioRepository) XUserCourseService {
	return &XUserCourseServiceImpl{
		XUserCourseRepository: xUserCourseRepository,
	}
}

func (u *XUserCourseServiceImpl) DeleteCurso(userId string) error {
	return u.XUserCourseRepository.DeleteOne(userId)
}

func (u *XUserCourseServiceImpl) FindAll() (users []response.UserCourseResponse, err error) {
	results, err := u.XUserCourseRepository.FindAll()
	if err != nil {
		return []response.UserCourseResponse{}, err
	}

	for _, value := range results {
		user := response.UserCourseResponse{
			IdUsuario:  value.UserId,
			CourseName: value.CourseName,
			State:      value.State,
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *XUserCourseServiceImpl) FindById(userId string) ([]response.UserCourseResponse, error) {
	results, err := u.XUserCourseRepository.FindById(userId)
	if err != nil {
		return []response.UserCourseResponse{}, err
	}

	var users []response.UserCourseResponse
	for _, value := range results {
		user := response.UserCourseResponse{
			IdUsuario:  value.UserId,
			CourseName: value.CourseName,
			State:      value.State,
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *XUserCourseServiceImpl) AgregarCurso(request *request.AgregarCurso) error {
	userCourse := model.UserCourse{
		UserId:     request.UserId,
		CourseName: request.CourseName,
		State:      request.State,
	}

	_, err := u.XUserCourseRepository.InsertOne(userCourse)
	if err != nil {
		return fmt.Errorf("failed to insert user course: %w", err)
	}

	return nil
}

func (u *XUserCourseServiceImpl) UpdateCurso(request *request.UpdateCurso) error {

	userCourse := model.UserCourse{
		State: request.State,
	}

	_, err := u.XUserCourseRepository.UpdateOne(userCourse)
	if err != nil {
		return fmt.Errorf("failed to update user course: %w", err)
	}

	return nil
}
