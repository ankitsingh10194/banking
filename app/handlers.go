package app

// import (
// 	"encoding/json"
// 	"encoding/xml"
// 	"log"
// 	"net/http"

// 	"github.com/ankit/banking/service"
// 	"github.com/gorilla/mux"
// )

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipcode"`
// }

// type CustomerHandler struct {
// 	service service.CustomerService
// }

// func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, req *http.Request) {

// 	customers, err := ch.service.GetAllCustomers()
// 	if err != nil {
// 		writeResponse(w, err.Code, err.AsMessage())
// 	}
// 	if req.Header.Get("Content-Type") == "application/xml" {
// 		w.Header().Set("Content-Type", "application/xml")
// 		xml.NewEncoder(w).Encode(customers)
// 	} else {
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(customers)
// 	}

// }
// func (ch *CustomerHandler) getCustomerById(w http.ResponseWriter, req *http.Request) {
// 	log.Println("getCustomerById() is running...")
// 	customerId := mux.Vars(req)["customer_Id"]
// 	customer, err := ch.service.GetCustomer(customerId)
// 	if err != nil {
// 		writeResponse(w, err.Code, err.AsMessage())
// 	} else {
// 		writeResponse(w, http.StatusOK, customer)
// 	}
// }

// func writeResponse(w http.ResponseWriter, code int, data any) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		panic(err)
// 	}

// }
