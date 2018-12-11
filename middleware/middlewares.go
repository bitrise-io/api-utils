package middleware

import (
	"net/http"

	"github.com/slapec93/bitrise-api-utils/httpresponse"
)

// CreateRedirectToHTTPSMiddleware ...
func CreateRedirectToHTTPSMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			scheme := r.Header.Get("X-Forwarded-Proto")
			if scheme != "" && scheme != "https" {
				target := "https://" + r.Host + r.URL.Path
				if len(r.URL.RawQuery) > 0 {
					target += "?" + r.URL.RawQuery
				}
				http.Redirect(w, r, target, http.StatusPermanentRedirect)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

// CreateOptionsRequestTerminatorMiddleware ...
func CreateOptionsRequestTerminatorMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				httpresponse.RespondWithJSONNoErr(w, 200, nil)
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}
