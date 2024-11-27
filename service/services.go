package service

import (
	"log"
	"net/http"
	"os"
	"routersmux/handler"
	"routersmux/helpers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func RouterHandler() {
	r := mux.NewRouter()
	r.HandleFunc("/create", handler.CreateUserDetails).Methods("POST")
	r.HandleFunc("/users{id}", handler.GetSingleUser).Methods("GET")
	r.HandleFunc("/users", handler.GetMultipleUser).Methods("GET")
	r.HandleFunc("/update{employee_id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete{id}", handler.Deleteuser).Methods("DELETE")

	err := godotenv.Load(".env")
	if err != nil {
		helpers.ErrorData.Println("Failed to Load the env file ")
	}
	log.Fatal(http.ListenAndServe(os.Getenv("localhost"), r))
}
