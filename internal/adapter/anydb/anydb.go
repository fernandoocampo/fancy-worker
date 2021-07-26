package anydb

// Record contains anydb record data.
type Record struct {
	ID     string
	Amount float64
}

// AnyDB implements any hypothetical database client.
type AnyDB struct{}
