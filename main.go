package main

import (
	"net/http"
	"os"

	"github.com/FenixAra/grocery-app/internal/config"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/handlers"
	"github.com/FenixAra/grocery-app/utils/log"
)

func main() {
	db.Init(config.DATABASE_URL)
	ok := db.RunMigration(config.DATABASE_URL)
	if !ok {
		os.Exit(1)
		return
	}

	l := log.NewLogger("")

	l.Info("Port: ", config.PORT)
	http.ListenAndServe(":"+config.PORT, handlers.GetRouter())
}
