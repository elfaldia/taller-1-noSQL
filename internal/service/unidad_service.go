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

type UnidadService interface {
	FindAll() ([]response.ObtenerUnidadResponde, error)
	FindByIdCurso(string) ([]response.ObtenerUnidadResponde, error)
	CreateOne(request.CrearUnidadRequest) (model.Unidad, error)
	//CreateMany(request.CrearUnidadesRequest) ([]model.Unidad, error)
}

type UnidadServiceImpl struct {
	UnidadRepository repository.UnidadRepository
	Validate         *validator.Validate
}

func NewUnidadServiceImpl(unidadRepository repository.UnidadRepository, validate *validator.Validate) (service UnidadService, err error) {
	if validate == nil {
		return nil, errors.New("validator no puede ser nil")
	}
	return &UnidadServiceImpl{
		UnidadRepository: unidadRepository,
		Validate:         validate,
	}, nil
}

func (u *UnidadServiceImpl) FindAll() (unidades []response.ObtenerUnidadResponde, err error) {
	resultados, err := u.UnidadRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, valor := range resultados {
		unidad := response.ObtenerUnidadResponde{
			Id:      valor.Id,
			Nombre:  valor.Nombre,
			Indice:  valor.Indice,
			IdCurso: valor.IdCurso,
		}
		unidades = append(unidades, unidad)
	}
	return unidades, nil
}

func (u *UnidadServiceImpl) FindByIdCurso(id string) (unidad []response.ObtenerUnidadResponde, err error) {
	resultado, err := u.UnidadRepository.FindByIdCurso(id)
	if err != nil {
		return unidad, err
	}

	var unidadesRes []response.ObtenerUnidadResponde

	for _, unidad := range resultado {
		res := response.ObtenerUnidadResponde{
			Id:      unidad.Id,
			Nombre:  unidad.Nombre,
			Indice:  unidad.Indice,
			IdCurso: unidad.IdCurso,
		}
		unidadesRes = append(unidadesRes, res)
	}

	return unidadesRes, nil
}

func (u *UnidadServiceImpl) CreateOne(req request.CrearUnidadRequest) (model.Unidad, error) {
	err := u.Validate.Struct(req)
	if err != nil {
		return model.Unidad{}, err
	}

	idCurso, err := primitive.ObjectIDFromHex(req.IdCurso)
	if err != nil {
		return model.Unidad{}, err
	}

	unidad := model.Unidad{
		Nombre:  req.Nombre,
		Indice:  req.Indice,
		IdCurso: idCurso,
	}

	_, err = u.UnidadRepository.InsertOne(unidad)
	if err != nil {
		return unidad, err
	}
	return unidad, nil
}

/*
func (u *UnidadServiceImpl) CreateMany(req request.CrearUnidadRequest) ([]model.Unidad, error) {
	return nil, errors.New("not implemented yet")
}*/
