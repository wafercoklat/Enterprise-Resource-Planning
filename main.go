package main

import (
	datasource "REVAMPS-PHP-GO/internal/datasources/sqlDB"
	"REVAMPS-PHP-GO/internal/domain/service"
	"REVAMPS-PHP-GO/internal/handlers"

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
