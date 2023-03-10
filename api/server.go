package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	db "simple_bank/db/sqlc"
	docs "simple_bank/docs"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

type ServerError struct {
	Error   error  `json:"error"`
	Message string `json:"message"`
}

func NewServerError(err error) ServerError {
	return ServerError{Error: err, Message: err.Error()}
}

func NewServer(store *db.Store) *Server {
	server := Server{store: store}

	// Create the routers and register the rout s
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	v1 := router.Group("/api/v1")
	account := v1.Group("/account")
	account.POST("/createaccount", server.CreateAccountHandler)
	account.GET("/getaccount/:id", server.GetAccountByIDHandler)
	account.GET("/getaccounts", server.GetAccountsHandler)
	account.PUT("/updateaccount", server.UpdateAccountHandler)
	account.DELETE("/deleteaccount/:id", server.DeleteAccountHandler)

	transfer := v1.Group("/transfer")
	transfer.POST("/initiatetransfer", server.InitiateTransferHandler)
	transfer.GET("/getincomingtransfers", server.GetIncomingTransferHandler)
	transfer.GET("/getoutgoingtransfers", server.GetOutgoingTransferHandler)
	transfer.GET("/getTransfersBetween", server.GetTransfersBetweenHandler)
	transfer.GET("/getTransfer/:id", server.GetTransferHandler)

	logger := v1.Group("/log")
	logger.GET("/getentry/:id", server.GetEntryHandler)
	logger.GET("/listaccountentries", server.GetAccountEntriesHandler)
	logger.GET("/listBankEntries", server.GetBankEntriesHandler)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	server.router = router

	return &server
}

func (s *Server) Start(address string) {
	log.Fatal(s.router.Run(address))
}
