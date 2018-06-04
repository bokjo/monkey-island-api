package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bokjo/monkey-island-api/api"
)

var a api.API

func TestGetCuddlyToys(t *testing.T) {

	a = api.API{}
	a.Init("postgres", "postgres", "postgres")
	defer a.Db.Close()

	clearTable("dogs")
	clearTable("monkeys")

	req, _ := http.NewRequest("GET", "/api/cuddly_toys", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "" {
		t.Errorf("Expected the 'error' key of the response to be set to ''. Got '%s'", m["error"])
	}

}

func clearTable(tableName string) {
	a.Db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	a.Db.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH 1", tableName))
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	a.Router.ServeHTTP(rec, req)

	return rec
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
