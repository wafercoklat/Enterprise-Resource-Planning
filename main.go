package main

import (
	datasource "REVAMPS-CMI-APPS/internal/datasources/mongoDB"
	"REVAMPS-CMI-APPS/internal/entity/service"
	"REVAMPS-CMI-APPS/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// DataSource
	getDataSource := datasource.NewData()

	// Service
	doServices := service.New(getDataSource)

	// Handler
	doHandler := handlers.NewHTTP(doServices)

	r := gin.New()
	r.GET("/account/:id/", doHandler.Retreive)

	r.Run(":8008")
}
