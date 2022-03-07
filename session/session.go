package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
}

func (s *Session) Init() *scs.SessionManager {
	var persist, secure bool

	minutes, err := strconv.Atoi(s.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	if strings.ToLower(s.CookiePersist) == "true" {
		persist = true
	}
	if strings.ToLower(s.CookieSecure) == "true" {
		secure = true
	}

	se := scs.New()
	se.Lifetime = time.Duration(minutes) * time.Minute
	se.Cookie.Persist = persist
	se.Cookie.Secure = secure
	se.Cookie.Name = s.CookieName
	se.Cookie.Domain = s.CookieDomain
	se.Cookie.SameSite = http.SameSiteLaxMode

	switch strings.ToLower(s.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:

	}

	return se
}
