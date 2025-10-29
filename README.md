# Mock Server Alumno
API REST mock server desarrollado con Go y el framework Gin que proporciona datos de alumnos.

## Prerrequisitos
- Go 1.25 o superior
- Docker

## Instalación
Clonar el repositorio:
```
git clone <url-repositorio>cd mock-server-alumno
```
### Instalar dependencias:
```
go mod download
```

## Desarrollo Local
Ejecutar el servidor localMock Server Alumno
API REST mock server desarrollado con Go y el framework Gin que proporciona datos de alumnos.
```
go run main.go
```

El servidor se iniciará en http://localhost:8088

Endpoints de la API
```
GET / - Endpoint de verificación de salud
GET /api/v1/alumnos - Obtener todos los alumnos
GET /api/v1/alumnos/:id - Obtener alumno por ID
```

### Docker Build y Ejecución
Construir la imagen Docker:
```
docker build -t mock-server-alumno .
```
Ejecutar el contenedor:

```
docker run -p 8088:8088 mock-server-alumno
```
El servidor estará accesible en http://localhost:8088
