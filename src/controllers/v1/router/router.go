package router

import (
	"log"
	"net/http"

	"projects/api.animalVids.com/pkg/types/routes"
	StatusHandler "projects/api.animalVids.com/src/controllers/v1/status"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")

		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

		log.Println("Inside V1 Middleware")
		next.ServeHTTP(w, r)
	})
}

// We are making a GetRoutes function for our subroutes we are using a map so we can easily use the key as the name later for our subroutes
// example the key of the map becomes the start of subroute ... /cubs
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
