package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/start", report)

	http.ListenAndServe(":8080", router)
}

type register struct {
	cpf    string //(Queria colocar o CPF como ID, provalmente só consiga fazer isso com banco. Ainda não encontrei)
	name   string
	age    int64
	state  string
	report string
}

var registers = []register{

	{cpf: "8444096358", name: "Ted", age:  35, state: "Rio de Janeiro", report: "violência doméstica"},

}

func report(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(registers)

}
