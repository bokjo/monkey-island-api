package main

import (
	"github.com/bokjo/monkey-island-api/api"
)

func main() {
	mia := api.API{}
	mia.Init("postgres", "postgres", "postgres")

	mia.Run()
}
