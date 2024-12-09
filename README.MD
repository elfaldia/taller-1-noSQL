# Guía para Iniciar el Proyecto

## Requisitos Previos

### 1. Instalar Go

[https://go.dev/doc/install](https://go.dev/doc/install)


## Se Utilizo MONGOATLAS PARA LA BASE DE datos:

Se necesita enviar un email al siguiente correo para poder acceder al cluster de atlas (permiso de ip):

## correo: cristian.nettle@alumnos.ucn.cl

Con esta solicitud se le enviara una solicitud al correo en el que envio la solicitud y tendra que acceder a mongoatlas y permitir el ip, para que asi tenga acceso al cluster de mongo atlas

(ya se le envio un correo al profesor Bastian Ruiz, pero no sabemos quien evaluara)

### 2. Instalar Dependencias

Para documentar el proyecto, instala **swag** utilizando el siguiente comando:

### `go install github.com/swaggo/swag/cmd/swag@latest`

## Configuración del Proyecto

### 1. Generar Documentación

En la raíz principal del proyecto, ejecutar el siguiente comando para recopilar los comentarios y generar la documentación:

## `swag init -g ./cmd/api/main.go -o docs`

### 2. Iniciar el Proyecto

Para iniciar el proyecto, ejecuta el siguiente comando dentro del directorio del proyecto:

### `go run cmd/api/main.go cmd/api/api.go`

## Acceder a la Documentación

Una vez iniciado el proyecto, puedes acceder a la documentación a través del siguiente enlace:

[http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)


# Implementación Prueba 

Se utilizo dynamoDB utilizando los servicios de AWS, las claves de para acceder a dynamodb se utilizo 
direnv.

## Implementar direnv en tu PC
El siguiente enlace muestra los pasos para tener direnv en windows
```
https://gist.github.com/rmtuckerphx/4ace28c1605300462340ffa7b7001c6d
```
luego de seguir los pasos en el inicio del proyecto crear el archivo
```
.envrc
```
En ese archivo .envrc agregar lo siguiente:

export aws_access_key_id="<access_key_id>"
export aws_secret_access_key="<aws_secret_access_key>"
export region="sa-east-1"

Las claves de acceso se encuentran en el informe que acompaña este trabajo


Las rutas nuevas implementadas para los requerimientos de la prueba:
```
Para registro de usuario
-------------
	router := service.Group("/user")
	router.GET("", userController.FindAll)
	router.GET("/:user_id", userController.FindById)
	router.POST("", userController.CreateUser)
	router.POST("/login", userController.Login)
	router.DELETE("/:user_id", userController.DeleteUser)
------------------------------------------------------
Para cursos-usuerio
-----------------------------------
	router := service.Group("/user_curso")
	router.GET("", userCursoController.FindAll)
	router.GET("/:user_id", userCursoController.FindByIdUser)
	router.POST("", userCursoController.CreateOne)
	router.DELETE("/:user_id/:curso_name", userCursoController.DeleteOne)
	router.PATCH("", userCursoController.UpdateOne)
```








