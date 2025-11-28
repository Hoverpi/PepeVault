package middlewares

import (
	"database/sql"
	"net/http"
	"time"

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

		var UserID string
		var expires time.Time
		err = db.QueryRow(`SELECT "User_ID", "Expires_at" FROM public."Sessions" WHERE "Token" = $1`, cookie).Scan(&UserID, &expires)
		if err != nil || time.Now().After(expires) {
			// invalid or expired session
			_, _ = db.Exec(`DELETE FROM sessions WHERE token = $1`, cookie)
			if ctx.ContentType() == "application/json" {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated"})
			} else {
				ctx.Redirect(http.StatusSeeOther, "/login")
			}
			ctx.Abort()
			return
		}

		ctx.Next()

	}
}
