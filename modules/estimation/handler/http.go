package estimation

import (
	domain "STACK-ERP/modules/estimation/domain"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	port domain.Services
}

func NewHandler(route *gin.Engine, port domain.Services) {
	http := Handler{
		port: port,
	}

	// Routes
	route.GET("/estimation/:id/", http.EstimationFindByID)
	route.GET("/estimation/list/", http.EstimationList)
	route.POST("/estimation/add/", http.EstimationCreate)
	route.PUT("/estimation/update/:id/", http.EstimationUpdate)
	route.DELETE("/estimation/delete/:id/", http.EstimationDelete)
}

func (h *Handler) EstimationFindByID(c *gin.Context) {
	val := h.port.FindByID(c.Param("id"))
	c.JSON(200, val)
}

func (h *Handler) EstimationList(c *gin.Context) {
	val := h.port.List()
	c.JSON(200, val)
}

func (h *Handler) EstimationCreate(c *gin.Context) {
	var request domain.Estimation
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val := h.port.Create(request)
	c.JSON(200, val)
}

func (h *Handler) EstimationUpdate(c *gin.Context) {
	var request domain.Estimation
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val := h.port.Update(c.Param("id"), request)
	c.JSON(200, val)
}

func (h *Handler) EstimationDelete(c *gin.Context) {
	val := h.port.Delete(c.Param("id"))
	c.JSON(200, val)
}
