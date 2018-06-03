package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// API struct
type API struct {
	Router *mux.Router
	Db     *sql.DB
}

// Init point for the API - DB
func (api *API) Init(username, password, dbname string) {

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", username, password, dbname)

	var err error

	api.Db, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	api.Router = mux.NewRouter()
	api.initRoutes()
}

// Run API
func (api *API) Run() {
	log.Fatal(http.ListenAndServe(":1234", api.Router))
}

func (api *API) initRoutes() {
	api.Router.HandleFunc("/", test).Methods("GET")
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TEST"))
}
