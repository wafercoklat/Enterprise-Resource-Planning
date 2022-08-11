package handler

import (
	"REVAMP-PHP-GO/internal/domain/ports"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	port ports.PortServices
}

func New(port ports.PortServices) *Handler {
	return &Handler{
		port: port,
	}
}

func (h *Handler) AccountFindByID(c *gin.Context) {
	val := h.port.AccountFindByID(c.Param("id"))
	c.JSON(200, val)
}

func (h *Handler) AccountList(c *gin.Context) {
	val := h.port.AccountList()
	c.JSON(200, val)
}

func (h *Handler) AccountCreate(c *gin.Context) {
	//Get data
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	// Set Data
	val := h.port.AccountCreate(request)
	c.JSON(200, val)
}

func (h *Handler) AccountUpdate(c *gin.Context) {
	//Get data
	var request map[string]interface{}
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	// Set Data
	val := h.port.AccountUpdate(c.Param("id"), request)
	c.JSON(200, val)
}

func (h *Handler) AccountDelete(c *gin.Context) {
	val := h.port.AccountDelete(c.Param("id"))
	c.JSON(200, val)
}
