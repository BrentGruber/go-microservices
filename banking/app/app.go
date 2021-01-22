package app

import (
	"log"
	"net/http"

	"github.com/BrentGruber/banking/domain"
	"github.com/BrentGruber/banking/service"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
