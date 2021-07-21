package workers

import (
	"log"

	"github.com/fernandoocampo/fancy-worker/internal/orders"
	"github.com/fernandoocampo/fancy-worker/internal/repositories"
)

// Worker defines a worker implementation.
type Worker struct {
	orderStream  chan orders.Order
	resultStream chan repositories.Record
}

// New creates a new worker.
func New(orderStream chan orders.Order) *Worker {
	return &Worker{
		orderStream:  orderStream,
		resultStream: make(chan repositories.Record),
	}
}

// Start starts the worker.
func (w *Worker) Start(done <-chan interface{}) error {
	defer close(w.resultStream)
	for {
		select {
		case <-done:
			log.Println("level", "INFO", "object", "workers.worker", "msg", "terminating operations by application request")
			return nil
		case order, ok := <-w.orderStream:
			if !ok {
				log.Println("level", "INFO", "object", "workers.worker", "msg", "terminating operations because order stream was closed")
				return nil
			}
			log.Println("level", "INFO", "object", "workers.worker", "msg", "new order to proces", "order", order)
			w.resultStream <- order.ToRecord()
		}
	}
}

// ResultStream return the worker result stream.
func (w *Worker) ResultStream() chan repositories.Record {
	if w == nil {
		return nil
	}
	return w.resultStream
}
