package app

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
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

// Start the server by creating a router and then starting server on our port with newly created router
func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)

	r := router.NewRouter()

	r.Init()

	http.ListenAndServe(s.port, r.Router)
}
