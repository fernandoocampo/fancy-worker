package workers

import (
	"log"

	"github.com/fernandoocampo/fancy-worker/internal/adapter/anydb"
	"github.com/fernandoocampo/fancy-worker/internal/orders"
)

// Worker defines a worker implementation.
type Worker struct {
	orderStream  chan orders.Order
	resultStream chan anydb.Record
}

// New creates a new worker.
func New(orderStream chan orders.Order) *Worker {
	return &Worker{
		orderStream:  orderStream,
		resultStream: make(chan anydb.Record),
	}
}

// Run make the worker run.
func (w *Worker) Run(done <-chan interface{}) error {
	defer close(w.resultStream)
	for {
		select {
		case <-done:
			log.Println(
				"level", "INFO",
				"object", "workers.worker",
				"method", "Run",
				"msg", "terminating operations by application request",
			)
			return nil
		case order, ok := <-w.orderStream:
			if !ok {
				log.Println("level", "INFO",
					"object", "workers.worker",
					"method", "Run",
					"msg", "terminating operations because order stream was closed",
				)
				return nil
			}
			w.processOrder(order)
		}
	}
}

// processOrder processes the given order
func (w *Worker) processOrder(order orders.Order) {
	log.Println(
		"level", "INFO",
		"object", "workers.worker",
		"method", "processOrder",
		"msg", "new order to proces",
		"order", order,
	)
	w.resultStream <- order.ToRecord()
}

// ResultStream return the worker result stream.
func (w *Worker) ResultStream() chan anydb.Record {
	if w == nil {
		return nil
	}
	return w.resultStream
}
