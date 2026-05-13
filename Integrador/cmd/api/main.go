package main

import (
	"integrador/internal/repository"
	"integrador/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Elegimos qué repositorio usar (Inyección)
	// Podrías cambiar esto por repository.NewPostgresRepo("...") más tarde
	repo := repository.NewMemoryRepo()
	serv := service.NewImageService(repo)

	// 2. Configuración de la API (Rutas)
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		// Aquí usarías CONDICIONALES para validar la petición
		err := serv.SubirImagen("1", "nueva_foto", "jpg", []string{"test"})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "Imagen guardada"})
	})

	r.Run(":8080") // Corre en el puerto 8080
}
