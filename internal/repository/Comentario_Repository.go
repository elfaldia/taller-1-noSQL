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

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		query := `
			MERGE (u:User {UserId: $userId})
			MERGE (c:Course {CourseName: $courseName})
			CREATE (u)-[r:COMENTA]->(c)
			SET 
				r.nombre = $nombre,
				r.fecha = $fecha,
				r.titulo = $titulo,
				r.detalle = $detalle,
				r.likes = $likes,
				r.dislikes = $dislikes
		`

		params := map[string]interface{}{
			"userId":     comentario.IdUsuario,
			"courseName": comentario.IdCurso,
			"fecha":      comentario.Fecha,
			"nombre":     comentario.Nombre,
			"titulo":     comentario.Titulo,
			"detalle":    comentario.Detalle,
			"likes":      comentario.Likes,
			"dislikes":   comentario.Dislikes,
		}
		_, err := tx.Run(ctx, query, params)
		return nil, err
	})

	if err != nil {
		return fmt.Errorf("falied to add comments in course: %w", err)
	}

	return nil
}

func (r *ComentarioRepositoryImpl) FindByCurso(cursoID string) ([]model.ComentarioCurso, error) {
	ctx := context.TODO()
	session := r.Driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)

	result, err := session.Run(ctx, `
		MATCH (c:Course {CourseName: $id_curso})<-[r:COMENTA]-(u:User)
		RETURN 
			c.CourseName AS courseName,
			u.UserId AS userId,
			r.nombre AS nombre, 
			r.fecha AS fecha, 
			r.titulo AS titulo, 
			r.detalle AS detalle, 
			r.likes AS likes, 
			r.dislikes AS dislikes
		ORDER BY r.likes DESC`,
		map[string]interface{}{
			"id_curso": cursoID,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to find comments: %w", err)
	}

	var comentarios []model.ComentarioCurso

	// Iterar sobre los resultados
	for result.Next(ctx) {
		record := result.Record()

		// Obtener valores del registro
		courseName, _ := record.Get("courseName")
		userId, _ := record.Get("userId")
		nombre, _ := record.Get("nombre")
		fecha, _ := record.Get("fecha")
		titulo, _ := record.Get("titulo")
		detalle, _ := record.Get("detalle")
		likes, _ := record.Get("likes")
		dislikes, _ := record.Get("dislikes")

		// Crear el comentario y añadirlo a la lista
		comentarios = append(comentarios, model.ComentarioCurso{
			IdUsuario: userId.(string),
			IdCurso:   courseName.(string),
			Nombre:    nombre.(string),
			Fecha:     fecha.(string),
			Titulo:    titulo.(string),
			Detalle:   detalle.(string),
			Likes:     int(likes.(int64)),
			Dislikes:  int(dislikes.(int64)),
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
