package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "simple_bank/db/sqlc"
)

type CreateAccountHandlerRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR RUB NGN"`
}

// CreateAccountHandler / @Create an account
// @Tags Account
// @Accept  json
// @Produce  json
// @Param CreateAccount body CreateAccountHandlerRequest true "Create_Account"
// @Success 200 {object} db.Account "ok"
// @Failure 400 {object} ServerError  "Something went wrong"
// @Failure 500 {object} ServerError  "Something went wrong"
// @Router /api/v1/account/createaccount [POST]
func (s *Server) CreateAccountHandler(context *gin.Context) {
	var request CreateAccountHandlerRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}

	account, err := s.store.CreateAccount(context, db.CreateAccountParams{
		Owner:    request.Owner,
		Balance:  0,
		Currency: request.Currency,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}

	context.JSON(http.StatusOK, account)
	return
}

type GetAccountByIDHandlerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// GetAccountByIDHandler / @Get an account
// @Tags Account
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} db.Account "ok"
// @Failure 400 {object} ServerError  "Something went wrong"
// @Failure 500 {object} ServerError  "Something went wrong"
// @Router /api/v1/account/getaccount/{id} [GET]
func (s *Server) GetAccountByIDHandler(context *gin.Context) {
	var request GetAccountByIDHandlerRequest
	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}

	account, err := s.store.GetAccount(context, request.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}

	context.JSON(http.StatusOK, account)

}

type GetAccountsHandlerRequest struct {
	PageID   int32 `form:"page_id" binding:"min=1"`
	PageSize int32 `form:"page_size" binding:"min=1,max=10"`
}

// GetAccountsHandler / @Get Accounts
// @Tags Account
// @Produce  json
// @Param page_id query int32 true "pageID"
// @Param page_size query int32 true "pageSize"
// @Success 200 {object} []db.Account "ok"
// @Failure 400 {object} ServerError  "We require all fields"
// @Router /api/v1/account/getaccounts [GET]
func (s *Server) GetAccountsHandler(context *gin.Context) {
	var request GetAccountsHandlerRequest
	err := context.ShouldBindQuery(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	params := db.ListAccountsParams{
		Limit:  request.PageSize,
		Offset: (request.PageID - 1) * request.PageSize,
	}

	accounts, err := s.store.ListAccounts(context, params)
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}

	context.JSON(http.StatusOK, accounts)
	return
}

type UpdateAccountHandlerRequest struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

// UpdateAccountHandler / @Update Account
// @Tags Account
// @Accept json
// @Produce  json
// @Param requestBody body UpdateAccountHandlerRequest  true "body"
// @Success 200 {object} db.Account "ok"
// @Failure 400 {object} ServerError  "We require all fields"
// @Router /api/v1/account/updateaccount [PUT]
func (s *Server) UpdateAccountHandler(context *gin.Context) {
	var request UpdateAccountHandlerRequest
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}
	account, err := s.store.UpdateAccount(context, db.UpdateAccountParams{
		ID:      request.ID,
		Balance: request.Balance,
	})
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}

	context.JSON(http.StatusOK, account)
	return
}

type DeleteAccountHandlerRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// DeleteAccountHandler / @Delete Account
// @Tags Account
// @Produce  json
// @Param id path int64  true "account id"
// @Success 200 {object} nil  "ok"
// @Failure 400 {object} ServerError  "We require all fields"
// @Router /api/v1/account/deleteaccount/{id} [DELETE]
func (s *Server) DeleteAccountHandler(context *gin.Context) {
	var request DeleteAccountHandlerRequest
	err := context.ShouldBindUri(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, NewServerError(err))
		return
	}

	err = s.store.DeleteAccount(context, request.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, NewServerError(err))
		return
	}

	context.JSON(http.StatusOK, struct{}{})
}
