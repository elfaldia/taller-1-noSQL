package repository

import (
	"github.com/elfaldia/taller-noSQL/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type CursoRepository interface {
	FindAll() ([]model.Curso, error)
	FindById(cursoId int) (model.Curso, error)
	// seguir ...
}

type CursoRepositoryImpl struct {
	CursoCollection *mongo.Collection
}


func NewCursoRepositoryImpl(cursoCollection *mongo.Collection) CursoRepository {
	return &CursoRepositoryImpl{CursoCollection: cursoCollection}
}


// FindAll implements CursoRepository.
func (c *CursoRepositoryImpl) FindAll() ([]model.Curso, error) {
	panic("unimplemented")
}

// FindById implements CursoRepository.
func (c *CursoRepositoryImpl) FindById(cursoId int) (model.Curso, error) {
	panic("unimplemented")
}

