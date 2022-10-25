package user

import (
	domain "STACK-ERP/modules/account/domain"

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
	route.GET("/account/:id/", http.AccountFindByID)
	route.GET("/account/list/", http.AccountList)
	route.POST("/account/add/", http.AccountCreate)
	route.PUT("/account/update/:id/", http.AccountUpdate)
	route.DELETE("/account/delete/:id/", http.AccountDelete)
}

func (h *Handler) AccountFindByID(c *gin.Context) {
	val := h.port.FindByID(c.Param("id"))
	c.JSON(200, val)
}

func (h *Handler) AccountList(c *gin.Context) {
	val := h.port.List()
	c.JSON(200, val)
}

func (h *Handler) AccountCreate(c *gin.Context) {
	var request domain.Account
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val := h.port.Create(request)
	c.JSON(200, val)
}

func (h *Handler) AccountUpdate(c *gin.Context) {
	var request domain.Account
	if err := c.ShouldBindJSON(&request); err != nil {
		panic(err)
	}

	val := h.port.Update(c.Param("id"), request)
	c.JSON(200, val)
}

func (h *Handler) AccountDelete(c *gin.Context) {
	val := h.port.Delete(c.Param("id"))
	c.JSON(200, val)
}
