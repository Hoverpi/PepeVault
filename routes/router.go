package routes

import (
	"PepeVault/handlers"
	"PepeVault/middlewares"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")

	// HTML Templates
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	})
	router.POST("/login", handlers.LoginUser(db))
	router.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	})

	protected := router.Group("/")
	protected.Use(middlewares.ValidateSession(db))
	{
		protected.GET("/panel", handlers.PanelHandler)

		// Vault Functions
		// protected.GET("/panel/vault", handlers.GetVaults)
	}

	return router
}
