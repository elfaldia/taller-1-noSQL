package repository

import (
	"context"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UnidadRepository interface {
	FindAll() ([]model.Unidad, error)
	FindByIdCurso(unidadId string) ([]model.Unidad, error)
	InsertOne(unidad model.Unidad) (model.Unidad, error)
	InsertMany(unidades []model.Unidad) ([]model.Unidad, error)
	DeleteUnidad(string) error
}

type UnidadRepositoryImpl struct {
	Collection *mongo.Collection
	Ctx        *context.Context
}

func NewUnidadRepositoryImpl(unidadCollection *mongo.Collection) UnidadRepository {
	return &UnidadRepositoryImpl{Collection: unidadCollection}
}

// Obtener todos las unidades
func (u *UnidadRepositoryImpl) FindAll() ([]model.Unidad, error) {
	unidades, err := u.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var resultado []model.Unidad
	err = unidades.All(context.TODO(), &resultado)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

// Obtener una unidad por un ID
func (u *UnidadRepositoryImpl) FindByIdCurso(CursoId string) ([]model.Unidad, error) {

	var resultado []model.Unidad

	objectID, err := primitive.ObjectIDFromHex(CursoId)
	if err != nil {
		return nil, err
	}

	unidades, err := u.Collection.Find(context.TODO(), bson.M{"id_curso": objectID})
	if err != nil {
		return nil, err
	}

	defer unidades.Close(context.TODO())

	err = unidades.All(context.TODO(), &resultado)
	if err != nil {
		return nil, err
	}

	return resultado, nil
}

// Insertar una unidad
func (u *UnidadRepositoryImpl) InsertOne(unidad model.Unidad) (model.Unidad, error) {

	insertarUnidad, err := u.Collection.InsertOne(context.TODO(), unidad)
	if err != nil {
		return unidad, err
	}

	unidad.Id = insertarUnidad.InsertedID.(primitive.ObjectID)

	return unidad, nil
}

// Ingresar muchas unidades en foramto de arrays de json
func (u *UnidadRepositoryImpl) InsertMany(unidades []model.Unidad) ([]model.Unidad, error) {
	var unidadesInsertados []model.Unidad
	var documentos []interface{}

	for _, unidad := range unidades {
		documentos = append(documentos, unidad)
	}

	resultado, err := u.Collection.InsertMany(context.TODO(), documentos)
	if err != nil {
		return nil, err
	}

	for i, id := range resultado.InsertedIDs {
		unidades[i].Id = id.(primitive.ObjectID)
		unidades[i].IdCurso = id.(primitive.ObjectID)
		unidadesInsertados = append(unidadesInsertados, unidades[i])
	}
	return unidadesInsertados, nil
}

// DeleteCurso implements CursoRepository.
func (c *UnidadRepositoryImpl) DeleteUnidad(idUnidadS string) error {

	idUnidad, err := primitive.ObjectIDFromHex(idUnidadS)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: idUnidad}}

	_, err = c.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
