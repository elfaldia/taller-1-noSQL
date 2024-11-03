package repository

import (
	"context"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CursoRepository interface {
	FindAll() ([]model.Curso, error)
	FindById(cursoId string) (model.Curso, error)
	InsertOne(curso model.Curso) (model.Curso, error)
	InsertMany(cursos []model.Curso) ([]model.Curso, error)
	DeleteCurso(string) error
}

type CursoRepositoryImpl struct {
	CursoCollection *mongo.Collection
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
	return results, nil
}

// FindById implements CursoRepository.
func (c *CursoRepositoryImpl) FindById(_id string) (model.Curso, error) {

	var result model.Curso

	objectID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return result, err
	}

	err = c.CursoCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// InsertMany implements CursoRepository.
func (c *CursoRepositoryImpl) InsertMany(cursos []model.Curso) ([]model.Curso, error) {
	var cursosInsertados []model.Curso
	var documentos []interface{}

	for _, curso := range cursos {
		documentos = append(documentos, curso)
	}

	resultado, err := c.CursoCollection.InsertMany(context.TODO(), documentos)
	if err != nil {
		return nil, err
	}

	for i, id := range resultado.InsertedIDs {
		cursos[i].Id = id.(primitive.ObjectID)
		cursosInsertados = append(cursosInsertados, cursos[i])
	}
	return cursosInsertados, nil
}

// InsertOne implements CursoRepository.
func (c *CursoRepositoryImpl) InsertOne(curso model.Curso) (model.Curso, error) {

	insertarCurso, err := c.CursoCollection.InsertOne(context.TODO(), curso)
	if err != nil {
		return curso, err
	}

	curso.Id = insertarCurso.InsertedID.(primitive.ObjectID)

	return curso, nil

}

// DeleteCurso implements CursoRepository.
func (c *CursoRepositoryImpl) DeleteCurso(idCursoString string) error {

	idCurso, err := primitive.ObjectIDFromHex(idCursoString)
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: idCurso}}

	_, err = c.CursoCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
