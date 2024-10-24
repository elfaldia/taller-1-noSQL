// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/curso": {
            "get": {
                "description": "get cursos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "curso"
                ],
                "summary": "Devuelve todos los carritos de la base de datos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateComentarioRequest": {
            "type": "object",
            "required": [
                "detalle",
                "dislikes",
                "fecha",
                "id_curso",
                "likes",
                "nombre",
                "titulo"
            ],
            "properties": {
                "detalle": {
                    "type": "string"
                },
                "dislikes": {
                    "type": "integer"
                },
                "fecha": {
                    "type": "string"
                },
                "id_curso": {
                    "type": "string"
                },
                "likes": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "titulo": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "response.ResponseUnidad": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
