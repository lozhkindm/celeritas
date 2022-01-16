package session

import (
	"github.com/alexedwards/scs/v2"
	"reflect"
	"testing"
)

func TestSession_Init(t *testing.T) {
	var (
		sm    *scs.SessionManager
		sKind reflect.Kind
		sType reflect.Type
	)

	s := &Session{
		CookieLifetime: "111",
		CookiePersist:  "true",
		CookieName:     "test",
		CookieDomain:   "test_domain",
		CookieSecure:   "false",
		SessionType:    "cookie",
	}

	session := s.Init()
	rv := reflect.ValueOf(session)

	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		sKind = rv.Kind()
		sType = rv.Type()
		rv = rv.Elem()
	}

	if !rv.IsValid() {
		t.Errorf("invalid kind or type; kind: %v, type %v", rv.Kind(), rv.Type())
	}

	if sKind != reflect.ValueOf(sm).Kind() {
		t.Errorf("wrong kind; expected %v, got %v", reflect.ValueOf(sm).Kind(), sKind)
	}

	if sType != reflect.ValueOf(sm).Type() {
		t.Errorf("wrong type; expected %v, got %v", reflect.ValueOf(sm).Type(), sType)
	}
}
