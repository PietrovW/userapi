package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PietrovW/useapi/models"
	"github.com/go-playground/validator"
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

// GetUsers godoc
// @Summary      Pobierz listę użytkowników
// @Description  Zwraca listę wszystkich użytkowników z opcjonalną paginacją
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page   query      int     false  "Numer strony"
// @Param        limit  query      int     false  "Liczba elementów na stronie"
// @Success      200    {object}   map[string]interface{}
// @Failure      400    {string}   string  "Nieprawidłowe parametry"
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser godoc
// @Summary      Pobierz użytkownika
// @Description  Zwraca użytkownika o podanym ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int     true  "ID użytkownika"
// @Success      200  {object}  models.User
// @Failure      400  {string}  string  "Nieprawidłowe ID"
// @Failure      404  {string}  string  "Użytkownik nie znaleziony"
// @Router       /users/{id} [get]
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

// CreateUser godoc
// @Summary      Dodaj nowego użytkownika
// @Description  Tworzy nowego użytkownika na podstawie danych wejściowych
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Nowy użytkownik"
// @Success      200   {object}  models.User
// @Failure      400   {string}  string  "Nieprawidłowe dane"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(user)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
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
	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(updatedUser)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
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
