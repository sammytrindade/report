package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/start", report)

	http.ListenAndServe(":8080", router)
}

type register struct {
	id uuid.UUID
	name string
	age int64
	state string
	report string
	
}

func report(w http.ResponseWriter, r *http.Request){
	
}