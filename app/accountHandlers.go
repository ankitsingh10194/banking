package app

import (
	"encoding/json"
	"net/http"

	"github.com/ankitsingh10194/banking/dto"
	"github.com/ankitsingh10194/banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) CreateAccount(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["customer_id"]
	var request dto.AccountRequest
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = id
		resp, appErr := ah.service.NewAccount(request)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusCreated, resp)
		}
	}

}

func (ah AccountHandler) MakeTransaction(w http.ResponseWriter, req *http.Request) {
	customer_id := mux.Vars(req)["customer_id"]
	account_id := mux.Vars(req)["account_id"]
	var r dto.TransactionRequest
	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		r.CustomerId = customer_id
		r.AccountId = account_id
		resp, appErr := ah.service.MakeTransaction(r)
		if appErr != nil {
			writeResponse(w, appErr.Code, appErr.Message)
		} else {
			writeResponse(w, http.StatusOK, resp)
		}
	}
}
