package main

import (
	"fmt"
	"microservicio/capaAccesoADatos/repository"
	controller "microservicio/capaControladores"
	service "microservicio/capaFachadaServices"

	"github.com/gin-gonic/gin"
)

func main() {
	tipoRepository := repository.NewTipoAudioRepository()
	tipoService := service.NewMetadataAudioService(tipoRepository)
	tipoController := controller.NewMetadataAudioController(tipoService)

	audioRepository := repository.NewAudioRepository()
	audioServiceObj := service.NewAudioService(audioRepository)
	audioController := controller.NewAudioController(audioServiceObj)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		fmt.Printf("[REST] %s %s recibido desde %s\n", c.Request.Method, c.Request.URL.Path, c.ClientIP())
		c.Next()
	})

	router.POST("/tipos", tipoController.RegistrarTipo)
	router.GET("/tipos", tipoController.ListarTipos)
	router.GET("/tipos/:idTipo", tipoController.ConsultarTipo)

	router.GET("/tipos/:idTipo/audios", audioController.ListarPorTipo)
	router.GET("/tipos/:idTipo/audios/:id", audioController.ConsultarDetalle)

	router.Run(":8080")
}
