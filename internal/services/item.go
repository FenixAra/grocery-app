package services

import (
	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/arrays"
	"github.com/FenixAra/grocery-app/utils/log"
	pgx "gopkg.in/jackc/pgx.v2"
)

type Item struct {
	l         *log.Logger
	dbConn    *db.DBConn
	item      *daos.Item
	priceCard *daos.PriceCard
	lineItem  *daos.LineItem
	category  *daos.Category
	tag       *daos.Tag
}

func NewItem(l *log.Logger, dbConn *db.DBConn) *Item {
	return &Item{
		l:         l,
		dbConn:    dbConn,
		item:      daos.NewItem(l, dbConn),
		priceCard: daos.NewPriceCard(l, dbConn),
		lineItem:  daos.NewLineItem(l, dbConn),
		tag:       daos.NewTag(l, dbConn),
		category:  daos.NewCategory(l, dbConn),
	}
}

func (i *Item) GetTags(categories []string) ([]string, error) {
	var tags []string
	tags = categories
	for _, category := range categories {
		res, err := i.category.GetParent(category)
		if err != nil && err != pgx.ErrNoRows {
			return nil, err
		}

		if len(res) > 0 {
			tags = arrays.AppendWithoutDuplicates(tags, res)
			res, err = i.GetTags(res)
			if err != nil {
				return nil, err
			}

			if len(res) > 0 {
				tags = arrays.AppendWithoutDuplicates(tags, res)
			}
		}
	}
	return tags, nil
}

func (i *Item) Save(item *dtos.Item) error {
	return i.dbConn.ExecuteInTransaction(func() error {
		err := i.priceCard.Upsert(models.NewPriceCard(&item.PriceCard))
		if err != nil {
			return err
		}

		for _, lineItem := range item.PriceCard.LineItems {
			err = i.lineItem.Upsert(models.NewLineItem(&lineItem))
			if err != nil {
				return err
			}
		}

		item.PriceCardID = item.PriceCard.ID
		item.Tags, err = i.GetTags(item.Categories)
		if err != nil {
			return err
		}

		item.Tags = append(item.Tags, item.ID)
		err = i.item.Upsert(models.NewItem(item))
		if err != nil {
			return err
		}

		return i.tag.Upsert(&models.Tag{
			ID:   item.ID,
			Name: item.Name,
			Type: models.TypeItem,
		})
	})
}
