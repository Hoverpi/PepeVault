package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVault(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body struct {
			service     string `form:"uuid_v4" json:"uuid_v4" binding:"required"`
			description string `form:"uuid_v4" json:"uuid_v4" binding:"required"`
			url         string `form:"uuid_v4" json:"uuid_v4" binding:"required"`
			password    string `form:"uuid_v4" json:"uuid_v4" binding:"required"`
		}
		log.Println(body.service, body.description, body.url, body.password)
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
