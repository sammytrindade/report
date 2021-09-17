package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

var register []Register

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/registro", Reports).Methods("GET")
	router.HandleFunc("/registro/{id}", Report).Methods("GET")
	router.HandleFunc("/registro/{id}", CreateReport).Methods("POST")
	router.HandleFunc("/registro{id}", DeleteReport).Methods("DELETE")
	AppendRegisters()

	http.ListenAndServe(":8080", router)
}

func Reports(w http.ResponseWriter, request *http.Request) {
	json.NewEncoder(w).Encode(register)
}
func Report(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for _, item := range register {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(register)

}
func CreateReport(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	var registred Register
	_ = json.NewDecoder(request.Body).Decode(&Register{})
	registred.ID = params["id"]
	register = append(register, registred)
	json.NewEncoder(w).Encode(register)
}
func DeleteReport(w http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	for index, item := range register {
		if item.ID == params["id"] {
			register = append(register[:index], register[index+1:]...)
			break

		}
		json.NewEncoder(w).Encode(register)
	}

}
func AppendRegisters() {
	register = append(register, Register{"Comércio Ilegal", &Applicant{"1", "Ana", "Sara", 21},
		&Address{"Rio de Janeiro", "Rio de Janeiro"}})

	register = append(register, Register{"Violência Doméstica", &Applicant{"2", "Joana", "Almeida", 45},
		&Address{"Espirito Santo", "Vitória"}})

	register = append(register, Register{"Roubo", &Applicant{"3", "Roberto", "Joaquim", 35},
		&Address{"São Paulo", "Campinas"}})

}

type Register struct {
	Report string
	*Applicant
	*Address
}

type Applicant struct {
	ID        string //`json: "id, omitempty"`
	Firstname string
	Lastname  string
	//CPF       string
	Age int32
}

type Address struct {
	State string
	City  string
}
