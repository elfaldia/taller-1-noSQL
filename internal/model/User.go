package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	UserId string `json:"user_id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
	Clave  string `json:"clave"`
}

func (u User) GetKey() (map[string]types.AttributeValue, error) {
	
	UserId, err := attributevalue.Marshal(u.UserId)
	if err != nil {
		return nil, err 
	}

	key := map[string]types.AttributeValue {
		"UserId" : UserId,
	}
	return key, nil
}