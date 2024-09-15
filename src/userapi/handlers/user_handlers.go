package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PietrovW/useapi/models"

	"github.com/gorilla/mux"
)

var users []models.User

var nextID int = 1

// Funkcja inicjalizująca przykładowych użytkowników (opcjonalna)
func InitUsers() {
	users = []models.User{
		{ID: NextID(), Name: "Jan Kowalski", Email: "jan.kowalski@example.com"},
		{ID: NextID(), Name: "Anna Nowak", Email: "anna.nowak@example.com"},
	}
}

func NextID() int {
	id := nextID
	nextID++
	return id
}

// Pobieranie wszystkich użytkowników
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Pobieranie użytkownika po ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Nieprawidłowe ID", http.StatusBadRequest)
		return
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "Użytkownik nie znaleziony", http.StatusNotFound)
}

// Dodawanie nowego użytkownika
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Nieprawidłowe dane", http.StatusBadRequest)
		return
	}
	user.ID = NextID()
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// Aktualizowanie istniejącego użytkownika
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Nieprawidłowe ID", http.StatusBadRequest)
		return
	}

	var updatedUser models.User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, "Nieprawidłowe dane", http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if user.ID == id {
			updatedUser.ID = user.ID
			users[index] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.Error(w, "Użytkownik nie znaleziony", http.StatusNotFound)
}

// Usuwanie użytkownika
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Nieprawidłowe ID", http.StatusBadRequest)
		return
	}

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			fmt.Fprintf(w, "Użytkownik o ID %d został usunięty", id)
			return
		}
	}
	http.Error(w, "Użytkownik nie znaleziony", http.StatusNotFound)
}
