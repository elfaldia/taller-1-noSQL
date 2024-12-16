package service

import (
	"errors"
	"fmt"
	"log"
	"time"

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
	DeleteCurso(string, string) error
	AddCourseRating(userId string, courseName string, rating int) error
	GetCourseRating(courseName string) (float64, error)
}

type XUserCourseServiceImpl struct {
	XUserCourseRepository repository.CursoUsuarioRepository
	UserService           UserService
	CursoService          CursoService
}

func NewXUserCourseServiceImpl(xUserCourseRepository repository.CursoUsuarioRepository, userService UserService, cursoService CursoService) XUserCourseService {
	return &XUserCourseServiceImpl{
		XUserCourseRepository: xUserCourseRepository,
		UserService:           userService,
		CursoService:          cursoService,
	}
}

func (u *XUserCourseServiceImpl) DeleteCurso(userId string, cursoName string) error {
	return u.XUserCourseRepository.DeleteOne(userId, cursoName)
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

	_, err := u.UserService.FindById(request.UserId)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	current_time := time.Now()
	formatedTime := current_time.Format("2006-January-02")

	userCourse := model.UserCourse{

		UserId:       request.UserId,
		CourseName:   request.CourseName,
		State:        request.State,
		StartDate:    formatedTime,
		ClasesVistas: 0,
	}

	_, err = u.XUserCourseRepository.InsertOne(userCourse)
	if err != nil {
		return fmt.Errorf("failed to insert user course: %w", err)
	}
	return nil
}

func (u *XUserCourseServiceImpl) UpdateCurso(request *request.UpdateCurso) error {

	if request.ClasesVistas < 0 {
		return fmt.Errorf("clases vistas no puede ser menor que 0")
	}

	cantidadTotal, err := u.CursoService.GetCantidadClases(request.CourseName)
	if err != nil {
		return err
	}

	log.Printf("%v", cantidadTotal)

	if cantidadTotal < request.ClasesVistas {
		return fmt.Errorf("clases vistas no puede ser mayor que el total de clases del curso")
	}

	userCourse := model.UserCourse{
		UserId:       request.UserId,
		CourseName:   request.CourseName,
		State:        request.State,
		ClasesVistas: request.ClasesVistas,
	}

	_, err = u.XUserCourseRepository.UpdateOne(userCourse)
	if err != nil {
		return fmt.Errorf("failed to update user course: %w", err)
	}

	return nil
}

func (u *XUserCourseServiceImpl) AddCourseRating(userId string, courseName string, rating int) error {

	res, err := u.UserService.FindById(userId)
	if err != nil {
		return err
	}

	if (res == response.UserResponse{}) {
		return errors.New("usuario no válido")
	}

	_, err = u.CursoService.FindById(courseName)
	if err != nil {
		return err
	}
	return u.XUserCourseRepository.AddCourseRating(userId, courseName, rating)
}

// GetCourseRating implements XUserCourseService.
func (u *XUserCourseServiceImpl) GetCourseRating(courseName string) (float64, error) {

	_, err := u.CursoService.FindById(courseName)
	if err != nil {
		return -1, err
	}
	return u.XUserCourseRepository.GetCourseRating(courseName)

}
