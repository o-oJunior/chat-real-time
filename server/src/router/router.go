package router

import (
	"net/http"
	"server/src/logger"

	"github.com/gorilla/mux"
)

type Router struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routerConfig(r)
}

func routerConfig(r *mux.Router) *mux.Router {
	routers := userRouters
	for _, router := range routers {
		logger.Info("[ROUTER (%s - %s)] Rota definida!", router.Method, router.URI)
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}
	return r
}
