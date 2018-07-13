package router

import (
	"log"
	"net/http"

	"projects/api.animalVids.com/pkg/types/routes"
	HomeHandler "projects/api.animalVids.com/src/controllers/home"
)

// this is the genral design of a middleware function, we can add logic inside to do whaterver we want with route before moving on
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware reached!")
		next.ServeHTTP(w, r)
	})
}

// this is using our Routes slice and Route struct from /pkg/types/routes
// we can build each route here which will be looped over in router.go to quickly build all of our routes
// HomeHandler is coming from our controllers/home it is basically a place we are creating our handlers in another file
func GetRoutes() routes.Routes {
	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}

}
