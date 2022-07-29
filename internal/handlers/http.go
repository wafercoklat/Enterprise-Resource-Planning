// handler to endPoint
package handlers

import (
	"REVAMPS-PHP-GO/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	service ports.PortService
}

func NewHTTP(getService ports.PortService) *HTTPHandler {
	return &HTTPHandler{
		service: getService,
	}
}

func (h *HTTPHandler) ViewDataById(c *gin.Context) {
	val, err := h.service.ViewDataById(c.Param("id"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, val)
}
