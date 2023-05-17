package api

import (
	"fmt"

	db "github.com/bulat-is-online/simplebank/db/sqlc"
	"github.com/bulat-is-online/simplebank/db/token"
	"github.com/bulat-is-online/simplebank/db/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves Http requst to banking service
type Server struct {
	config     util.Config
	store      db.Store //will allow to interact with db
	tokenMaker token.Maker
	router     *gin.Engine //will help to send each API request to correct handler
}

// NewServer creates a new Http server and setups routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("config.TokenSymmetricKey")
	if err != nil {
		return nil, fmt.Errorf("cannot create tokenmaker %w", err)
	}

	//creating new ser
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	//adding custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default() // calling new router
	//users
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	//route, one or multiple handlers, the last function should be real handler, and others middlewares
	// need to implement method for the server struct because to have an access to server object to save new account in db
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.PUT("/accounts", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	// in list /accounts because parameter are sent in url
	router.GET("/accounts", server.listAccount)

	//transfers
	router.POST("/transfers", server.createTransfer)
	server.router = router
}

// Start runs Http server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
