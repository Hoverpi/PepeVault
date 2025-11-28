package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"time"
)

// GenerateSecureToken returns a URL-safe base64 token
func GenerateSecureToken(nbytes int) (string, error) {
	b := make([]byte, nbytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// CreateSession creates a session token, stores in DB, and sets cookie
func CreateSession(w http.ResponseWriter, db *sql.DB, userID string, ttl time.Duration) (string, error) {
	token, err := GenerateSecureToken(32)
	if err != nil {
		return "", err
	}
	expires := time.Now().Add(ttl)

	// upsert into sessions table
	// User table: INSERT INTO public."User" ("Username", "Password") VALUES ('{"Admin"}', '{"Admin"}');
	// Session table: INSERT INTO public."Session" ("Token", "User_ID", "Expires_at") VALUES ('{"Admin"}', '{"Admin"}');
	_, err = db.Exec(`INSERT INTO public."Sessions" ("Token", "User_ID", "Expires_at") VALUES ($1, $2, $3)`, token, userID, expires)
	if err != nil {
		return "", err
	}

	cookie := &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		Expires:  expires,
		MaxAge:   int(ttl.Seconds()),
		HttpOnly: true,
		Secure:   false, // set to true in production (requires HTTPS)
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	return token, nil
}

// GetSessionUser returns user_id if token valid, else error
func GetSessionUser(db *sql.DB, token string) (int64, error) {
	if token == "" {
		return 0, errors.New("no token")
	}
	var userID int64
	var expires time.Time
	err := db.QueryRow(`SELECT "User_ID", "Expires_at" FROM public."Sessions" WHERE Token = $1`, token).Scan(&userID, &expires)
	if err != nil {
		return 0, err
	}
	if time.Now().After(expires) {
		return 0, errors.New("expired")
	}
	return userID, nil
}

// DestroySession - delete token from DB and clear cookie
func DestroySession(w http.ResponseWriter, db *sql.DB, token string) error {
	// User table: DELETE FROM public."User" WHERE "Username" = 'Pepe';
	// Sessions table: DELETE FROM public."Sessions" WHERE "Token" = '';
	_, _ = db.Exec(`DELETE FROM public."Sessions" WHERE "Token" = $1`, token) // TODO: ignore error for now
	// clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false, // set to true in production (requires HTTPS)
		SameSite: http.SameSiteLaxMode,
	})
	return nil
}
