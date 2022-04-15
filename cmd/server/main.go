package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"

	"boilerplate/internal/database"
	barHandler "boilerplate/internal/handlers/bar"
	fooHandler "boilerplate/internal/handlers/foo"
	barRepo "boilerplate/internal/repositiories/bar"
	barService "boilerplate/internal/services/bar"
	"boilerplate/internal/services/foo"
)

func main() {
	println("boilerplate")
	newRelicApp, err := newrelic.NewApplication()
	if err != nil {
		panic("newrelic create panic " + err.Error())
	}

	dbPool := database.NewDBPool()
	client := http.DefaultClient

	fooService := foo.New(client)
	fHandler := fooHandler.New(fooService)

	bRepo := barRepo.New(newRelicApp, dbPool)
	bService := barService.New(bRepo)
	bHandler := barHandler.New(bService)

	router := mux.NewRouter()
	router.Use(nrgorilla.Middleware(newRelicApp))

	router.HandleFunc("/foo", fHandler.Handle)
	router.HandleFunc("/bar", bHandler.Handle)

	http.ListenAndServe("", router)
}
