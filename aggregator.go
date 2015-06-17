package aggregator

// simple test of aggregating system.

type Message struct {
	timestamp int64
	key       string
	value     int
}

// Partitioner accepts the message and publishes them to the corresponding aggregator.
type Partitioner interface {
	Accept(Message)
}

// Aggregator accepts the message.
type Aggregator interface {
	Accept(Message)
}

type Scheduler interface {
	RegisterAggregator(string, Aggregator)
	GetAggregator(int64, string) Aggregator
}

type Environment struct {
}
