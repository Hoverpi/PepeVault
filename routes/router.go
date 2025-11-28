package routes

import (
	"PepeVault/handlers"
	"PepeVault/middlewares"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./static")

	// HTML Templates
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", handlers.LoginUser(db))
	router.GET("/register", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "register.html", nil)
	})

	protected := router.Group("/")
	protected.Use(middlewares.ValidateSession(db))
	{
		protected.GET("/vault/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "vault.html", gin.H{
				"ShowModal": false,
			})
		})
		protected.GET("/vault/new", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "vault.html", gin.H{
				"ShowModal": true,
			})
		})
		protected.POST("/vault/new", handlers.CreateVault(db))
	}

	return router
}
