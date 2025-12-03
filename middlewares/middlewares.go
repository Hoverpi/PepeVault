package middlewares

import (
	"PepeVault/auth"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateSession(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("session")
		if err != nil || cookie == "" {
			// Not authenticated
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			} else {
				ctx.Redirect(http.StatusSeeOther, "/login")
			}
			ctx.Abort()
			return
		}

		userID, err := auth.GetSessionUser(db, cookie)
		if err != nil {
			// if it's expired, remove it from DB and clear cookie
			if err == auth.ErrSessionExpired {
				_ = auth.DestroySession(ctx.Writer, db, cookie)
			} else {
				// could be sql.ErrNoRows or other DB error — ensure session deleted to be safe
				_, _ = db.Exec(`DELETE FROM public."sessions" WHERE "token" = $1`, cookie)
			}

			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			} else {
				ctx.Redirect(http.StatusSeeOther, "/login")
			}
			ctx.Abort()
			return
		}

		// Attach user id to context for handlers that need it
		ctx.Set("user_id", userID)

		ctx.Next()
	}
}
