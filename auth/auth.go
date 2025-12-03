package auth

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"net"
	"log"
	"net/http"
	"time"
)

// ErrSessionExpired is returned when a session exists but is past its expiry.
var ErrSessionExpired = errors.New("session expired")

// GenerateSecureToken returns a URL-safe base64 token
func GenerateSecureToken(nbytes int) (string, error) {
	b := make([]byte, nbytes)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// CreateSession creates a session token, stores it in DB (with ip + ua), and sets cookie.
// r is used to extract IP and User-Agent.
func CreateSession(w http.ResponseWriter, r *http.Request, db *sql.DB, userID string, ttl time.Duration) (string, error) {
	token, err := GenerateSecureToken(32)
	if err != nil {
		return "", err
	}
	expires := time.Now().Add(ttl)

	// extract IP (handle X-Forwarded-For)
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		// r.RemoteAddr may include port, so strip it
		host, _, err2 := net.SplitHostPort(r.RemoteAddr)
		if err2 == nil {
			ip = host
		} else {
			ip = r.RemoteAddr
		}
	}
	userAgent := r.UserAgent()

	// Insert session
	_, err = db.Exec(
		`INSERT INTO public."sessions" ("token", "user_id", "ip_address", "user_agent", "expires_at") VALUES ($1, $2, $3, $4, $5)`,
		token, userID, ip, userAgent, expires,
	)
	log.Println("ERROR 1=", err)
	if err != nil {
		return "", err
	}

	// set cookie
	cookie := &http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		Expires:  expires,
		MaxAge:   int(ttl.Seconds()),
		HttpOnly: true,
		Secure:   false, // set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	return token, nil
}

// GetSessionUser returns user_id (UUID string) if token valid, else error.
// Returns auth.ErrSessionExpired if found but expired.
func GetSessionUser(db *sql.DB, token string) (string, error) {
	if token == "" {
		return "", errors.New("no token")
	}
	var userID string
	var expires time.Time
	err := db.QueryRow(`SELECT "user_id", "expires_at" FROM public."sessions" WHERE "token" = $1`, token).Scan(&userID, &expires)
	log.Println("ERROR 2=", err)
	if err != nil {
		return "", err
	}
	if time.Now().After(expires) {
		return "", ErrSessionExpired
	}
	return userID, nil
}

// DestroySession - delete token from DB and clear cookie
func DestroySession(w http.ResponseWriter, db *sql.DB, token string) error {
	_, _ = db.Exec(`DELETE FROM public."sessions" WHERE "token" = $1`, token) // ignore DB error here intentionally
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
