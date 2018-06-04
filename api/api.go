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
	cuddlyToyService := model.CuddlyToyService{DB: api.Db}
	cuddlyToyHandler := handlers.CuddlyToyHandler{CuddlyToyService: cuddlyToyService}
	api.Router.HandleFunc("/api/cuddly_toys", cuddlyToyHandler.GetCuddlyToys).Methods(http.MethodGet)

	//Dogs
	dogService := model.DogService{DB: api.Db}
	dogHandler := handlers.DogHandler{DogService: dogService}
	api.Router.HandleFunc("/api/cuddly_toys/dogs", dogHandler.GetAllDogs).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/dogs", dogHandler.CreateDog).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/cuddly_toys/dogs/{id:[0-9]+}", dogHandler.GetDog).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/dogs/{id:[0-9]+}", dogHandler.UpdateDog).Methods(http.MethodPut)
	api.Router.HandleFunc("/api/cuddly_toys/dogs/{id:[0-9]+}", dogHandler.DeleteDog).Methods(http.MethodDelete)

	//Monkeys
	monkeyService := model.MonkeyService{DB: api.Db}
	monkeyHandler := handlers.MonkeyHandler{MonkeyService: monkeyService}
	api.Router.HandleFunc("/api/cuddly_toys/monkeys", monkeyHandler.GetAllMonkeys).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys", monkeyHandler.CreateMonkey).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys/{id:[0-9]+}", monkeyHandler.GetMonkey).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys/{id:[0-9]+}", monkeyHandler.UpdateMonkey).Methods(http.MethodPut)
	api.Router.HandleFunc("/api/cuddly_toys/monkeys/{id:[0-9]+}", monkeyHandler.DeleteMonkey).Methods(http.MethodDelete)

	//Weapons
	weaponService := model.WeaponService{DB: api.Db}
	weaponHandler := handlers.WeaponHandler{WeaponService: weaponService}
	api.Router.HandleFunc("/api/weapons", weaponHandler.GetAllWeapons).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/weapons", weaponHandler.CreateWeapon).Methods(http.MethodPost)
	api.Router.HandleFunc("/api/weapons/{id:[0-9]+}", weaponHandler.GetWeapon).Methods(http.MethodGet)
	api.Router.HandleFunc("/api/weapons/{id:[0-9]+}", weaponHandler.UpdateWeapon).Methods(http.MethodPut)
	api.Router.HandleFunc("/api/weapons/{id:[0-9]+}", weaponHandler.DeleteWeapon).Methods(http.MethodDelete)

	// Ghosts
	ghostService := model.GhostService{}
	ghostHandler := handlers.GhostHandler{GhostService: ghostService}
	api.Router.HandleFunc("/ghosts", ghostHandler.GetGhosts).Methods(http.MethodGet)

}
