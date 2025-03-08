package order

import (
	"testing"

	"github.com/Rhymond/go-money"
	"github.com/stretchr/testify/assert"
)

func TestComputeTotal(t *testing.T){
	o := Order{
		ID:   "45",
		CurrencyAlphaCode: "INR",
		Items: []Item{
			{
				ID:       "458",
				Quantity: 2,
				UnitPrice: money.New(100, "INR"),
			},
		},
	}
	total, err := o.ComputeTotal()
	assert.NoError(t, err)
	assert.Equal(t, 200,total.Amount())
	assert.Equal(t, "INR",total.Currency())

}