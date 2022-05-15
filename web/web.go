package web

import "net/http"

type Controller func(w http.ResponseWriter, r *http.Request)

type Middleware func(controller Controller) (wrapped Controller)

func (c Controller) With(middlewares ...Middleware) Controller {
	wrapped := c
	for _, midWare := range middlewares {
		wrapped = midWare(wrapped)
	}

	return wrapped
}
