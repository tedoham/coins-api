package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/tedoham/coins-api/dto"
	"github.com/tedoham/coins-api/errs"
	mock "github.com/tedoham/coins-api/mocks/service"
)

var router *mux.Router
var ch AccountHandlers
var mockService *mock.MockAccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = mock.NewMockAccountService(ctrl)
	ch = AccountHandlers{mockService}

	router = mux.NewRouter()
	router.HandleFunc("/accounts", ch.getAccounts)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_GetAccounts_return_accounts_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyAccounts := []dto.AccountResponse{
		{"1001", "Ashish", "110011"},
		{"1002", "Rob", "110011"},
	}
	mockService.EXPECT().ListAllAccount().Return(dummyAccounts, nil)
	request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_GetAccounts_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().ListAllAccount().Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/accounts", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}

func setupPayment(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = mock.NewMockAccountService(ctrl)
	ch = AccountHandlers{mockService}

	router = mux.NewRouter()
	router.HandleFunc("/payments", ch.getPayments)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_GetPayments_return_payments_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setupPayment(t)
	defer teardown()

	dummyPayments := []dto.TransactionResponse{
		{"1001", "Ashish", "to-rob", "outgoing", "USD", 23423},
		{"1002", "Rob", "to-ashish", "incoming", "USD", 23534},
	}
	mockService.EXPECT().ListAllPayments().Return(dummyPayments, nil)
	request, _ := http.NewRequest(http.MethodGet, "/payments", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}
}

func Test_GetPayments_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setupPayment(t)
	defer teardown()

	mockService.EXPECT().ListAllPayments().Return(nil, errs.NewUnexpectedError("some database error"))
	request, _ := http.NewRequest(http.MethodGet, "/payments", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
