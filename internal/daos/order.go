package daos

import (
	"fmt"
	"time"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Order struct {
	l  *log.Logger
	db *db.DBConn
}

func NewOrder(l *log.Logger, db *db.DBConn) *Order {
	return &Order{
		l:  l,
		db: db,
	}
}

func (v *Order) Persist(Order *models.Order) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Order VALUES (%s, %s, %s, %s, %s, %s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Order.ID), qa.Append(Order.UserID), qa.Append(Order.RegisterID), qa.Append(Order.EmployeeID),
		qa.Append(Order.Amount), qa.Append(Order.Bill), qa.Append(Order.Inventories), qa.Append(time.Now().UTC()))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Order) Upsert(Order *models.Order) error {
	err := v.Persist(Order)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Order SET user_id = %s, register_id = %s, employee_id = %s, amount = %s,
		bill = %s, inventories = %s WHERE id = %s`, qa.Append(Order.UserID), qa.Append(Order.RegisterID),
			qa.Append(Order.EmployeeID), qa.Append(Order.Amount), qa.Append(Order.Bill),
			qa.Append(Order.Inventories), qa.Append(Order.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Order) Get(id string) (*models.Order, error) {
	Order := &models.Order{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Order WHERE id = $1`, id).Scan(
		&Order.ID,
		&Order.UserID,
		&Order.RegisterID,
		&Order.EmployeeID,
		&Order.Amount,
		&Order.Bill,
		&Order.Inventories,
		&Order.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return Order, nil
}

func (v *Order) GetAll() ([]models.Order, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Order`)
	if err != nil {
		return nil, err
	}

	var Orders []models.Order
	for rows.Next() {
		var Order models.Order
		err = rows.Scan(
			&Order.ID,
			&Order.UserID,
			&Order.RegisterID,
			&Order.EmployeeID,
			&Order.Amount,
			&Order.Bill,
			&Order.Inventories,
			&Order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		Orders = append(Orders, Order)
	}

	return Orders, nil
}
