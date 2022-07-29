package main

import (
	datasource "REVAMPS-CMI-APPS/internal/datasources/sqlDB"
	"REVAMPS-CMI-APPS/internal/domain/service"
	"REVAMPS-CMI-APPS/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init
	dataSource := datasource.NewData()    // Data Source
	services := service.New(dataSource)   // Start Service
	handler := handlers.NewHTTP(services) // Start handler

	// Create Routes
	r := gin.New()
	r.GET("/account/:id/", handler.ViewDataById)

	r.Run(":8008")
}
