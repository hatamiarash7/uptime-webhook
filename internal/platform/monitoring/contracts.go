package monitoring

// Monitor is the monitoring interface
type Monitor interface {
	Record(events []Event)
}
