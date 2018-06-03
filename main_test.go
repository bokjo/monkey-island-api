package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/bokjo/monkey-island-api/api"
)

var a api.API

func TestMain(m *testing.M) {

	a = api.API{}
	a.Init("postgres", "postgres", "postgres")

	ensureTableExists("dogs")
	ensureTableExists("monkeys")
	ensureTableExists("weapons")

	code := m.Run()

	clearTable("dogs")
	clearTable("monkeys")
	clearTable("weapons")

	os.Exit(code)

}

func ensureTableExists(tableName string) {
	if _, err := a.Db.Exec(tableCreationQuery(tableName)); err != nil {
		log.Fatal(err)
	}
}

func tableCreationQuery(tableName string) string {

	replace := "energy_level"

	if tableName == "weapons" {
		replace = "power_level"
	}
	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s
(
id SERIAL,
name TEXT NOT NULL,
%s TEXT NOT NULL,
CONSTRAINT %s_pkey PRIMARY KEY (id)
)`, tableName, replace, tableName)
}

func clearTable(tableName string) {
	a.Db.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	a.Db.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH 1", tableName))
}
