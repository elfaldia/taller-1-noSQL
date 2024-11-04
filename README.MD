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