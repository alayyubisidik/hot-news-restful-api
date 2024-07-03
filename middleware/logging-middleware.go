package middleware

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		next(w, r, ps)
	}
}
 