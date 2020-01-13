package middleware

// This was copied from galera-healthcheck in https://github.com/cloudfoundry/cf-mysql-release

import (
	"crypto/subtle"
	"net/http"
)

type BasicAuth struct {
	Username, Password string
}

func NewBasicAuth(username, password string) Middleware {
	return BasicAuth{
		Username: username,
		Password: password,
	}
}

func (b BasicAuth) Wrap(next http.Handler) http.Handler {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if ok &&
			secureCompare(username, b.Username) &&
			secureCompare(password, b.Password) {
			next.ServeHTTP(rw, req)
		} else {
			rw.Header().Set("WWW-Authenticate", "Basic realm=\"Authorization Required\"")
			http.Error(rw, "Not Authorized", http.StatusUnauthorized)
		}
	})
	return handler
}

func secureCompare(a, b string) bool {
	x := []byte(a)
	y := []byte(b)
	return subtle.ConstantTimeCompare(x, y) == 1
}
