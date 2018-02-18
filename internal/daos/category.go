package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Category struct {
	l  *log.Logger
	db *db.DBConn
}

func NewCategory(l *log.Logger, db *db.DBConn) *Category {
	return &Category{
		l:  l,
		db: db,
	}
}

func (v *Category) Persist(category *models.Category) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO category VALUES (%s, %s, %s, %s) ON CONFLICT DO NOTHING`, qa.Append(category.ID),
		qa.Append(category.Name), qa.Append(category.Description), qa.Append(category.ParentID))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Category) Upsert(category *models.Category) error {
	err := v.Persist(category)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE category SET name = %s, description = %s, parent_id = %s 
		 WHERE id = %s`, qa.Append(category.Name), qa.Append(category.Description), qa.Append(category.ParentID),
			qa.Append(category.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Category) Get(id string) (*models.Category, error) {
	category := &models.Category{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM category WHERE id = $1`, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.ParentID,
	)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (v *Category) GetAll() ([]models.Category, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM category`)
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.ID, &category.Name, &category.Description, &category.ParentID)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}
