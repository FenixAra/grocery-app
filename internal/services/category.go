package services

import (
	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
)

type Category struct {
	l        *log.Logger
	dbConn   *db.DBConn
	category *daos.Category
	tag      *daos.Tag
}

func NewCategory(l *log.Logger, dbConn *db.DBConn) *Category {
	return &Category{
		l:        l,
		dbConn:   dbConn,
		category: daos.NewCategory(l, dbConn),
		tag:      daos.NewTag(l, dbConn),
	}
}

func (c *Category) Save(category *dtos.Category) error {
	err := c.category.Upsert(models.NewCategory(category))
	if err != nil {
		return err
	}

	return c.tag.Upsert(&models.Tag{
		ID:   category.ID,
		Name: category.Name,
		Type: models.TypeCategory,
	})
}
