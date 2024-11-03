package repository

import (
	"context"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClaseRepository interface {
	FindAll() ([]model.Clase, error)
	FindAllByIdUnidad(string) ([]model.Clase, error)
	FindById(string) (model.Clase, error)
	InsertOne(clase model.Clase) (model.Clase, error)
	InsertMany(clases []model.Clase) ([]model.Clase, error)
	DeleteClase(string) error
}

type ClaseRepositoryImpl struct {
	ClaseCollection *mongo.Collection
}

func NewClaseRepositoryImpl(claseCollection *mongo.Collection) ClaseRepository {
	return &ClaseRepositoryImpl{ClaseCollection: claseCollection}
}

func (c *ClaseRepositoryImpl) FindAll() ([]model.Clase, error) {

	cursor, err := c.ClaseCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []model.Clase
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (c *ClaseRepositoryImpl) FindById(_id string) (clase model.Clase, err error) {

	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return clase, err
	}

	err = c.ClaseCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&clase)
	if err != nil {
		return clase, err
	}
	return clase, nil
}

// InsertMany implements ClaseRepository.
func (c *ClaseRepositoryImpl) InsertMany(clases []model.Clase) ([]model.Clase, error) {
	var documentos []interface{}
	var clasesInsertados []model.Clase
	for _, clase := range clases {
		documentos = append(documentos, clase)
	}

	resultado, err := c.ClaseCollection.InsertMany(context.TODO(), documentos)
	if err != nil {
		return nil, err
	}

	for i, id := range resultado.InsertedIDs {
		clases[i].Id = id.(primitive.ObjectID)
		clasesInsertados = append(clasesInsertados, clases[i])
	}
	return clasesInsertados, nil
}

// InsertOne implements ClaseRepository.
func (c *ClaseRepositoryImpl) InsertOne(clase model.Clase) (model.Clase, error) {

	insertarClase, err := c.ClaseCollection.InsertOne(context.TODO(), clase)
	if err != nil {
		return clase, err
	}
	clase.Id = insertarClase.InsertedID.(primitive.ObjectID)
	return clase, nil
}

func (c *ClaseRepositoryImpl) FindAllByIdUnidad(idUnidad string) (clases []model.Clase, err error) {

	objectUnidadId, err := primitive.ObjectIDFromHex(idUnidad)
	if err != nil {
		return clases, err
	}
	cursor, err := c.ClaseCollection.Find(context.TODO(), bson.M{"id_unidad": objectUnidadId})
	if err != nil {
		return clases, err
	}
	err = cursor.All(context.TODO(), &clases)
	if err != nil {
		return nil, err
	}
	return clases, nil
}

func (c *ClaseRepositoryImpl) DeleteClase(idClaseS string) error {

	idClase, err := primitive.ObjectIDFromHex(idClaseS)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: idClase}}

	_, err = c.ClaseCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
