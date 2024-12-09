package model

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserCourse struct {
	UserId       string `json:"user_id"`
	CourseName   string `json:"course_name,omitempty"` // omitempty hace que el campo sea opcional en el JSON
	State        string `json:"state"`
	ClasesVistas int    `json:"clases_vistas"`
	StartDate    string `json:"start_date"`
}

func (uc UserCourse) GetKey() (map[string]types.AttributeValue, error) {
	UserId, err := attributevalue.Marshal(uc.UserId)
	if err != nil {
		return nil, err
	}

	key := map[string]types.AttributeValue{
		"UserId": UserId,
	}

	if uc.CourseName != "" {
		CourseId, err := attributevalue.Marshal(uc.CourseName)
		if err != nil {
			return nil, err
		}
		key["CourseName"] = CourseId
	}

	return key, nil
}
