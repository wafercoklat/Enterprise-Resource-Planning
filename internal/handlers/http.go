// handler to endPoint
package handlers

import (
	"REVAMPS-CMI-APPS/internal/entity/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service ports.PortService
}

func NewHTTP(services ports.PortService) *HTTPHandler {
	return &HTTPHandler{
		service: services,
	}
}

func (hdl *HTTPHandler) Retreive(c *gin.Context) {
	api, err := hdl.service.Retreive(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, api)
}
