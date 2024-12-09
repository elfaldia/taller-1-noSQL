package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/elfaldia/taller-noSQL/internal/model"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(userId string) (model.User, error)
	InsertOne(user model.User) (model.User, error)
	UpdateOne(user model.User) (model.User, error)
	DeleteOne(string) error
}

type UserRepositoryImpl struct {
	tableName      string
	UserCollection *dynamodb.Client
}

func NewUserRepositoryImpl(userCollection *dynamodb.Client) UserRepository {
	return &UserRepositoryImpl{
		UserCollection: userCollection,
		tableName:      "Users",
	}
}

// FindAll implements UserRepository.
func (u *UserRepositoryImpl) FindAll() (users []model.User, err error) {

	input := &dynamodb.ScanInput{
		TableName: &u.tableName,
	}

	result, err := u.UserCollection.Scan(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to scan table: %w", err)
	}

	err = attributevalue.UnmarshalListOfMaps(result.Items, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal scan result items: %w", err)
	}

	return users, nil

}

// DeleteONe implements UserRepository.
func (u *UserRepositoryImpl) DeleteOne(UserId string) error {

	user := model.User{
		UserId: UserId,
	}

	key, err := user.GetKey()
	if err != nil {
		return err
	}

	_, err = u.UserCollection.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(u.tableName),
		Key:       key,
	})
	if err != nil {
		return err
	}
	return nil
}

// FindById implements UserRepository.
func (u *UserRepositoryImpl) FindById(userId string) (user model.User, err error) {
	userI := model.User{
		UserId: userId,
	}

	key, err := userI.GetKey()

	if err != nil {
		return model.User{}, err
	}

	response, err := u.UserCollection.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(u.tableName),
	})

	if err != nil {
		return model.User{}, err
	}
	err = attributevalue.UnmarshalMap(response.Item, &user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// InsertOne implements UserRepository.
func (u *UserRepositoryImpl) InsertOne(user model.User) (model.User, error) {

	user.UserId = user.Email

	av, err := attributevalue.MarshalMap(user)
	if err != nil {
		return model.User{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(u.tableName),
	}

	_, err = u.UserCollection.PutItem(context.TODO(), input)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// UpdateOne implements UserRepository.
func (u *UserRepositoryImpl) UpdateOne(user model.User) (model.User, error) {
	update := expression.Set(expression.Name("Nombre"), expression.Value(user.Nombre)).
		Set(expression.Name("Email"), expression.Value(user.Email)).
		Set(expression.Name("Clave"), expression.Value(user.Clave))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return model.User{}, err
	}

	key, err := user.GetKey()
	if err != nil {
		return model.User{}, err
	}

	response, err := u.UserCollection.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(u.tableName),
		Key:                       key,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
		ReturnValues:              "ALL_NEW",
	})
	if err != nil {
		return model.User{}, err
	}

	var updatedUser model.User
	err = attributevalue.UnmarshalMap(response.Attributes, &updatedUser)
	if err != nil {
		return model.User{}, err
	}

	return updatedUser, nil
}
