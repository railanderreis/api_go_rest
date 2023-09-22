package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/railanderreis/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonaliades)
	log.Fatal(http.ListenAndServe(":8000", r))
}
