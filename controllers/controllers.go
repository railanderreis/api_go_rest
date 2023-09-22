package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/railanderreis/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonaliades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}
