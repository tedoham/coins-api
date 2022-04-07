package app

import (
	"encoding/json"
	"net/http"

	"github.com/tedoham/coins-api/dto"
	"github.com/tedoham/coins-api/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ch *AccountHandlers) getAccounts(w http.ResponseWriter, r *http.Request) {

	accounts, err := ch.service.ListAllAccount()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, accounts)
	}
}

func (ch *AccountHandlers) getPayments(w http.ResponseWriter, r *http.Request) {

	payments, err := ch.service.ListAllPayments()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, payments)
	}
}

func (h *AccountHandlers) makeTransfer(w http.ResponseWriter, r *http.Request) {

	// decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		// make transaction
		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
