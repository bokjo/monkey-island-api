package main

import (
	"github.com/bokjo/monkey-island-api/api"
)

func main() {
	mia := api.API{}
	// TODO: Implement config, ENV variables some other way go of reatrieving DB cridentials...
	mia.Init("postgres", "postgres", "postgres")

	mia.Run()

}
