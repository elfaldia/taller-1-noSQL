package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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

	log.Printf("Parameters: %v", map[string]interface{}{
		"comentario_id": comentario.ComentarioID,
		"id_curso":      comentario.IdCurso,
		"id_usuario":    comentario.IdUsuario,
		"nombre":        comentario.Nombre,
		"fecha":         comentario.Fecha,
		"titulo":        comentario.Titulo,
		"detalle":       comentario.Detalle,
		"likes":         comentario.Likes,
		"dislikes":      comentario.Dislikes,
	})

	_, err := session.Run(ctx, `
		CREATE (c:Comentario {
			comentario_id: $comentario_id,
			curso_id: $id_curso,
			usuario_id: $id_usuario,
			nombre: $nombre,
			fecha: $fecha,
			titulo: $titulo,
			detalle: $detalle,
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
			"detalle":       comentario.Detalle,
			"likes":         comentario.Likes,
			"dislikes":      comentario.Dislikes,
		})

	if err != nil {
		return fmt.Errorf("failed to insert comment: %w", err)
	}
	return nil
}

func (r *ComentarioRepositoryImpl) FindByCurso(cursoID string) ([]model.ComentarioCurso, error) {
	ctx := context.TODO()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
		MATCH (c:Comentario {curso_id: $id_curso})
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

		// Recuperar y convertir valores
		comentarioID, _ := record.Get("c.comentario_id")
		usuarioID, _ := record.Get("c.usuario_id")
		nombre, _ := record.Get("c.nombre")
		fecha, _ := record.Get("c.fecha")
		titulo, _ := record.Get("c.titulo")
		detalle, _ := record.Get("c.detalle")
		likes, _ := record.Get("c.likes")
		dislikes, _ := record.Get("c.dislikes")

		// Crear el comentario y añadirlo a la lista
		comentarios = append(comentarios, model.ComentarioCurso{
			ComentarioID: comentarioID.(string),
			IdCurso:      cursoID,
			IdUsuario:    usuarioID.(string),
			Nombre:       nombre.(string),
			Fecha:        fecha.(string),
			Titulo:       titulo.(string),
			Detalle:      detalle.(string),
			Likes:        int(likes.(int64)),
			Dislikes:     int(dislikes.(int64)),
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
