package order

import (
	"errors"
	"github.com/Rhymond/go-money"
)

type Item struct {
	ID        string
	Quantity  int
	UnitPrice *money.Money
}

type Order struct {
	ID                string
	CurrencyAlphaCode string
	Items             []Item
}

// ComputeTotal calculates the total price of the order
func (o *Order) ComputeTotal() (*money.Money, error) {
	if len(o.Items) == 0 {
		return nil, errors.New("order has no items")
	}

	total := money.New(0, o.CurrencyAlphaCode)
	for _, item := range o.Items {
		subtotal := item.UnitPrice.Multiply(int64(item.Quantity))
		if err != nil {
			return nil, err
		}
		total, err = total.Add(subtotal)
		if err != nil {
			return nil, err
		}
	}
	return total, nil
}
