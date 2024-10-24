package service

import (
	"context"
	"errors"

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
	CreateCurso(request.CreateCursoRequest) error
	CreateManyCursos(request.CreateManyCursoRequest) error
	FindAll() ([]response.CursoReponse, error)
	FindById(string) (response.CursoReponse, error)
	AddComentarioCurso(comentario model.ComentarioCurso) error
	GetComentariosByCursoId(cursoID primitive.ObjectID) ([]model.ComentarioCurso, error) // Agregar este método
}

type CursoServiceImpl struct {
	CursoRepository repository.CursoRepository
	Validate        *validator.Validate
	db              *mongo.Database
}

func NewCursoServiceImpl(cursoRepository repository.CursoRepository, validate *validator.Validate, db *mongo.Database) (service CursoService, err error) {
	if validate == nil {
		return nil, errors.New("validator no puede ser nil")
	}
	return &CursoServiceImpl{
		CursoRepository: cursoRepository,
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

// InsertMany implements CursoService.
func (c *CursoServiceImpl) CreateManyCursos(req request.CreateManyCursoRequest) error {
	panic("unimplemented")
}

// InsertOne implements CursoService.
// InsertOne implements CursoService.
func (c *CursoServiceImpl) CreateCurso(req request.CreateCursoRequest) error {
	// Validar el cuerpo de la solicitud
	err := c.Validate.Struct(req)
	if err != nil {
		return err
	}

	// Crear el objeto curso a partir del request
	curso := model.Curso{
		Nombre:           req.Nombre,
		Descripcion:      req.Descripcion,
		Valoracion:       req.Valoracion, // Asegúrate de incluir la valoración en la estructura
		ImagenMiniatura:  req.ImagenMiniatura,
		ImagenBanner:     req.ImagenBanner,
		CantidadUsuarios: req.CantidadUsuarios,
	}

	// Insertar el curso en la base de datos
	_, err = c.CursoRepository.InsertOne(curso)
	if err != nil {
		return err
	}

	return nil
}

func (s *CursoServiceImpl) AddComentarioCurso(comentario model.ComentarioCurso) error {
	// Acceder a la colección de comentarios
	collection := s.db.Collection("comentarios_curso")

	// Insertar el comentario en la colección
	_, err := collection.InsertOne(context.TODO(), comentario)
	if err != nil {
		return err
	}

	return nil
}

func (s *CursoServiceImpl) GetComentariosByCursoId(cursoID primitive.ObjectID) ([]model.ComentarioCurso, error) {
	var comentarios []model.ComentarioCurso
	collection := s.db.Collection("comentarios_curso")

	// Buscar los comentarios por el ID del curso
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
