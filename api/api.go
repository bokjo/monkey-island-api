package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/bokjo/monkey-island-api/handlers"
)

// API struct
type API struct {
	Router *mux.Router
	Db     *sql.DB
}

// Init point for the API - DB
func (api *API) Init(username, password, dbname string) {

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

	var err error

	api.Db, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	api.Router = mux.NewRouter()
	api.initRoutes()
}

// Run the API
func (api *API) Run() {
	log.Fatal(http.ListenAndServe(":1234", api.Router))
}

func (api *API) initRoutes() {

	//Cuddly Toys
	api.Router.HandleFunc("/api/cuddly_toys", handlers.GetCuddlyToys).Methods(http.MethodGet)

	//Dogs
	api.Router.HandleFunc("/api/cuddly_toys/dogs", handlers.GetAllDogs).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/dogs", handlers.CreateDog).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/cuddly_toys/dogs/{id:[0-9]+}", handlers.GetDog).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/dogs/{id:[0-9]+}", handlers.UpdateDog).Methods(http.MethodPut)
	api.Router.HandleFunc("/api/cuddly_toys/dogs/{id:[0-9]+}", handlers.DeleteDog).Methods(http.MethodDelete)

	//Monkeys
	api.Router.HandleFunc("/api/cuddly_toys/monkeys", handlers.GetAllMonkeys).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys", handlers.CreateMonkey).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys/{id:[0-9]+}", handlers.GetMonkey).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys/{id:[0-9]+}", handlers.UpdateMonkey).Methods(http.MethodPut)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys/{id:[0-9]+}", handlers.DeleteMonkey).Methods(http.MethodDelete)

	//Weapons
	api.Router.HandleFunc("/api/weapons", handlers.GetAllWeapons).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/weapons", handlers.CreateWeapon).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/weapons/{id:[0-9]+}", handlers.GetWeapon).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/weapons/{id:[0-9]+}", handlers.UpdateWeapon).Methods(http.MethodPut)
	api.Router.HandleFunc("/api/weapons/{id:[0-9]+}", handlers.DeleteWeapon).Methods(http.MethodDelete)

	// Ghosts
	api.Router.HandleFunc("/ghosts", handlers.GetGhosts).Methods(http.MethodGet)

}
