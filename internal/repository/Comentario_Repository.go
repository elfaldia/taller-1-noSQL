package repository

import (
	"context"
	"fmt"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ComentarioRepository interface {
	InsertOne(comentario model.ComentarioCurso) error
	FindByCurso(cursoID string) ([]model.ComentarioCurso, error)
	DeleteOne(comentarioID string) error
}

type ComentarioRepositoryImpl struct {
	Driver neo4j.DriverWithContext
}

func NewComentarioRepositoryImpl(driver neo4j.DriverWithContext) ComentarioRepository {
	return &ComentarioRepositoryImpl{Driver: driver}
}

func (r *ComentarioRepositoryImpl) InsertOne(comentario model.ComentarioCurso) error {
	ctx := context.TODO()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx , `
		CREATE (c:Comentario {
			comentario_id: $comentario_id,
			curso_id: $id_curso,
			usuario_id: $id_usuario,
			nombre: $nombre,
			fecha: $fecha,
			titulo: $titulo,
			contenido: $detalle,
			likes: $likes,
			dislikes: $dislikes
		})`,
		map[string]interface{}{
			"comentario_id": comentario.ComentarioID,
			"curso_id":      comentario.IdCurso,
			"usuario_id":    comentario.IdUsuario,
			"nombre":        comentario.Nombre,
			"fecha":         comentario.Fecha,
			"titulo":        comentario.Titulo,
			"contenido":     comentario.Detalle,
			"likes":         comentario.Likes,
			"dislikes":      comentario.Dislikes,
		})

	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}
	return nil
}

func (r *ComentarioRepositoryImpl) FindByCurso(cursoID string) ([]model.ComentarioCurso, error) {
	ctx:= context.TODO()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	// Convierte cursoID a ObjectID
	cursoIDObj, err := primitive.ObjectIDFromHex(cursoID)
	if err != nil {
		return nil, fmt.Errorf("invalid curso_id: %w", err)
	}

	result, err := session.Run(ctx, `
		MATCH (c:Comentario {id_curso: $id_curso})
		RETURN c.comentario_id, c.usuario_id, c.nombre, c.fecha, c.titulo, c.detalle, c.likes, c.dislikes
		ORDER BY c.likes DESC`,
		map[string]interface{}{
			"id_curso": cursoID,
		})

	if err != nil {
		return nil, fmt.Errorf("failed to find comments: %w", err)
	}

	var comentarios []model.ComentarioCurso
	for result.Next(ctx) {
		record := result.Record()

		// Convierte comentario_id a primitive.ObjectID
		comentarioIDStr, ok := record.Values[0].(string)
		if !ok {
			return nil, fmt.Errorf("failed to cast comentario_id to string")
		}
		comentarioID, err := primitive.ObjectIDFromHex(comentarioIDStr)
		if err != nil {
			return nil, fmt.Errorf("invalid comentario_id: %w", err)
		}

		// Convierte usuario_id a primitive.ObjectID
		usuarioIDStr, ok := record.Values[1].(string)
		if !ok {
			return nil, fmt.Errorf("failed to cast usuario_id to string")
		}
		usuarioID, err := primitive.ObjectIDFromHex(usuarioIDStr)
		if err != nil {
			return nil, fmt.Errorf("invalid usuario_id: %w", err)
		}

		// Conversión de otros campos
		nombre, _ := record.Values[2].(string)
		fecha, _ := record.Values[3].(string)
		titulo, _ := record.Values[4].(string)
		detalle, _ := record.Values[5].(string)
		likes := int(record.Values[6].(int64))
		dislikes := int(record.Values[7].(int64))

		// Agrega el comentario a la lista
		comentarios = append(comentarios, model.ComentarioCurso{
			ComentarioID: comentarioID,
			IdCurso:      cursoIDObj,
			IdUsuario:    usuarioID,
			Nombre:       nombre,
			Fecha:        fecha,
			Titulo:       titulo,
			Detalle:      detalle,
			Likes:        likes,
			Dislikes:     dislikes,
		})
	}

	return comentarios, nil
}

func (r *ComentarioRepositoryImpl) DeleteOne(comentarioID string) error {
	ctx := context.TODO()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.Run(ctx, `
		MATCH (c:Comentario {comentario_id: $comentario_id})
		DELETE c`,
		map[string]interface{}{
			"comentario_id": comentarioID,
		})

	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}
	return nil
}
