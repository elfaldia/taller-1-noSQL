package service

import (
	"errors"

	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/go-playground/validator"
)

type CursoService interface {
	CreateCurso(request.CreateCursoRequest) error
	CreateManyCursos(request.CreateManyCursoRequest) error
	FindAll() ([]response.CursoReponse, error)
	FindById(string) (response.CursoReponse, error)
}

type CursoServiceImpl struct {
	CursoRepository repository.CursoRepository
	Validate        *validator.Validate
}

func NewCursoServiceImpl(cursoRepository repository.CursoRepository, validate *validator.Validate) (service CursoService, err error) {
	if validate == nil {
		return nil, errors.New("validator no puede ser nil")
	}
	return &CursoServiceImpl{
		CursoRepository: cursoRepository,
		Validate:        validate,
	}, nil
}

// FindAll implements CursoService.
func (c *CursoServiceImpl) FindAll() (cursos []response.CursoReponse, err error) {
	results, err := c.CursoRepository.FindAll()
	if err != nil {
		return nil, err
	}

	for _, value := range results {
		curso := response.CursoReponse{
			Id:               value.Id,
			Nombre:           value.Nombre,
			Descripcion:      value.Descripcion,
			ImagenMiniatura:  value.ImagenMiniatura,
			ImagenBanner:     value.ImagenBanner,
			Valoracion:       value.Valoracion,
			CantidadUsuarios: value.CantidadUsuarios,
		}
		cursos = append(cursos, curso)
	}
	return cursos, nil
}

// FindById implements CursoService.
func (c *CursoServiceImpl) FindById(_id string) (curso response.CursoReponse, err error) {
	data, err := c.CursoRepository.FindById(_id)
	if err != nil {
		return response.CursoReponse{}, err
	}
	res := response.CursoReponse{
		Id:               data.Id,
		Nombre:           data.Nombre,
		Descripcion:      data.Descripcion,
		ImagenMiniatura:  data.ImagenMiniatura,
		ImagenBanner:     data.ImagenBanner,
		Valoracion:       data.Valoracion,
		CantidadUsuarios: data.CantidadUsuarios,
	}
	return res, nil
}

// InsertMany implements CursoService.
func (c *CursoServiceImpl) CreateManyCursos(request.CreateManyCursoRequest) error {
	panic("unimplemented")
}

// InsertOne implements CursoService.
func (c *CursoServiceImpl) CreateCurso(request.CreateCursoRequest) error {
	panic("unimplemented")
}
