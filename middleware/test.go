package middleware

import (
	"net/http"

	"github.com/bitrise-io/api-utils/httpresponse"
)

// TestHandler ...
func TestHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		httpresponse.RespondWithSuccessNoErr(rw, map[string]string{"message": "Success"})
	}
	return http.HandlerFunc(fn)
}
