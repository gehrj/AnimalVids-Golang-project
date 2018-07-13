package router

import (
	"log"
	"net/http"

	"projects/api.animalVids.com/pkg/types/routes"
	StatusHandler "projects/api.animalVids.com/src/controllers/v1/status"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("v1 middleware reached!")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (SubRoute map[string]routes.SubRoutePackage) {
	SubRoute = map[string]routes.SubRoutePackage{
		"/v1": routes.SubRoutePackage{
			Routes: routes.Routes{
				routes.Route{"Status", "GET", "/status", StatusHandler.Index},
			},
			Middleware: Middleware,
		},
	}
	return
}
