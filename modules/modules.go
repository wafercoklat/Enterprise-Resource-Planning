package modules

import (
	Acchandl "STACK-ERP/modules/account/handler"
	Accserv "STACK-ERP/modules/account/service"
	Esthandl "STACK-ERP/modules/estimation/handler"
	Estserv "STACK-ERP/modules/estimation/service"
	"STACK-ERP/port"

	"github.com/gin-gonic/gin"
)

func New(g *gin.Engine, db port.PortRepo) {
	Acchandl.NewHandler(g, Accserv.NewService(db)) // Account Module
	Esthandl.NewHandler(g, Estserv.NewService(db)) // Estimation Module

}
