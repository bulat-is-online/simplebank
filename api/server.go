package api

import (
	db "github.com/bulat-is-online/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves Http requst to banking service
type Server struct {
	store  db.Store    //will allow to interact with db
	router *gin.Engine //will help to send each API request to correct handler
}

// NewServer creates a new Http server and setups routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store} //creating new server
	router := gin.Default()         // calling new router

	//adding custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//route, one or multiple handlers, the last function should be real handler, and others middlewares
	// need to implement method for the server struct because to have an access to server object to save new account in db
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.PUT("/accounts", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	// in list /accounts because parameter are sent in url
	router.GET("/accounts", server.listAccount)

	//same for transfers
	router.POST("/transfers", server.createTransfer)

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
