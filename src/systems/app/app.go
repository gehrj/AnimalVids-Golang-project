package app

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
	"projects/api.animalVids.com/src/systems/router"
)

// this is our server struct it holds our port and the db our app will be using
type Server struct {
	port string
	Db   *xorm.Engine
}

// this function creates a server for us
func NewServer() Server {
	return Server{}
}

// Init all vals for our server struct
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initalizing server...")
	s.port = ":" + port
	s.Db = db
}

// Start the server by creating a router and then creating a handler that also takes in our router which combined becomes our handler we need for our server that
// has our routes and also is able to accept requests from outside sources. We then make a server because this allows for more control over behavior.
// once our server is made we call ListenAndServe on it to start our server.
func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)

	// we are grabbing a new router here
	r := router.NewRouter()

	// this is where we initialize our router
	r.Init()

	// this is for cross-origin resource sharing, since our api will be needing to communicate with things outside of itself this is necessary
	// to let it know what kind of requests it should accept. the handlers.Cors(options)(r.Router) does this for us. Putting it inside the logginghandler lets us
	// get information in our terminal about the cors.
	// LogginHandler looks like this func LogginHandler(out io.Writer, h http.Handler) http.Handler
	// so that is why it can take handlers.CORS as its second argument
	// basically we are creating a ton of handlers that all turn into one handler in our assignment (http.Handler)
	handler := handlers.LoggingHandler(os.Stdout, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "PATCH", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "Cache-Control", "X-App-Token"}),
		handlers.ExposedHeaders([]string{""}),
		handlers.MaxAge(1000),
		handlers.AllowCredentials(),
	)(r.Router))
	// handlers.RecoveryHandler is a function with options that returns a function that expects and http.Handler argument and returns a new http.Handler
	// that is why it looks like a double function call same thing as what handlers.CORS looks like it follows same process
	handler = handlers.RecoveryHandler(handlers.PrintRecoveryStack(true))(handler)

	// this is creating a server that lets us have more control than just doing http.ListenAndServe(s.port, handler)
	// using "0.0.0.0" will match all ip addresses currently assigned to device :D
	newServer := &http.Server{
		Handler:      handler,
		Addr:         "0.0.0.0" + s.port,
		WriteTimeout: 15 + time.Second,
		ReadTimeout:  15 + time.Second,
	}

	// here we are starting our server just as if we did http.ListenAndServe(s.port, handler)
	log.Fatal(newServer.ListenAndServe())
}
