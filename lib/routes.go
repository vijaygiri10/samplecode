package lib

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//GetRoutes ...
func GetRoutes() http.Handler {
	// Creating Mux Router Object
	router := mux.NewRouter().StrictSlash(true)

	//Registering HTTP End Point with Mux Router
	router.HandleFunc("/", indexHandler).Methods("GET")
	router.HandleFunc("/open/{email_id}/{account_id}/{url_index}/{url_digest}", EmailOpen).Methods("GET")

	allowedHeaders := []string{"X-Requested-With", "Content-Type", "Authorization"}
	allowedMethods := []string{"GET", "HEAD", "OPTIONS"}
	allowedOrigins := []string{"*"}

	return handlers.CORS(handlers.AllowedHeaders(allowedHeaders), handlers.AllowedMethods(allowedMethods), handlers.AllowedOrigins(allowedOrigins))(router)

}
