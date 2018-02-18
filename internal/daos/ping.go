package daos

import (
	"errors"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/utils/log"
)

var (
	ErrNoRowsInserted = errors.New("No rows inserted")
)

type Ping struct {
	l      *log.Logger
	dbConn *db.DBConn
}

func NewPing(l *log.Logger, dbConn *db.DBConn) *Ping {
	return &Ping{
		l:      l,
		dbConn: dbConn,
	}
}

func (p *Ping) Ping() (bool, error) {
	var v int32
	err := p.dbConn.GetQueryer().QueryRow("SELECT 1").Scan(&v)
	return (v == 1), err
}
