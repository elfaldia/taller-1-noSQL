package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/elfaldia/taller-noSQL/internal/model"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type CursoUsuarioRepository interface {
	FindAll() ([]model.UserCourse, error)
	FindById(userId string) ([]model.UserCourse, error)
	InsertOne(userCourse model.UserCourse) (model.UserCourse, error)
	UpdateOne(userCourse model.UserCourse) (model.UserCourse, error)
	DeleteOne(string, string) error
	AddCourseRating(userId string, courseName string, rating int) error
	GetCourseRating(courseName string) (float64, error)
}

type CursoUsuarioRepositoryImpl struct {
	tableName            string
	UserCourseCollection *dynamodb.Client
	DriverNeo4j          *neo4j.DriverWithContext
}

func NewCursoUsuarioRepositoryImpl(userCourseCollection *dynamodb.Client, userCourseGraph *neo4j.DriverWithContext) CursoUsuarioRepository {
	return &CursoUsuarioRepositoryImpl{
		UserCourseCollection: userCourseCollection,
		tableName:            "CursoUsuario",
		DriverNeo4j:          userCourseGraph,
	}
}

func (u *CursoUsuarioRepositoryImpl) InsertOne(userCourse model.UserCourse) (model.UserCourse, error) {
	item, err := attributevalue.MarshalMap(userCourse)
	if err != nil {
		return model.UserCourse{}, fmt.Errorf("failed to marshal user course: %w", err)
	}

	input := &dynamodb.PutItemInput{
		TableName: &u.tableName,
		Item:      item,
	}

	_, err = u.UserCourseCollection.PutItem(context.TODO(), input)
	if err != nil {
		return model.UserCourse{}, fmt.Errorf("failed to put item in DynamoDB: %w", err)
	}

	return userCourse, nil
}

func (u *CursoUsuarioRepositoryImpl) UpdateOne(userCourse model.UserCourse) (model.UserCourse, error) {
	key, err := attributevalue.MarshalMap(map[string]string{
		"UserId":     userCourse.UserId,
		"CourseName": userCourse.CourseName,
	})
	if err != nil {
		return model.UserCourse{}, fmt.Errorf("failed to marshal key: %w", err)
	}

	update, err := attributevalue.MarshalMap(userCourse)
	if err != nil {
		return model.UserCourse{}, fmt.Errorf("failed to marshal user course: %w", err)
	}

	input := &dynamodb.UpdateItemInput{
		TableName: &u.tableName,
		Key:       key,
		AttributeUpdates: map[string]types.AttributeValueUpdate{
			"State": {
				Value:  update["State"],
				Action: types.AttributeActionPut,
			},
		},
		ReturnValues: types.ReturnValueUpdatedNew,
	}

	_, err = u.UserCourseCollection.UpdateItem(context.TODO(), input)
	if err != nil {
		return model.UserCourse{}, fmt.Errorf("failed to update item in DynamoDB: %w", err)
	}

	return userCourse, nil
}

func (u *CursoUsuarioRepositoryImpl) DeleteOne(userId string, cursoName string) error {
	userCourse := model.UserCourse{
		UserId:     userId,
		CourseName: cursoName,
	}

	key, err := userCourse.GetKey()
	if err != nil {
		return fmt.Errorf("failed to get key: %w", err)
	}

	_, err = u.UserCourseCollection.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: &u.tableName,
		Key:       key,
	})
	if err != nil {
		return fmt.Errorf("failed to delete item in DynamoDB: %w", err)
	}

	return nil
}

func (u *CursoUsuarioRepositoryImpl) FindAll() ([]model.UserCourse, error) {
	input := &dynamodb.ScanInput{
		TableName: &u.tableName,
	}

	result, err := u.UserCourseCollection.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}

	var userCourses []model.UserCourse
	err = attributevalue.UnmarshalListOfMaps(result.Items, &userCourses)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal scan result items: %w", err)
	}

	return userCourses, nil
}

func (u *CursoUsuarioRepositoryImpl) FindById(userId string) ([]model.UserCourse, error) {
	keyCondition := expression.Key("UserId").Equal(expression.Value(userId))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build expression: %w", err)
	}

	input := &dynamodb.QueryInput{
		TableName:                 &u.tableName,
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	}

	result, err := u.UserCourseCollection.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query table: %w", err)
	}

	var userCourses []model.UserCourse
	err = attributevalue.UnmarshalListOfMaps(result.Items, &userCourses)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal query result items: %w", err)
	}

	return userCourses, nil
}

func (u *CursoUsuarioRepositoryImpl) AddCourseRating(userId string, courseName string, rating int) error {
    ctx := context.TODO()

    // Inicia una sesión de escritura
    session := (*u.DriverNeo4j).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
    defer session.Close(ctx)

    _, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            MERGE (u:User {UserId: $userId})
            MERGE (c:Course {CourseName: $courseName})
            MERGE (u)-[r:RATED]->(c)
            SET r.Score = $rating
        `
        params := map[string]interface{}{
            "userId":    userId,
            "courseName": courseName,
            "rating":    rating,
        }

        _, err := tx.Run(ctx, query, params)
        return nil, err
    })

    if err != nil {
        return fmt.Errorf("failed to add course rating: %w", err)
    }
    return nil
}

func (u *CursoUsuarioRepositoryImpl) GetCourseRating(courseName string) (float64, error) {
    ctx := context.TODO()

    session := (*u.DriverNeo4j).NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
    defer session.Close(ctx)

    result, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
        query := `
            MATCH (:Course {CourseName: $courseName})<-[r:RATED]-(:User)
            RETURN avg(r.Score) as averageRating
        `
        params := map[string]interface{}{
            "courseName": courseName,
        }

        records, err := tx.Run(ctx, query, params)
        if err != nil {
            return nil, err
        }

        if records.Next(ctx) {
            value, ok := records.Record().Get("averageRating")
            if !ok {
                return 0, fmt.Errorf("averageRating not found in the result")
            }

            if averageRating, isFloat := value.(float64); isFloat {
                return averageRating, nil
            }
            return 0, fmt.Errorf("averageRating is not a float64")
        }

        return 0, nil // Sin resultados
    })

    if err != nil {
        return 0, fmt.Errorf("failed to get course rating: %w", err)
    }

    if avgRating, ok := result.(float64); ok {
        return avgRating, nil
    }

    return 0, fmt.Errorf("unexpected result type")
}

