package main

import (
	"flag"

	"projects/api.animalVids.com/src/systems/app"
	DB "projects/api.animalVids.com/src/systems/db"
)

// these are the variables that we are setting with our flags in the init function
var port string
var dbhost string
var dbport string
var dbuser string
var dbpass string
var dboptions string
var dbdatabase string

// this function runs before main runs and it has a bunch of flags in it. These allow the user to set the variables in the command line, if they do not set them they will be set to the
// default variable which is the second argument. for example port will be set to 8080 if they do not set a port in command line
func init() {
	flag.StringVar(&port, "port", "8080", "Assigning the port that the server should listen on.")
	flag.StringVar(&dbhost, "dbhost", "127.0.0.1", "Set the host for the application")
	flag.StringVar(&dbport, "dbport", "3306", "Set the dbport for the application")
	flag.StringVar(&dbuser, "dbuser", "test", "Set the user for the application")
	flag.StringVar(&dbpass, "dbpass", "test", "Set the password for the application")
	flag.StringVar(&dboptions, "dboptions", "parseTime=true", "Set the options for the application")
	flag.StringVar(&dbdatabase, "dbdatabase", "animal_vids", "Set the database for the application")

	flag.Parse()
}

// the main function of our api application, this connects to our database and initializes and starts our server
func main() {
	// this is where we are connecting to the server and passing in the varables we set by default or with our flag entries
	db, err := DB.Connect(dbhost, dbport, dbuser, dbpass, dbdatabase, dboptions)
	if err != nil {
		panic(err)
	}
	// this method off of app grabs a server struct and assigns it to s
	s := app.NewServer()

	// Init is a method off of our server struct we are setting our port and db variables on our server struct
	// Start is another method off of our server struct it is creating our router and and then starting our server on the proper server with the router passed in
	s.Init(port, db)
	s.Start()
}
