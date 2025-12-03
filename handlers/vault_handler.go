package handlers

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVault(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body struct {
			Service     string `form:"service" json:"service" binding:"required"`
			Description string `form:"description" json:"description" binding:"required"`
			URL         string `form:"url" json:"url" binding:"required"`
			Password    string `form:"password" json:"password" binding:"required"`
		}
		// Will bind depending on Content-Type (form or json). If required fields missing -> error.
		if err := ctx.ShouldBind(&body); err != nil {
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing data"})
			} else {
				ctx.HTML(http.StatusBadRequest, "vault.html", gin.H{"error": "Missing data"})
			}
			return
		}

	}
}
