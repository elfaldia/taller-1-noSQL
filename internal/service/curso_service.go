package service

import (
	"errors"
	"log"
	"math/rand"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CursoService interface {
	CreateCurso(*request.CreateCursoRequest) (idCurso string, err error)
	FindAll() ([]response.CursoReponse, error)
	FindById(string) (response.CursoReponse, error)
	AddComentarioCurso(comentario model.ComentarioCurso) error
	GetComentariosByCursoId(cursoID string) ([]model.ComentarioCurso, error)
	DeleteCurso(string)
	GetRandomId() (primitive.ObjectID, error)
	GetCantidadClases(string) (int, error)
}

type CursoServiceImpl struct {
	CursoRepository      repository.CursoRepository
	UnidadService        UnidadService
	ClaseService         ClaseService
	ComentarioRepository repository.ComentarioRepository
	Validate             *validator.Validate
	db                   *mongo.Database
}

func NewCursoServiceImpl(
	cursoRepository repository.CursoRepository,
	comentarioRepository repository.ComentarioRepository,
	validate *validator.Validate,
	db *mongo.Database,
	unidadService UnidadService,
	claseService ClaseService,
) (service CursoService, err error) {
	if validate == nil {
		return nil, errors.New("validator no puede ser nil")
	}
	return &CursoServiceImpl{
		CursoRepository:      cursoRepository,
		ComentarioRepository: comentarioRepository,
		UnidadService:        unidadService,
		ClaseService:         claseService,
		Validate:             validate,
		db:                   db,
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

// InsertOne implements CursoService.
func (c *CursoServiceImpl) CreateCurso(req *request.CreateCursoRequest) (idCurso string, err error) {

	err = c.Validate.Struct(req)
	if err != nil {
		return "", err
	}

	curso := model.Curso{
		Nombre:           req.Nombre,
		Descripcion:      req.Descripcion,
		Valoracion:       req.Valoracion,
		ImagenMiniatura:  req.ImagenMiniatura,
		ImagenBanner:     req.ImagenBanner,
		CantidadUsuarios: req.CantidadUsuario,
	}

	curso, err = c.CursoRepository.InsertOne(curso)
	if err != nil {
		return "", err
	}

	for _, unidadReq := range req.Unidades {

		createUnidadReq := request.CrearUnidadRequest{
			Nombre:  unidadReq.NombreUnidad,
			Indice:  unidadReq.IndiceUnidad,
			IdCurso: curso.Id.Hex(),
		}

		unidad, err := c.UnidadService.CreateOne(createUnidadReq)
		if err != nil {
			c.DeleteCurso(curso.Id.Hex())
			return "", err
		}
		for _, claseReq := range unidadReq.Clases {

			createClaseReq := request.CreateClaseRequest{
				Nombre:            claseReq.NombreClase,
				Descripcion:       claseReq.Descripcion,
				Video:             claseReq.Video,
				Indice:            claseReq.IndiceClase,
				MaterialAdicional: claseReq.MaterialAdicional,
				IdUnidad:          unidad.Id.Hex(),
			}
			_, err := c.ClaseService.CreateClase(createClaseReq)
			if err != nil {
				c.DeleteCurso(curso.Id.Hex())
				return "", err
			}
		}
	}
	return curso.Id.Hex(), nil
}

func (c *CursoServiceImpl) AddComentarioCurso(comentario model.ComentarioCurso) error {
	if comentario.ComentarioID == "" {
		comentario.ComentarioID = "newComentarioID" // Puedes generar un nuevo ID aquí, como un UUID o similar
	}

	if comentario.IdCurso == "" || comentario.IdUsuario == "" {
		return errors.New("id_curso y id_usuario no pueden ser vacíos")
	}

	return c.ComentarioRepository.InsertOne(comentario)
}

func (c *CursoServiceImpl) GetComentariosByCursoId(cursoID string) ([]model.ComentarioCurso, error) {
	if _, err := primitive.ObjectIDFromHex(cursoID); err != nil {
		return nil, errors.New("cursoID inválido")
	}

	return c.ComentarioRepository.FindByCurso(cursoID)
}

func (c *CursoServiceImpl) DeleteCurso(cursoId string) {

	c.CursoRepository.DeleteCurso(cursoId)
	unidades, _ := c.UnidadService.FindByIdCurso(cursoId)
	for _, unidad := range unidades {
		clases, _ := c.ClaseService.FindAllByIdUnidad(unidad.Id.Hex())
		for _, clase := range clases {
			c.ClaseService.DeleteClase(clase.Id.Hex())
		}
		c.UnidadService.DeleteUnidad(unidad.Id.Hex())
	}

}

func (c *CursoServiceImpl) GetRandomId() (primitive.ObjectID, error) {

	cursos, err := c.FindAll()
	if err != nil {
		return primitive.NilObjectID, err
	}
	var ids []primitive.ObjectID
	for _, curso := range cursos {
		ids = append(ids, curso.Id)
	}
	randomIndex := rand.Intn(len(ids))
	return ids[randomIndex], nil
}

func (c *CursoServiceImpl) GetCantidadClases(cursoId string) (int, error) {

	unidades, err := c.UnidadService.FindByIdCurso(cursoId)
	if err != nil {
		return 0, err
	}

	cantidad_clases := 0
	for _, unidad := range unidades {

		log.Printf("%s", unidad.Id.Hex())
		cant, err := c.ClaseService.FindAllByIdUnidad(unidad.Id.Hex())
		if err != nil {
			return 0, err
		}
		cantidad_clases = cantidad_clases + len(cant)
	}

	return cantidad_clases, nil
}
