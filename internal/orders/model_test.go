package orders_test

import (
	"testing"

	"github.com/fernandoocampo/fancy-worker/internal/orders"
	"github.com/fernandoocampo/fancy-worker/internal/repositories"
	"github.com/stretchr/testify/assert"
)

func TestTransformOrderToRecord(t *testing.T) {
	// GIVEN
	order := orders.Order{
		ID:     "123",
		Amount: 13.45,
	}
	expectedRecord := repositories.Record{
		ID:     "123",
		Amount: 13.45,
	}

	// WHEN
	got := order.ToRecord()

	// THEN
	assert.Equal(t, expectedRecord, got)
}
