package main

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
)

type Alumno struct {
	Id              int       `fake:"{number:1}" json:"id"`
	Nombre          string    `fake:"{firstname}" json:"nombre"`
	Apellido        string    `fake:"{lastname}" json:"apellido"`
	NroDocumento    int       `fake:"{number:10000000}" json:"nro_documento"`
	TipoDocumento   string    `fake:"{randomstring:[DNI,Pasaporte,LE]}" json:"tipo_documento"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"` //Default Go format (RFC3339)
	FechaIngreso    time.Time `json:"fecha_ingreso"`    //Default Go format (RFC3339)
	Sexo            string    `fake:"{randomstring:[M,F]}" json:"sexo"`
	NroLegajo       int       `fake:"{number:1000}" json:"nro_legajo"`
	EspecialidadId  int       `fake:"{number:1}" json:"especialidad_id"`
}

func setupRouter() *gin.Engine {
	alumnos := generateMockData()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Mock API is running",
		})
	})

	r.GET("/api/v1/alumnos", func(c *gin.Context) {

		c.JSON(http.StatusOK, alumnos)
	})

	r.GET("/api/v1/alumnos/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
			return
		}

		fmt.Println("Searching for alumno with ID:", id)
		idx := slices.IndexFunc(alumnos, func(a Alumno) bool { return a.Id == id })
		if idx != -1 {
			c.JSON(http.StatusOK, alumnos[idx])
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Alumno not found"})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8088")
}

func generateMockData() []Alumno {
	gofakeit.Seed(0)
	alumnos := []Alumno{}
	gofakeit.Slice(&alumnos)
	return alumnos
}
