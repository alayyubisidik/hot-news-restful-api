package middleware

import "github.com/julienschmidt/httprouter"

func ChainMiddleware(h httprouter.Handle, middlewares ...func(httprouter.Handle) httprouter.Handle) httprouter.Handle {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}