package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Payment struct {
	l  *log.Logger
	db *db.DBConn
}

func NewPayment(l *log.Logger, db *db.DBConn) *Payment {
	return &Payment{
		l:  l,
		db: db,
	}
}

func (v *Payment) Persist(Payment *models.Payment) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Payment VALUES (%s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Payment.ID), qa.Append(Payment.BookingID), qa.Append(Payment.Mode), qa.Append(Payment.PaymentRef),
		qa.Append(Payment.Amount), qa.Append(Payment.Status))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Payment) Upsert(Payment *models.Payment) error {
	err := v.Persist(Payment)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Payment SET booking_id = %s, mode = %s, payment_ref = %s, amount = %s,
		status = %s   
		 WHERE id = %s`, qa.Append(Payment.BookingID), qa.Append(Payment.Mode), qa.Append(Payment.PaymentRef),
			qa.Append(Payment.Amount), qa.Append(Payment.Status), qa.Append(Payment.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Payment) Get(id string) (*models.Payment, error) {
	Payment := &models.Payment{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Payment WHERE id = $1`, id).Scan(
		&Payment.ID,
		&Payment.BookingID,
		&Payment.Mode,
		&Payment.PaymentRef,
		&Payment.Amount,
		&Payment.Status,
	)
	if err != nil {
		return nil, err
	}

	return Payment, nil
}

func (v *Payment) GetAll() ([]models.Payment, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Payment`)
	if err != nil {
		return nil, err
	}

	var Payments []models.Payment
	for rows.Next() {
		var Payment models.Payment
		err = rows.Scan(
			&Payment.ID,
			&Payment.BookingID,
			&Payment.Mode,
			&Payment.PaymentRef,
			&Payment.Amount,
			&Payment.Status,
		)
		if err != nil {
			return nil, err
		}

		Payments = append(Payments, Payment)
	}

	return Payments, nil
}
