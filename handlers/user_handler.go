package handlers

import (
	"PepeVault/auth"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// https://www.geeksforgeeks.org/html/http-headers-content-type/

func LoginUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body struct {
			Username  string `form:"username" json:"username" binding:"required"`
			Password string `form:"password" json:"password" binding:"required"`
		}
		// Will bind depending on Content-Type (form or json). If required fields missing -> error.
		if err := ctx.ShouldBind(&body); err != nil {
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing username or password"})
			} else {
				ctx.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Missing username or password"})
			}
			return
		}

		var userID string
		var storedPassword string
		err := db.QueryRow(`SELECT "id", "password_hash" FROM public."users" WHERE "username" = $1`, body.Username).Scan(&userID, &storedPassword)

		if err != nil {
			if err == sql.ErrNoRows { // QueryRow return it
				if ctx.ContentType() == "application/json" {
					ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
				} else {
					ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
				}
				return
			}
			// other DB error
			log.Println("login db error:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// TODO: Compare master password
		// if err := body.Username == Password; err != nil {}
		if body.Password != storedPassword {
			// password mismatch
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			} else {
				ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
			}
			return
		}

		// Create session cookie
		const sessionTTL = 24 * time.Hour
		token, err := auth.CreateSession(ctx.Writer, ctx.Request, db, userID, sessionTTL)
		if err != nil {
			log.Println("create session error:", err)
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			} else {
				ctx.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Internal server error"})
			}
			return
		}

		// Use the token so it is not an unused variable; return for API, redirect for form
		if ctx.ContentType() == "application/json" {
			ctx.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
		} else {
			ctx.Redirect(http.StatusSeeOther, "/vault")
		}
	}
}

func RegisterUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body struct {
			Username  string `form:"username" json:"username" binding:"required"`
			Password string `form:"password" json:"password" binding:"required"`
			Confirm_Password string `form:"confirm-password" json:"confirm-password" binding:"required"`
		}
		// Will bind depending on Content-Type (form or json). If required fields missing -> error.
		if err := ctx.ShouldBind(&body); err != nil {
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing username or password"})
			} else {
				ctx.HTML(http.StatusBadRequest, "login.html", gin.H{"error": "Missing username or password"})
			}
			return
		}

		var userID string
		var storedPassword string
		err := db.QueryRow(`SELECT "id", "password_hash" FROM public."users" WHERE "username" = $1`, body.Username).Scan(&userID, &storedPassword)

		if err != nil {
			if err == sql.ErrNoRows { // QueryRow return it
				if ctx.ContentType() == "application/json" {
					ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
				} else {
					ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
				}
				return
			}
			// other DB error
			log.Println("login db error:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// TODO: Compare master password
		// if err := body.Username == Password; err != nil {}
		if body.Password != storedPassword {
			// password mismatch
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			} else {
				ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"error": "Invalid credentials"})
			}
			return
		}

		// Create session cookie
		const sessionTTL = 24 * time.Hour
		token, err := auth.CreateSession(ctx.Writer, ctx.Request, db, userID, sessionTTL)
		if err != nil {
			log.Println("create session error:", err)
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			} else {
				ctx.HTML(http.StatusInternalServerError, "login.html", gin.H{"error": "Internal server error"})
			}
			return
		}

		// Use the token so it is not an unused variable; return for API, redirect for form
		if ctx.ContentType() == "application/json" {
			ctx.JSON(http.StatusOK, gin.H{"ok": true, "token": token})
		} else {
			ctx.Redirect(http.StatusSeeOther, "/vault")
		}
	}
}