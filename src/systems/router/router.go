package router

import (
	"github.com/gorilla/mux"
	"projects/api.animalVids.com/pkg/types/routes"
	V1SubRoutes "projects/api.animalVids.com/src/controllers/v1/router"
)

// we are creating a router struct to contain our router
type Router struct {
	Router *mux.Router
}

// this function is attaching our global middleware to our router and it is also looping through our routes and attaching them to our handler
// the chaining looks weird but it is essentially attaching our path and handler (route) to the router and attaching additional things gorilla mux can use (look at docs if forget their purpose)
// GetRoutes() and Middleware are coming from our router package this can be found at router/routes.go
func (r *Router) Init() {
	// use takes in a middleware function as its argument
	r.Router.Use(Middleware)

	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	// here we are getting our subroutes and then looping through them to attach them all to our main router
	// since these are maps we are looping through we are grabbing the name which turns into the path our subroutes begin with example /cubs
	v1SubRoutes := V1SubRoutes.GetRoutes()
	for name, pack := range v1SubRoutes {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
}

// remember when there is just a return at the bottom it is returning the variable you declared after function arguments
// for this function we are programatically creating our subroutes, path is where the subroute begins for example /cubs/..... so it is the cubs part of sub route
// subroutes then is all the routes that have the prefix /cubs so could be like /cubs/win, /cubs/areAwesome, /cubs/player/12 ect..
// the middleware is then the middleware that all of these subroutes share
// finally it returns the subrouter which is attached to our main router
func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return
}

// this is creating a router for us the router is contained in our router struct
func NewRouter() (r Router) {
	// strict slash is making example.com/user be the same as example.com/user/
	r.Router = mux.NewRouter().StrictSlash(true)
	return
}
