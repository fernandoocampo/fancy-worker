package orders

import (
	"log"

	"github.com/fernandoocampo/fancy-worker/internal/repositories"
)

// Order order data.
type Order struct {
	ID     string
	Amount float64
}

// ToRecord transform the order to a repository record
func (o Order) ToRecord() repositories.Record {
	result := repositories.Record{
		ID:     o.ID,
		Amount: o.Amount,
	}
	log.Println("level", "INFO", "object", "orders.Order", "msg", "transforming order to record", "order", o, "record", result)
	return result
}
