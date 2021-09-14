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

//var registers = []register{

//	{cpf: "8444096358", name: "Ted", age:  35, state: "Rio de Janeiro", report: "violência doméstica"},
//	{cpf: "9547078545", name: "Robin", age:  35, state: "São Paulo", report: "roubo"},
//Poderia usar Map aqui também? (Estudar e analisar LISTA E MAP para procedimento
//}


var register1 = register{

	("15092785958"),("Joana"),28, "Paraná", "Assalto",
}

var register2 = register{

	("894556215415"),("Fred"),18, "Minas Gerais", "Maus tratos",
}

var registers[] = register(register1,register2)

func report(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(registers)

}
