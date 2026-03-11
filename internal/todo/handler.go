package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi! Welcome to my server!")
}

func JsonTest(w http.ResponseWriter, r *http.Request) {
	res := JsonResponse{Title: "API Response", Message: "This is an API response"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
