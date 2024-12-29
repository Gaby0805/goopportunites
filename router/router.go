package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	// Inicialize Router
	router := gin.Default()
	// Inicialize Routes
	inicializateRouter(router)

	// Run server
	router.Run(":8080")
}
