package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ankitsingh10194/banking/domain"
	"github.com/ankitsingh10194/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityTest() {
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables are not define")
	}
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_ADDR") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatal("DB Environment variables are not define")
	}
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func Start() {
	sanityTest()
	router := mux.NewRouter()
	dbClient := getDbClient()
	//wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	ch1 := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDbSqlx(dbClient))}
	ah := AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}
	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customers/{customer_Id:[0-9]+}", ch1.getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customersByStatus/{status}", ch1.GetAllCustomersByStatus).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)
	//router.HandleFunc("/customer/{customer_Id:[0-9]+}", getCustomerById).Methods(http.MethodGet)

	// start server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}
