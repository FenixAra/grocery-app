package daos

import (
	"fmt"

	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Tag struct {
	l  *log.Logger
	db *db.DBConn
}

func NewTag(l *log.Logger, db *db.DBConn) *Tag {
	return &Tag{
		l:  l,
		db: db,
	}
}

func (v *Tag) Persist(Tag *models.Tag) error {
	qa := pgx.QueryArgs{}
	q := fmt.Sprintf(`INSERT INTO Tag VALUES (%s, %s, %s) ON CONFLICT DO NOTHING`,
		qa.Append(Tag.ID), qa.Append(Tag.Name), qa.Append(Tag.Type))
	ct, err := v.db.GetQueryer().Exec(q, qa...)
	if err != nil {
		return err
	}

	if ct.RowsAffected() == 0 {
		return ErrNoRowsInserted
	}
	return nil
}

func (v *Tag) Upsert(Tag *models.Tag) error {
	err := v.Persist(Tag)
	if err != nil {
		qa := pgx.QueryArgs{}
		q := fmt.Sprintf(`UPDATE Tag SET name = %s, type = %s 
		 WHERE id = %s`, qa.Append(Tag.Name), qa.Append(Tag.Type), qa.Append(Tag.ID))
		_, err := v.db.GetQueryer().Exec(q, qa...)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (v *Tag) Get(id string) (*models.Tag, error) {
	Tag := &models.Tag{}
	err := v.db.GetQueryer().QueryRow(`SELECT * FROM Tag WHERE id = $1`, id).Scan(
		&Tag.ID,
		&Tag.Name,
		&Tag.Type,
	)
	if err != nil {
		return nil, err
	}

	return Tag, nil
}

func (v *Tag) GetAll() ([]models.Tag, error) {
	rows, err := v.db.GetQueryer().Query(`SELECT * FROM Tag`)
	if err != nil {
		return nil, err
	}

	var Tags []models.Tag
	for rows.Next() {
		var Tag models.Tag
		err = rows.Scan(
			&Tag.ID,
			&Tag.Name,
			&Tag.Type,
		)
		if err != nil {
			return nil, err
		}

		Tags = append(Tags, Tag)
	}

	return Tags, nil
}
