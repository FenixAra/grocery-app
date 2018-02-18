package db

import (
	"github.com/FenixAra/grocery-app/internal/config"
	"github.com/FenixAra/grocery-app/utils/log"

	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func RunMigration(url string) bool {
	l := log.NewLogger("")
	m, err := migrate.New(config.MIGRATION_FILE_PATH, url)
	if err != nil {
		l.Fatal("Unable to get migrate instance. Err: ", err)
		return false
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		l.Fatal("Unable to migrate the database. Err: ", err)
		return false
	}

	return true
}
