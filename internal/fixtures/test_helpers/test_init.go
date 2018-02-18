package test_helpers

import (
	"math/rand"
	"sync"
	"time"

	"github.com/FenixAra/grocery-app/internal/config"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/utils/log"
)

var once sync.Once

func TestInit() (*log.Logger, *db.DBConn) {
	db.Init(config.DATABASE_URL_TEST)
	l := log.NewLogger("")
	dbConn := new(db.DBConn)
	dbConn.Init(l)
	once.Do(func() {
		db.RunMigration(config.DATABASE_URL_TEST)
		rand.Seed(time.Now().UTC().UnixNano())
	})

	return l, dbConn
}
