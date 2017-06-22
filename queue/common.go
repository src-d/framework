package queue

import (
	"errors"
	"io"
	"net/url"
	"time"
)

// Priority represents a priority level.
type Priority uint8

const (
	// PriorityUrgent represents an urgent priority level.
	PriorityUrgent Priority = 8
	// PriorityNormal represents a normal priority level.
	PriorityNormal Priority = 4
	// PriorityLow represents a low priority level.
	PriorityLow Priority = 0
)

var (
	// ErrAlreadyClosed is the error returned when trying to close an already closed
	// connection.
	ErrAlreadyClosed = errors.New("already closed")
	// ErrEmptyJob is the error returned when an empty job is published.
	ErrEmptyJob = errors.New("invalid empty job")
	// ErrTxNotSupported is the error returned when the transaction receives a
	// callback does not know how to handle.
	ErrTxNotSupported = errors.New("transactions not supported")
	// ErrUnsupportedProtocol is the error returned when a Broker does not know how
	// to connect to a given URL
	ErrUnsupportedProtocol = errors.New("unsupported protocol")
)

const (
	protoAMQP   string = "amqp"
	protoMemory        = "memory"
)

// Broker represents a message broker.
type Broker interface {
	// Queue returns a Queue from the with the given name.
	Queue(string) (Queue, error)
	// Close closes the connection of the Broker.
	Close() error
}

// NewBroker creates a new Broker based on the given URI. Possible URIs are
//   amqp://<host>[:port]
//   memory://
func NewBroker(uri string) (Broker, error) {
	url, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	switch url.Scheme {
	case protoAMQP:
		return NewAMQPBroker(uri)
	case protoMemory:
		return NewMemoryBroker(), nil
	default:
		return nil, ErrUnsupportedProtocol
	}
}

// TxCallback is a function to be called in a transaction.
type TxCallback func(q Queue) error

// Queue represents a message queue.
type Queue interface {
	// Publish publishes the given Job to the queue.
	Publish(*Job) error
	// Publish publishes the given Job to the queue with a given delay.
	PublishDelayed(*Job, time.Duration) error
	// Transaction executes the passed TxCallback inside a transaction.
	Transaction(TxCallback) error
	// Consume returns a JobIter for the queue.
	Consume() (JobIter, error)
	// RepublishBuried republish all jobs in the buried queue to the main one
	RepublishBuried() error
}

// JobIter represents an iterator over a set of Jobs.
type JobIter interface {
	// Next returns the next Job in the iterator. It should block until the
	// job becomes available. Returns ErrAlreadyClosed if the iterator is
	// closed.
	Next() (*Job, error)
	io.Closer
}
