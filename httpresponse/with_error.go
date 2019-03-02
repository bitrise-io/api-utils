package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// RespondWithJSON ...
func RespondWithJSON(w http.ResponseWriter, httpCode int, respModel interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(&respModel); err != nil {
		return errors.Wrapf(err, "Failed to respond (encode) with JSON for response model: %#v", respModel)
	}
	return nil
}

// RespondWithError ...
func RespondWithError(w http.ResponseWriter, errMsg string, httpErrCode int) error {
	return RespondWithJSON(w, httpErrCode, StandardErrorRespModel{
		Message: errMsg,
	})
}

// RespondWithSuccess ...
func RespondWithSuccess(w http.ResponseWriter, respModel interface{}) error {
	return RespondWithJSON(w, http.StatusOK, respModel)
}

// RespondWithBadRequestError ...
func RespondWithBadRequestError(w http.ResponseWriter, errMsg string) error {
	return RespondWithError(w, errMsg, http.StatusBadRequest)
}

// RespondWithUnauthorized ...
func RespondWithUnauthorized(w http.ResponseWriter) error {
	return RespondWithError(w, "Unauthorized", http.StatusUnauthorized)
}

// RespondWithForbidden ...
func RespondWithForbidden(w http.ResponseWriter) error {
	return RespondWithError(w, "Forbidden", http.StatusForbidden)
}

// RespondWithMethodNotAllowedErrorWithMessage ...
func RespondWithMethodNotAllowedErrorWithMessage(w http.ResponseWriter, errMsg string) error {
	return RespondWithError(w, errMsg, http.StatusMethodNotAllowed)
}

// RespondWithMethodNotAllowedError ...
func RespondWithMethodNotAllowedError(w http.ResponseWriter) error {
	return RespondWithMethodNotAllowedErrorWithMessage(w, "Method Not Allowed")
}

// RespondWithNotFoundErrorWithMessage ...
func RespondWithNotFoundErrorWithMessage(w http.ResponseWriter, errMsg string) error {
	return RespondWithError(w, errMsg, http.StatusNotFound)
}

// RespondWithNotFoundError ...
func RespondWithNotFoundError(w http.ResponseWriter) error {
	return RespondWithNotFoundErrorWithMessage(w, "Not Found")
}
