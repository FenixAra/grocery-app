package db

import (
	"os"
	"sync"

	"github.com/FenixAra/grocery-app/internal/config"
	"github.com/FenixAra/grocery-app/utils/log"

	"gopkg.in/jackc/pgx.v2"
)

//Pool od database connection
var ConnPool *pgx.ConnPool
var once sync.Once

//Init the connection to DB
func Init(url string) error {
	if ConnPool == nil {
		once.Do(func() {
			connConfig, err := pgx.ParseURI(url)
			if err != nil {
				log.NewLogger("").Fatal("Invalid Database URL, Err: ", err)
				os.Exit(1)
				return
			}
			poolConfig := pgx.ConnPoolConfig{
				ConnConfig:     connConfig,
				MaxConnections: config.MAX_DB_CONNECTIONS,
			}
			ConnPool, err = pgx.NewConnPool(poolConfig)
			if err != nil {
				log.NewLogger("").Fatal("Unable to connect to Database, Err: ", err)
				os.Exit(1)
				return
			}
			log.NewLogger("").Info("Successfully established database connection to ", poolConfig.Database)
		})
	}
	return nil
}

// Interface to abstract the queryer(dbconnection or transaction)
type Queryer interface {
	Exec(sql string, arguments ...interface{}) (pgx.CommandTag, error)
	Query(sql string, args ...interface{}) (*pgx.Rows, error)
	QueryRow(sql string, args ...interface{}) *pgx.Row
}

type DBConn struct {
	conn          *pgx.ConnPool
	tx            *pgx.Tx
	isTransaction bool
	l             *log.Logger
}

func NewDBConn(l *log.Logger) *DBConn {
	return &DBConn{
		conn: ConnPool,
		l:    l,
	}
}

// Initialize the DB connection and assign the existing db connection
func (db *DBConn) Init(l *log.Logger) {
	db.conn = ConnPool
	db.l = l
}

func (db *DBConn) GetQueryer() Queryer {
	if db.isTransaction {
		return db.tx
	} else {
		return db.conn
	}
}

// ExecuteInTransaction executes the given function in DB transaction, i.e. It commits
// only if there is not error otherwise it is rolledback.
func (db *DBConn) ExecuteInTransaction(f func() error) (err error) {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}
	db.tx = tx
	db.isTransaction = true

	defer func() {
		if r := recover(); r != nil {
			db.l.Fatal("Recovered in function ", r)
			db.rollbackTransaction(tx)
		}
		db.isTransaction = false
	}()

	err = f()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	err = tx.Commit()
	if err != nil {
		db.rollbackTransaction(tx)
		return err
	}
	return nil
}

func (db *DBConn) rollbackTransaction(tx *pgx.Tx) {
	err := tx.Rollback()
	if err != nil {
		db.l.Error("Error While rollback, Err: ", err)
	}
}
