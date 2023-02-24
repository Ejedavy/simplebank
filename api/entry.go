package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "simple_bank/db/sqlc"
)

type GetEntryHandlerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetEntryHandler(context *gin.Context) {
	var request GetEntryHandlerRequest
	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	entry, err := s.store.GetEntry(context, request.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	context.JSON(http.StatusOK, entry)
}

type GetAccountEntriesHandlerRequest struct {
	AccountID int64 `json:"account_id" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=1"`
	PageID    int32 `form:"page_id" binding:"required,min=1"`
}

func (s *Server) GetAccountEntriesHandler(context *gin.Context) {
	var request GetAccountEntriesHandlerRequest
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
	entries, err := s.store.ListAccountEntries(context, db.ListAccountEntriesParams{
		AccountID: request.AccountID,
		Limit:     request.PageSize,
		Offset:    (request.PageID - 1) * request.PageSize,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}
	context.JSON(http.StatusOK, entries)
}

type GetBankEntriesHandlerRequest struct {
	PageSize int32 `form:"page_size" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}

func (s *Server) GetBankEntriesHandler(context *gin.Context) {
	var request GetAccountEntriesHandlerRequest
	err := context.ShouldBindQuery(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	entries, err := s.store.ListBankEntries(context, db.ListBankEntriesParams{
		Limit:  request.PageSize,
		Offset: (request.PageID - 1) * request.PageSize,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}

	context.JSON(http.StatusOK, entries)
}
