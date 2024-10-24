package repository

import (
	"context"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ComentarioClaseRepository interface {
	FindAllByIdClase(string) ([]model.ComentarioClase, error)
	FindById(string) (model.ComentarioClase, error)
	InsertOne(model.ComentarioClase) (model.ComentarioClase, error)
}

type ComentarioClaseRepositoryImpl struct {
	ComentarioClaseCollection *mongo.Collection
}

func NewComentarioClaseRepositoryImpl(comentarioClaseCollection *mongo.Collection) ComentarioClaseRepository {
	return &ComentarioClaseRepositoryImpl{ComentarioClaseCollection: comentarioClaseCollection}
}

func (c *ComentarioClaseRepositoryImpl) FindAllByIdClase(idClase string) (comentarios []model.ComentarioClase, err error) {

	objectClaseID, err := primitive.ObjectIDFromHex(idClase)
	if err != nil {
		return []model.ComentarioClase{}, err
	}
	cursos, err := c.ComentarioClaseCollection.Find(context.TODO(), bson.M{"id_clase": objectClaseID})
	if err != nil {
		return []model.ComentarioClase{}, err
	}
	err = cursos.All(context.TODO(), &comentarios)
	if err != nil {
		return []model.ComentarioClase{}, err
	}
	return comentarios, nil

}

func (c *ComentarioClaseRepositoryImpl) FindById(_id string) (comentario model.ComentarioClase, err error) {

	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return model.ComentarioClase{}, err
	}
	err = c.ComentarioClaseCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&comentario)
	if err != nil {
		return model.ComentarioClase{}, err
	}
	return comentario, nil
}

func (c *ComentarioClaseRepositoryImpl) InsertOne(comentario model.ComentarioClase) (model.ComentarioClase, error) {

	insertComentario, err := c.ComentarioClaseCollection.InsertOne(context.TODO(), comentario)
	if err != nil {
		return model.ComentarioClase{}, err
	}
	comentario.Id = insertComentario.InsertedID.(primitive.ObjectID)
	return comentario, nil
}
