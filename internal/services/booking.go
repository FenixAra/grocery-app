package services

import (
	"errors"
	"time"

	"github.com/FenixAra/grocery-app/dtos"
	"github.com/FenixAra/grocery-app/internal/daos"
	"github.com/FenixAra/grocery-app/internal/daos/db"
	"github.com/FenixAra/grocery-app/internal/models"
	"github.com/FenixAra/grocery-app/utils/log"
	"github.com/pborman/uuid"
)

var (
	ErrNotEnoughInventory   = errors.New("Not enough inventory for item")
	ErrNoRegistersAvailable = errors.New("No Registers available")
)

type Booking struct {
	l         *log.Logger
	dbConn    *db.DBConn
	booking   *daos.Booking
	inventory *daos.Inventory
	item      *daos.Item
	lineItem  *daos.LineItem
	account   *daos.Account
	discount  *daos.Discount
	register  *daos.Register
}

func NewBooking(l *log.Logger, dbConn *db.DBConn) *Booking {
	return &Booking{
		l:         l,
		dbConn:    dbConn,
		booking:   daos.NewBooking(l, dbConn),
		inventory: daos.NewInventory(l, dbConn),
		item:      daos.NewItem(l, dbConn),
		lineItem:  daos.NewLineItem(l, dbConn),
		account:   daos.NewAccount(l, dbConn),
		discount:  daos.NewDiscount(l, dbConn),
		register:  daos.NewRegister(l, dbConn),
	}
}

func (b *Booking) blockInventories(itemID string, count int) ([]string, error) {
	barCodes, err := b.inventory.GetAvailable(itemID)
	if err != nil {
		return nil, err
	}

	if len(barCodes) < count {
		return nil, ErrNotEnoughInventory
	}

	err = b.inventory.InventoryStatus(barCodes[0:count], models.InventoryBlocked, models.InventoryAvailable)
	if err != nil {
		return nil, err
	}

	return barCodes[0:count], nil
}

func (b *Booking) LineItemModelArrayToDtoArray(lineItems []models.LineItem, count int) []dtos.BillLineItem {
	var billLineItems []dtos.BillLineItem
	for i := 0; i < count; i++ {
		for _, lineItem := range lineItems {
			billLineItems = append(billLineItems, dtos.BillLineItem{
				ID:          lineItem.ID,
				PriceCardID: lineItem.PriceCardID,
				Name:        lineItem.Name,
				Code:        lineItem.Code,
				Description: lineItem.Description,
				Amount:      lineItem.Amount,
			})
		}
	}
	return billLineItems
}

func (b *Booking) DiscountModelArrayToDtoArray(discounts []models.Discount, count int) []dtos.BillDiscount {
	var billDiscounts []dtos.BillDiscount
	for i := 0; i < count; i++ {
		for _, discount := range discounts {
			billDiscounts = append(billDiscounts, dtos.BillDiscount{
				ID:          discount.ID,
				Name:        discount.Name,
				Code:        discount.Code,
				Description: discount.Description,
				Amount:      discount.Amount,
			})
		}
	}
	return billDiscounts
}

func (b *Booking) ConfirmBooking(req *dtos.BookingData) (*dtos.GenerateBillResponse, error) {
	res, err := b.GenerateBill(req)
	if err != nil {
		b.l.Error("Err:", err)
		return nil, err
	}

	err = b.dbConn.ExecuteInTransaction(func() error {
		register, err := b.register.Get(res.RegisterID)
		if err != nil {
			b.l.Error("Err:", err)
			return err
		}

		err = b.register.SetStatus(res.RegisterID, models.RegisterAvailable, models.RegisterOccupied)
		if err != nil {
			b.l.Error("Err:", err)
			return err
		}

		err = b.inventory.InventoryStatus(res.Inventories, models.InventorySold, models.InventoryBlocked)
		if err != nil {
			b.l.Error("Err:", err)
			return err
		}

		return b.booking.Upsert(&models.Booking{
			ID:          uuid.New(),
			AccountID:   res.AccountID,
			RegisterID:  res.RegisterID,
			EmployeeID:  register.AccountID.String,
			Amount:      res.Amount,
			Bill:        res.Bill,
			Inventories: res.Inventories,
			CreatedAt:   time.Now().UTC(),
		})
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *Booking) GenerateBill(req *dtos.BookingData) (*dtos.GenerateBillResponse, error) {
	res := &dtos.GenerateBillResponse{}
	var amount int
	err := b.dbConn.ExecuteInTransaction(func() error {
		account, err := b.account.Get(req.AccountID)
		if err != nil {
			b.l.Error("Err:", err)
			return err
		}

		registers, err := b.register.GetAvailable()
		if err != nil {
			b.l.Error("Err:", err)
			return err
		}

		if len(registers) == 0 {
			return ErrNoRegistersAvailable
		}

		err = b.register.SetStatus(registers[0], models.RegisterOccupied, models.RegisterAvailable)
		if err != nil {
			b.l.Error("Err:", err)
			return err
		}

		res.RegisterID = registers[0]
		res.AccountID = account.ID
		for itemID, count := range req.Items {
			item, err := b.item.Get(itemID)
			if err != nil {
				b.l.Error("Err:", err)
				return err
			}

			inventories, err := b.blockInventories(itemID, count)
			if err != nil {
				b.l.Error("Err:", err)
				return err
			}

			res.Inventories = append(res.Inventories, inventories...)
			lineItems, err := b.lineItem.GetForPriceCardID(item.PriceCardID)
			if err != nil {
				b.l.Error("Err:", err)
				return err
			}

			res.Bill.LineItems = append(res.Bill.LineItems, b.LineItemModelArrayToDtoArray(lineItems, count)...)
			discounts, err := b.discount.GetForTags(item.Tags)
			if err != nil {
				b.l.Error("Err:", err)
				return err
			}

			res.Bill.Discounts = append(res.Bill.Discounts, b.DiscountModelArrayToDtoArray(discounts, count)...)
		}

		for _, lineItem := range res.Bill.LineItems {
			amount += lineItem.Amount
		}

		for _, discount := range res.Bill.Discounts {
			amount -= discount.Amount
		}

		res.Amount = amount
		res.Bill.Amount = amount

		return nil
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
