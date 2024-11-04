## Pasos para iniciar el proyecto
### Installar Go:
#### `https://go.dev/doc/install`

## Necesitas instalar estas dependencias:
### Para documentar el proyecto:
#### `go install github.com/swaggo/swag/cmd/swag@latest`

## comando para recopilar los comentarios:
#### `swag init -g ./cmd/api/main.go -o docs`


## Para iniciar el proyecto utilizar en el directorio del proyecto este comando:
#### `go run cmd/api/main.go cmd/api/api.go`

## Endpoint para ver la documentacion:
#### `http://localhost:8080/swagger/index.html`