package main

import (
	datasources "REVAMP-PHP-GO/internal/datasources/sqlDB"
	"REVAMP-PHP-GO/internal/domain/services"
	"REVAMP-PHP-GO/internal/handler"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO Database
	dialect := "mysql"
	dsn := "root:@tcp(localhost:3306)/enterprise022?parseTime=true"
	datasources, err := datasources.New(dialect, dsn, 1, 1)
	if err != nil {
		fmt.Printf("%s", err)
	}

	// TODO Services
	service := services.New(datasources)

	// TODO Handler
	handler := handler.New(service)

	// TODO Routes
	r := gin.Default()
	r.GET("/account/:id/", handler.AccountFindByID)
	r.GET("/account/list/", handler.AccountList)
	r.POST("/account/add/", handler.AccountCreate)
	r.PUT("/account/update/:id/", handler.AccountUpdate)
	r.DELETE("/account/delete/:id/", handler.AccountDelete)

	if err := r.Run(":8099"); err != nil {
		panic(err.Error())
	}

}
