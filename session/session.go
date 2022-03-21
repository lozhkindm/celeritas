package session

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieName     string
	CookieDomain   string
	CookieSecure   string
	SessionType    string
	DBPool         *sql.DB
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
		se.Store = mysqlstore.New(s.DBPool)
	case "postgres", "postgresql":
		se.Store = postgresstore.New(s.DBPool)
	default:

	}

	return se
}
