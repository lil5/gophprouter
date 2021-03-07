package gophprouter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Handle func(httprouter.Params)

type Router struct {
	HTTPR *httprouter.Router
}

func NewRouter() *Router {
	r := Router{
		HTTPR: httprouter.New(),
	}

	return &r
}

func (r *Router) Serve(method string, path string, body string) bool {
	handle, params, _ := r.HTTPR.Lookup(method, path)

	if handle == nil {
		return false
	}

	bodyR := strings.NewReader(body)
	req := httptest.NewRequest(method, path, bodyR)
	w := httptest.NewRecorder()

	handle(w, req, params)

	resBody, _ := io.ReadAll(w.Body)

	fmt.Print(string(resBody))

	return true
}

func (r *Router) Handle(method string, path string, h httprouter.Handle) {
	//httpHandle := func(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	//	handle(params)
	//}
	r.HTTPR.Handle(method, path, h)
}

func (r *Router) DELETE(path string, h httprouter.Handle) {
	r.Handle(http.MethodDelete, path, h)
}
func (r *Router) GET(path string, h httprouter.Handle) {
	r.Handle(http.MethodGet, path, h)
}

func (r *Router) HEAD(path string, h httprouter.Handle) {
	r.Handle(http.MethodHead, path, h)
}
func (r *Router) OPTIONS(path string, h httprouter.Handle) {
	r.Handle(http.MethodOptions, path, h)
}
func (r *Router) PATCH(path string, h httprouter.Handle) {
	r.Handle(http.MethodPatch, path, h)
}
func (r *Router) POST(path string, h httprouter.Handle) {
	r.Handle(http.MethodPost, path, h)
}
func (r *Router) PUT(path string, h httprouter.Handle) {
	r.Handle(http.MethodPut, path, h)
}
