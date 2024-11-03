package service

import (
	"errors"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ComentarioClaseService interface {
	FindAllByIdClase(string) ([]response.ComentarioClaseResponse, error)
	FindById(string) (response.ComentarioClaseResponse, error)
	CreateComentarioClase(request.CreateComentarioClase) (response.ComentarioClaseResponse, error)
}

type ComentarioClaseServiceImpl struct {
	ComentarioClaseRepository repository.ComentarioClaseRepository
	ClaseService              ClaseService
	Validate                  *validator.Validate
}

func NewComentarioClaseServiceImpl(comentarioClaseRep repository.ComentarioClaseRepository, claseService ClaseService, validate *validator.Validate) (service ComentarioClaseService, err error) {
	if validate == nil {
		return nil, errors.New("validator no puede ser nil")
	}
	return &ComentarioClaseServiceImpl{
		ComentarioClaseRepository: comentarioClaseRep,
		ClaseService:              claseService,
		Validate:                  validate,
	}, nil
}

func (c *ComentarioClaseServiceImpl) FindAllByIdClase(idClase string) (comentarios []response.ComentarioClaseResponse, err error) {

	data, err := c.ComentarioClaseRepository.FindAllByIdClase(idClase)
	if err != nil {
		return []response.ComentarioClaseResponse{}, err
	}
	for _, value := range data {
		comentario := response.ComentarioClaseResponse{
			Id:       value.Id,
			Nombre:   value.Nombre,
			Fecha:    value.Fecha,
			Titulo:   value.Titulo,
			Detalle:  value.Detalle,
			Likes:    value.Likes,
			Dislikes: value.Dislikes,
			IdClase:  value.IdClase,
		}
		comentarios = append(comentarios, comentario)
	}
	return comentarios, nil
}

func (c *ComentarioClaseServiceImpl) FindById(_id string) (response.ComentarioClaseResponse, error) {
	data, err := c.ComentarioClaseRepository.FindById(_id)
	if err != nil {
		return response.ComentarioClaseResponse{}, err
	}
	res := response.ComentarioClaseResponse{
		Id:       data.Id,
		Nombre:   data.Nombre,
		Fecha:    data.Fecha,
		Titulo:   data.Titulo,
		Detalle:  data.Detalle,
		Likes:    data.Likes,
		Dislikes: data.Dislikes,
		IdClase:  data.IdClase,
	}
	return res, nil
}

func (c *ComentarioClaseServiceImpl) CreateComentarioClase(req request.CreateComentarioClase) (response.ComentarioClaseResponse, error) {
	err := c.Validate.Struct(req)
	if err != nil {
		return response.ComentarioClaseResponse{}, err
	}

	_, err = c.ClaseService.FindById(req.IdClase)
	if err != nil {
		return response.ComentarioClaseResponse{}, err
	}

	idClase, err := primitive.ObjectIDFromHex(req.IdClase)
	if err != nil {
		return response.ComentarioClaseResponse{}, err
	}

	comentario := model.ComentarioClase{
		Nombre:   req.Nombre,
		Fecha:    req.Fecha,
		Titulo:   req.Titulo,
		Detalle:  req.Detalle,
		Likes:    req.Likes,
		Dislikes: req.Dislikes,
		IdClase:  idClase,
	}
	data, err := c.ComentarioClaseRepository.InsertOne(comentario)
	if err != nil {
		return response.ComentarioClaseResponse{}, err
	}

	res := response.ComentarioClaseResponse{
		Id:       data.Id,
		Nombre:   data.Nombre,
		Fecha:    data.Fecha,
		Titulo:   data.Titulo,
		Detalle:  data.Detalle,
		Likes:    data.Likes,
		Dislikes: data.Dislikes,
		IdClase:  data.IdClase,
	}
	return res, nil
}
