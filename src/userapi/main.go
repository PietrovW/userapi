package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PietrovW/useapi/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Inicjalizacja przykładowych użytkowników
	handlers.InitUsers()

	// Tworzenie nowego routera
	router := mux.NewRouter()

	// Mapowanie ścieżek
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	// Uruchomienie serwera
	fmt.Println("Serwer działa na porcie 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
