package repository

import (
	"context"
	"log"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CursoRepository interface {
	FindAll() ([]model.Curso, error)
	FindById(cursoId string) (model.Curso, error)
	InsertOne(model.Curso) (model.Curso, error)
	InsertMany([]model.Curso) ([]model.Curso, error)

	// seguir ...
}

type CursoRepositoryImpl struct {
	CursoCollection *mongo.Collection
	Ctx             *context.Context
}

func NewCursoRepositoryImpl(cursoCollection *mongo.Collection) CursoRepository {
	return &CursoRepositoryImpl{CursoCollection: cursoCollection}
}

// FindAll implements CursoRepository.
func (c *CursoRepositoryImpl) FindAll() ([]model.Curso, error) {
	cursor, err := c.CursoCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	var results []model.Curso
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return nil, err
	}
	log.Printf("imagen: %s", results[0].ImagenBanner)
	return results, nil
}

// FindById implements CursoRepository.
func (c *CursoRepositoryImpl) FindById(_id string) (model.Curso, error) {
	panic("unimplemented")
}

// InsertMany implements CursoRepository.
func (c *CursoRepositoryImpl) InsertMany([]model.Curso) ([]model.Curso, error) {
	panic("unimplemented")
}

// InsertOne implements CursoRepository.
func (c *CursoRepositoryImpl) InsertOne(model.Curso) (model.Curso, error) {
	panic("unimplemented")
}
