package services

import (
	"errors"

	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/utils/log"
)

type Ping struct {
	dbConn *db.DBConn
	l      *log.Logger
	ping   *daos.Ping
}

var (
	ErrUnableToPingDB = errors.New("Unable to ping database")
)

func NewPing(l *log.Logger, dbConn *db.DBConn) *Ping {
	return &Ping{
		l:      l,
		dbConn: dbConn,
		ping:   daos.NewPing(l, dbConn),
	}
}

func (p *Ping) Ping() error {
	ok, err := p.ping.Ping()
	if err != nil {
		return err
	}

	if !ok {
		return ErrUnableToPingDB
	}
	return nil
}
