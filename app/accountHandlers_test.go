package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/tedoham/coins-api/dto"
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

func Test_GetAccounts(t *testing.T) {

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
