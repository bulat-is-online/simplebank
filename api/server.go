package api

import (
	db "github.com/bulat-is-online/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves Http requst to banking service
type Server struct {
	store  *db.Store   //will allow to interact with db
	router *gin.Engine //will help to send each API request to correct handler
}

// NewServer creates a new Http server and setups routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store} //creating new server
	router := gin.Default()         // calling new router

	//route, one or multiple handlers, the last function should be real handler, and others middlewares
	// need to implement method for the server struct because to have an access to server object to save new account in db
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	// in list /accounts because parameter are sent in url
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs Http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
