package providers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RequestParamsInterface ...
type RequestParamsInterface interface {
	Get(req *http.Request) map[string]string
}

// RequestParams ...
type RequestParams struct{}

// Get ...
func (r *RequestParams) Get(req *http.Request) map[string]string {
	return mux.Vars(req)
}

// RequestParamsMock ...
type RequestParamsMock struct {
	Params map[string]string
}

// Get ...
func (r *RequestParamsMock) Get(req *http.Request) map[string]string {
	return r.Params
}
