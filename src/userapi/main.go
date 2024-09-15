package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/PietrovW/useapi/handlers"

	"github.com/gorilla/mux"
	_ "github.com/swaggo/http-swagger/example/gorilla/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title           Przykładowe API użytkowników
// @version         1.0
// @description     To jest przykładowe API do zarządzania użytkownikami.
// @termsOfService  http://example.com/terms/

// @contact.name   Dział Wsparcia
// @contact.url    http://example.com/support
// @contact.email  support@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8000
// @BasePath  /v2
func main() {
	// Inicjalizacja przykładowych użytkowników
	handlers.InitUsers()

	// Tworzenie nowego routera
	router := mux.NewRouter()
	uri, err := url.Parse("http://localhost:1323/api/v3")
	if err != nil {
		panic(err)
	}

	// Mapowanie ścieżek
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.BeforeScript(`const UrlMutatorPlugin = (system) => ({
			rootInjects: {
			  setScheme: (scheme) => {
				const jsonSpec = system.getState().toJSON().spec.json;
				const schemes = Array.isArray(scheme) ? scheme : [scheme];
				const newJsonSpec = Object.assign({}, jsonSpec, { schemes });
		  
				return system.specActions.updateJsonSpec(newJsonSpec);
			  },
			  setHost: (host) => {
				const jsonSpec = system.getState().toJSON().spec.json;
				const newJsonSpec = Object.assign({}, jsonSpec, { host });
		  
				return system.specActions.updateJsonSpec(newJsonSpec);
			  },
			  setBasePath: (basePath) => {
				const jsonSpec = system.getState().toJSON().spec.json;
				const newJsonSpec = Object.assign({}, jsonSpec, { basePath });
		  
				return system.specActions.updateJsonSpec(newJsonSpec);
			  }
			}
		  });`),
		httpSwagger.Plugins([]string{"UrlMutatorPlugin"}),
		httpSwagger.UIConfig(map[string]string{
			"onComplete": fmt.Sprintf(`() => {
			  window.ui.setScheme('%s');
			  window.ui.setHost('%s');
			  window.ui.setBasePath('%s');
			}`, uri.Scheme, uri.Host, uri.Path),
		}),
	))
	// Uruchomienie serwera
	fmt.Println("Serwer działa na porcie 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
