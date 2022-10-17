package main

import (
	"context"
	"github.com/justinas/nosurf"
	"net/http"
)

func CSRFCheck(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func IsAuthenticated(next http.Handler) http.Handler {
	if app.IsLogin && app.SessionManager.GetBool(context.Background(), "user_id") {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}

	return nil
}
