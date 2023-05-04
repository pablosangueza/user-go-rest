package main

import (
	"log"

	"github.com/dom/user/internal/api"
	"github.com/dom/user/internal/config"
	"github.com/dom/user/internal/database"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r, err := api.SetupRouter(db)
	if err != nil {
		panic(err)
	}
	r.AppRouter.Logger.Fatal(r.AppRouter.Start(":8080"))

}
