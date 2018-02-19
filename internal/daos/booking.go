package daos

import (
	"fmt"
	"time"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Booking struct {
	l  *log.Logger
	db *db.DBConn
}

func NewBooking(l *log.Logger, db *db.DBConn) *Booking {
	return &Booking{
		l:  l,
		db: db,
	}
}

func (v *Booking) Persist(Booking *models.Booking) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Booking VALUES (%s, %s, %s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Booking.ID), qa.Append(Booking.AccountID), qa.Append(Booking.RegisterID), qa.Append(Booking.EmployeeID),
		qa.Append(Booking.Amount), qa.Append(Booking.Bill), qa.Append(Booking.Inventories), qa.Append(time.Now().UTC()))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Booking) Upsert(Booking *models.Booking) error {
	err := v.Persist(Booking)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Booking SET user_id = %s, register_id = %s, employee_id = %s, amount = %s,
		bill = %s, inventories = %s WHERE id = %s`, qa.Append(Booking.AccountID), qa.Append(Booking.RegisterID),
			qa.Append(Booking.EmployeeID), qa.Append(Booking.Amount), qa.Append(Booking.Bill),
			qa.Append(Booking.Inventories), qa.Append(Booking.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Booking) Get(id string) (*models.Booking, error) {
	Booking := &models.Booking{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Booking WHERE id = $1`, id).Scan(
		&Booking.ID,
		&Booking.AccountID,
		&Booking.RegisterID,
		&Booking.EmployeeID,
		&Booking.Amount,
		&Booking.Bill,
		&Booking.Inventories,
		&Booking.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return Booking, nil
}

func (v *Booking) GetAll() ([]models.Booking, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Booking`)
	if err != nil {
		return nil, err
	}

	var Bookings []models.Booking
	for rows.Next() {
		var Booking models.Booking
		err = rows.Scan(
			&Booking.ID,
			&Booking.AccountID,
			&Booking.RegisterID,
			&Booking.EmployeeID,
			&Booking.Amount,
			&Booking.Bill,
			&Booking.Inventories,
			&Booking.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		Bookings = append(Bookings, Booking)
	}

	return Bookings, nil
}
