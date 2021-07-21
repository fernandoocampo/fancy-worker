package workers_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/fernandoocampo/fancy-worker/internal/orders"
	"github.com/fernandoocampo/fancy-worker/internal/repositories"
	"github.com/fernandoocampo/fancy-worker/internal/workers"
)

func TestProcessData(t *testing.T) {
	// GIVEN
	request := orders.Order{
		ID:     "ABC123",
		Amount: 65.5,
	}
	expectedRecord := repositories.Record{
		ID:     "ABC123",
		Amount: 65.5,
	}
	orderStream := make(chan orders.Order)
	doneStream := make(chan interface{})
	worker := workers.New(orderStream)
	resultData := make([]repositories.Record, 0)
	wg := sync.WaitGroup{}

	// WHEN
	wg.Add(3)
	go func(done chan interface{}) {
		defer wg.Done()
		err := worker.Start(done)
		if err != nil {
			doneStream <- true
			t.Errorf("unexpected error: %s", err)
		}
	}(doneStream)

	go func(done chan interface{}, orders chan orders.Order, request orders.Order) {
		defer wg.Done()
		defer close(orders)
		select {
		case <-done:
			return
		case orders <- request:
		}
	}(doneStream, orderStream, request)

	go func(done chan interface{}, result chan repositories.Record) {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case record, ok := <-result:
				if !ok {
					return
				}
				resultData = append(resultData, record)
			}
		}
	}(doneStream, worker.ResultStream())
	wg.Wait()
	close(doneStream)
	// THEN
	assert.Equal(t, 1, len(resultData))
	assert.Equal(t, expectedRecord, resultData[0])
}
