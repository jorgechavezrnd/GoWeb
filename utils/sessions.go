package utils

import (
	"net/http"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.com/jorgechavezrnd/go_rest/models"
)

const (
	cookieName   = "go_session"
	cookieExpire = 24 * 2 * time.Hour // dos dias
)

// Sessions ...
var Sessions = struct {
	m map[string]*models.User
	sync.RWMutex
}{m: make(map[string]*models.User)}

// SetSession ...
func SetSession(user *models.User, w http.ResponseWriter) {
	Sessions.Lock()
	defer Sessions.Unlock()

	uuidValue, _ := uuid.NewV4()
	uuid := uuidValue.String()
	Sessions.m[uuid] = user

	cookie := &http.Cookie{
		Name:    cookieName,
		Value:   uuid,
		Path:    "/",
		Expires: time.Now().Add(cookieExpire),
	}

	http.SetCookie(w, cookie)
}

// GetUser ...
func GetUser(r *http.Request) *models.User {
	Sessions.Lock()
	defer Sessions.Unlock()

	uuid := getValCookie(r)
	if user, ok := Sessions.m[uuid]; ok {
		return user
	}

	return &models.User{}
}

// DeleteSession ...
func DeleteSession(w http.ResponseWriter, r *http.Request) {
	Sessions.Lock()
	defer Sessions.Unlock()

	delete(Sessions.m, getValCookie(r))

	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

func getValCookie(r *http.Request) string {
	if cookie, err := r.Cookie(cookieName); err == nil {
		return cookie.Value // uuid
	}
	return ""
}

// IsAuthenticated ...
func IsAuthenticated(r *http.Request) bool {
	return getValCookie(r) != ""
}
