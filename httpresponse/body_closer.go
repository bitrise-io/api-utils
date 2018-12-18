package httpresponse

import (
	"log"
	"net/http"
)

// RequestBodyCloseWithErrorLog ...
func RequestBodyCloseWithErrorLog(r *http.Request) {
	err := r.Body.Close()
	if err != nil {
		log.Printf(" [!] Exception: RequestBodyCloseWithErrorLog: %+v", err)
	}
}
