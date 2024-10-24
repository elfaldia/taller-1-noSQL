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

type ClaseService interface {
	FindAllByIdUnidad(string) ([]response.ClaseReponse, error)
	FindById(string) (response.ClaseReponse, error)
	CreateClase(request.CreateClaseRequest) (response.ClaseReponse, error)
	CreateManyClase([]request.CreateClaseRequest) error
}

type ClaseServiceImpl struct {
	ClaseRepository repository.ClaseRepository
	Validate        *validator.Validate
}

func NewClaseServiceImpl(claseRepository repository.ClaseRepository, validate *validator.Validate) (service ClaseService, err error) {
	if validate == nil {
		return nil, errors.New("validator no puede ser nil")
	}
	return &ClaseServiceImpl{
		ClaseRepository: claseRepository,
		Validate:        validate,
	}, nil
}

// FindAllByIdUnidad implements ClaseService.
func (c *ClaseServiceImpl) FindAllByIdUnidad(idUnidad string) (clases []response.ClaseReponse, err error) {
	data, err := c.ClaseRepository.FindAllByIdUnidad(idUnidad)
	if err != nil {
		return clases, err
	}
	for _, value := range data {
		clase := response.ClaseReponse{
			Id:                value.Id,
			Nombre:            value.Nombre,
			Descripcion:       value.Descripcion,
			Video:             value.Video,
			Indice:            value.Indice,
			MaterialAdicional: value.MaterialAdicional,
			IdUnidad:          value.IdUnidad,
		}
		clases = append(clases, clase)
	}
	return clases, nil
}

// FindById implements ClaseService.
func (c *ClaseServiceImpl) FindById(_id string) (response.ClaseReponse, error) {
	data, err := c.ClaseRepository.FindById(_id)
	if err != nil {
		return response.ClaseReponse{}, err
	}

	res := response.ClaseReponse{
		Id:                data.Id,
		Nombre:            data.Nombre,
		Descripcion:       data.Descripcion,
		Video:             data.Video,
		Indice:            data.Indice,
		MaterialAdicional: data.MaterialAdicional,
		IdUnidad:          data.IdUnidad,
	}
	return res, nil
}

// InsertMany implements ClaseService.
func (c *ClaseServiceImpl) CreateManyClase(requests []request.CreateClaseRequest) error {

	for _, req := range requests {
		err := c.Validate.Struct(req)
		if err != nil {
			return err
		}
	}
	// logica para validar que la unidad existe

	var clases []model.Clase

	for _, req := range requests {
		esUnico := c.isIndiceUniqueByUnidad(req.Indice, req.IdUnidad)
		if !esUnico {
			return errors.New("indice de clase no es único")
		}
		idUnidad, err := primitive.ObjectIDFromHex(req.IdUnidad)
		if err != nil {
			return err
		}
		clase := model.Clase{
			Nombre:            req.Nombre,
			Indice:            req.Indice,
			Video:             req.Video,
			Descripcion:       req.Descripcion,
			MaterialAdicional: req.MaterialAdicional,
			IdUnidad:          idUnidad,
		}
		clases = append(clases, clase)
	}

	_, err := c.ClaseRepository.InsertMany(clases)
	if err != nil {
		return err
	}
	return nil

}

// InsertOne implements ClaseService.
func (c *ClaseServiceImpl) CreateClase(req request.CreateClaseRequest) (response.ClaseReponse, error) {
	err := c.Validate.Struct(req)
	if err != nil {
		return response.ClaseReponse{}, err
	}

	// validar que idUnidad existe

	esUnico := c.isIndiceUniqueByUnidad(req.Indice, req.IdUnidad)
	if !esUnico {
		return response.ClaseReponse{}, errors.New("indice de clase no es único")
	}

	idUnidad, err := primitive.ObjectIDFromHex(req.IdUnidad)
	if err != nil {
		return response.ClaseReponse{}, err
	}

	clase := model.Clase{
		Nombre:            req.Nombre,
		Indice:            req.Indice,
		Video:             req.Video,
		Descripcion:       req.Descripcion,
		MaterialAdicional: req.MaterialAdicional,
		IdUnidad:          idUnidad,
	}

	data, err := c.ClaseRepository.InsertOne(clase)
	if err != nil {
		return response.ClaseReponse{}, err
	}
	res := response.ClaseReponse{
		Id:                data.Id,
		Nombre:            data.Nombre,
		Descripcion:       data.Descripcion,
		Video:             data.Video,
		Indice:            data.Indice,
		MaterialAdicional: data.MaterialAdicional,
		IdUnidad:          data.IdUnidad,
	}
	return res, nil

}

func (c *ClaseServiceImpl) isIndiceUniqueByUnidad(indice int, idUnidad string) bool {
	clases, err := c.FindAllByIdUnidad(idUnidad)
	if err != nil {
		return false
	}

	for _, value := range clases {
		if value.Indice == indice {
			return false
		}
	}
	return true
}
