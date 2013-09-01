package main

import (
	"fmt"
	"github.com/dancannon/gonews/app"
	"github.com/dancannon/gonews/controllers"
	"github.com/dancannon/gonews/util"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer(mode string) *http.Server {
	// Create application
	app.Init(mode)

	// Setup router
	app.SetRouter(initRouting())

	// Setup logging handler
	loggingHandler := handlers.CombinedLoggingHandler(util.GetLogFile("access.log"), app.Router)

	// Create and start server
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", "3000"),
		Handler: loggingHandler,
	}
}

func StartServer(server *http.Server) {
	server.ListenAndServe()
}

func initRouting() *mux.Router {
	r := mux.NewRouter()

	new(controllers.IndexController).Init(r.PathPrefix("/").Subrouter())
	new(controllers.PostController).Init(r.PathPrefix("/posts").Subrouter())

	// Add handler for static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return r
}
