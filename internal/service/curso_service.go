package service

import (
	"context"
	"errors"
	"math/rand"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/elfaldia/taller-noSQL/internal/repository"
	"github.com/elfaldia/taller-noSQL/internal/request"
	"github.com/elfaldia/taller-noSQL/internal/response"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CursoService interface {
	CreateCurso(*request.CreateCursoRequest) (idCurso string, err error)
	FindAll() ([]response.CursoReponse, error)
	FindById(string) (response.CursoReponse, error)
	AddComentarioCurso(comentario model.ComentarioCurso) error
	GetComentariosByCursoId(cursoID primitive.ObjectID) ([]model.ComentarioCurso, error)
	DeleteCurso(string)
	GetRandomId() (primitive.ObjectID, error)
	GetCantidadClases(string) (int, error)
}

type CursoServiceImpl struct {
	CursoRepository repository.CursoRepository
	UnidadService   UnidadService
	ClaseService    ClaseService
	Validate        *validator.Validate
	db              *mongo.Database
}

func NewCursoServiceImpl(
	cursoRepository repository.CursoRepository,
	validate *validator.Validate,
	db *mongo.Database,
	unidadService UnidadService,
	claseService ClaseService,
) (service CursoService, err error) {
	if validate == nil {
		return nil, errors.New("validator no puede ser nil")
	}
	return &CursoServiceImpl{
		CursoRepository: cursoRepository,
		UnidadService:   unidadService,
		ClaseService:    claseService,
		Validate:        validate,
		db:              db, // Asegúrate de pasar la base de datos aquí
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

func (s *CursoServiceImpl) AddComentarioCurso(comentario model.ComentarioCurso) error {

	collection := s.db.Collection("comentarios_curso")

	_, err := collection.InsertOne(context.TODO(), comentario)
	if err != nil {
		return err
	}

	return nil
}

func (s *CursoServiceImpl) GetComentariosByCursoId(cursoID primitive.ObjectID) ([]model.ComentarioCurso, error) {
	var comentarios []model.ComentarioCurso
	collection := s.db.Collection("comentarios_curso")

	filter := bson.M{"id_curso": cursoID}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &comentarios); err != nil {
		return nil, err
	}

	return comentarios, nil
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
		cant, err := c.ClaseService.FindAllByIdUnidad(unidad.Id.String())
		if err != nil {
			return 0, err
		}
		cantidad_clases = cantidad_clases + len(cant)
	}

	return cantidad_clases, nil

}
