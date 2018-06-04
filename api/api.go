package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/bokjo/monkey-island-api/handlers"
	"github.com/bokjo/monkey-island-api/model"
)

// API struct
type API struct {
	Router *mux.Router
	Db     *sql.DB
}

// Init point for the API
func (api *API) Init(username, password, dbname string) {

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

	var err error

	api.Db, err = model.DBConnect(connectionString)

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
	api.Router.Handle("/api/cuddly_toys", handlers.GetCuddlyToys(api.Db)).Methods(http.MethodGet)

	//Dogs
	api.Router.Handle("/api/cuddly_toys/dogs", handlers.GetAllDogs(api.Db)).Methods(http.MethodGet)
	api.Router.Handle("/api/cuddly_toys/dogs", handlers.CreateDog(api.Db)).Methods(http.MethodPost)
	api.Router.Handle("/api/cuddly_toys/dogs/{id:[0-9]+}", handlers.GetDog(api.Db)).Methods(http.MethodGet)
	api.Router.Handle("/api/cuddly_toys/dogs/{id:[0-9]+}", handlers.UpdateDog(api.Db)).Methods(http.MethodPut)
	api.Router.Handle("/api/cuddly_toys/dogs/{id:[0-9]+}", handlers.DeleteDog(api.Db)).Methods(http.MethodDelete)

	//Monkeys
	api.Router.Handle("/api/cuddly_toys/monkeys", handlers.GetAllMonkeys(api.Db)).Methods(http.MethodGet)
	api.Router.Handle("/api/cuddly_toys/monkeys", handlers.CreateMonkey(api.Db)).Methods(http.MethodPost)
	api.Router.Handle("/api/cuddly_toys/monkeys/{id:[0-9]+}", handlers.GetMonkey(api.Db)).Methods(http.MethodGet)
	api.Router.Handle("/api/cuddly_toys/monkeys/{id:[0-9]+}", handlers.UpdateMonkey(api.Db)).Methods(http.MethodPut)
	api.Router.Handle("/api/cuddly_toys/monkeys/{id:[0-9]+}", handlers.DeleteMonkey(api.Db)).Methods(http.MethodDelete)

	//Weapons
	api.Router.Handle("/api/weapons", handlers.GetAllWeapons(api.Db)).Methods(http.MethodGet)
	api.Router.Handle("/api/weapons", handlers.CreateWeapon(api.Db)).Methods(http.MethodPost)
	api.Router.Handle("/api/weapons/{id:[0-9]+}", handlers.GetWeapon(api.Db)).Methods(http.MethodGet)
	api.Router.Handle("/api/weapons/{id:[0-9]+}", handlers.UpdateWeapon(api.Db)).Methods(http.MethodPut)
	api.Router.Handle("/api/weapons/{id:[0-9]+}", handlers.DeleteWeapon(api.Db)).Methods(http.MethodDelete)

	// Ghosts
	api.Router.HandleFunc("/ghosts", handlers.GetGhosts).Methods(http.MethodGet)

}
