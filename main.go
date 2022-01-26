package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thonguyen/rest-api-golang/controllers"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.HomeLink)
	router.HandleFunc("/people", controllers.GetPeople).Methods("GET")
	router.HandleFunc("/people/month", controllers.GetPeopleByMonth).Methods("GET")
	router.HandleFunc("/people/salary", controllers.GetPeopleBySalary).Methods("GET")
	router.HandleFunc("/people/{id}", controllers.GetPeopleById).Methods("GET")
	router.HandleFunc("/people", controllers.CreatePeople).Methods("POST")
	router.HandleFunc("/people/update/{id}", controllers.UpdatePeople).Methods("PUT")
	router.HandleFunc("/people/delete/{id}", controllers.DeletePeople).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
