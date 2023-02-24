package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "simple_bank/db/sqlc"
)

type InitiateTransferHandlerRequest struct {
	SenderID   int64 `json:"sender_id" binding:"required,min=1"`
	ReceiverID int64 `json:"receiver_id" binding:"required,min=1"`
	Amount     int64 `json:"amount" binding:"required,min=1"`
}

// InitiateTransferHandler / @Make Transfer
// @Tags Transfers
// @Accept json
// @Produce  json
// @Param requestBody body InitiateTransferHandlerRequest true "requestBody"
// @Success 200 {object} db.TransferTXResult "ok"
// @Failure 400 {object} ServerError  "Something went wrong"
// @Failure 500 {object} ServerError  "Something went wrong"
// @Router /api/v1/transfer/initiatetransfer [POST]
func (s *Server) InitiateTransferHandler(context *gin.Context) {
	var request InitiateTransferHandlerRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	result, err := s.store.TransferTX(context, db.TransferTXParam{
		SenderID:   request.SenderID,
		ReceiverID: request.ReceiverID,
		Amount:     request.Amount,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	context.JSON(http.StatusOK, result)
}

type GetIncomingTransferHandlerRequest struct {
	ToAccountID int64 `json:"to_account_id" binding:"required,min=1"`
	PageSize    int32 `form:"page_size" binding:"required,min=1"`
	PageID      int32 `form:"page_id" binding:"required,min=1"`
}

// GetIncomingTransferHandler  / @Get Incoming Transfers
// @Tags Transfers
// @Accept json
// @Produce  json
// @Param page_id query int32 true "pageID"
// @Param page_size query int32 true "pageSize"
// @Param to_account_id body int64 true "to_account_id"
// @Success 200 {object} []db.Transfer "ok"
// @Failure 400 {object} ServerError  "Something went wrong"
// @Failure 500 {object} ServerError  "Something went wrong"
// @Router /api/v1/transfer/getincomingtransfers [GET]
func (s *Server) GetIncomingTransferHandler(context *gin.Context) {
	var request GetIncomingTransferHandlerRequest
	err := context.ShouldBindQuery(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	err = context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	transfers, err := s.store.ListIncomingTransfers(context, db.ListIncomingTransfersParams{
		ToAccountID: request.ToAccountID,
		Limit:       request.PageSize,
		Offset:      (request.PageID - 1) * request.PageSize,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	context.JSON(http.StatusOK, transfers)
}

type GetOutgoingTransferHandlerRequest struct {
	FromAccountID int64 `json:"from_account_id" binding:"required,min=1"`
	PageSize      int32 `form:"page_size" binding:"required,min=1"`
	PageID        int32 `form:"page_id" binding:"required,min=1"`
}

// GetOutgoingTransferHandler  / @Get Outgoing Transfers
// @Tags Transfers
// @Accept json
// @Produce  json
// @Param page_id query int32 true "pageID"
// @Param page_size query int32 true "pageSize"
// @Param from_account_id body int64 true "from_account_id"
// @Success 200 {object} []db.Transfer "ok"
// @Failure 400 {object} ServerError  "Something went wrong"
// @Failure 500 {object} ServerError  "Something went wrong"
// @Router /api/v1/transfer/getoutgoingtransfers [GET]
func (s *Server) GetOutgoingTransferHandler(context *gin.Context) {
	var request GetOutgoingTransferHandlerRequest
	err := context.ShouldBindQuery(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	err = context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	transfers, err := s.store.ListOutgoingTransfers(context, db.ListOutgoingTransfersParams{
		FromAccountID: request.FromAccountID,
		Limit:         request.PageSize,
		Offset:        (request.PageID - 1) * request.PageSize,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	context.JSON(http.StatusOK, transfers)

}

type GetTransfersBetweenHandlerRequest struct {
	Account1 int64 `json:"account_1" binding:"required,min=1"`
	Account2 int64 `json:"account_2" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}

func (s *Server) GetTransfersBetweenHandler(context *gin.Context) {
	var request GetTransfersBetweenHandlerRequest
	err := context.ShouldBindQuery(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	err = context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	transfers, err := s.store.ListTransfersBetween(context, db.ListTransfersBetweenParams{
		Account1: request.Account1,
		Account2: request.Account2,
		Limit:    request.PageSize,
		Offset:   (request.PageID - 1) * request.PageSize,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	context.JSON(http.StatusOK, transfers)
}

type GetTransferHandlerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetTransferHandler  / @Get a Transfer
// @Tags Transfers
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} db.Transfer "ok"
// @Failure 400 {object} ServerError  "Something went wrong"
// @Failure 500 {object} ServerError  "Something went wrong"
// @Router /api/v1/transfer/getTransfer/{id} [GET]
func (s *Server) GetTransferHandler(context *gin.Context) {
	var request GetTransferHandlerRequest
	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	transfer, err := s.store.GetTransfer(context, request.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	context.JSON(http.StatusOK, transfer)
}
