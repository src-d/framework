package queue

type Priority uint8

const (
	PriorityUrgent Priority = 8
	PriorityNormal Priority = 4
	PriorityLow    Priority = 0
)

type Broker interface {
	Queue(string) (Queue, error)
	Close() error
}

type Queue interface {
	Publish(*Job) error
	Consume() (JobIter, error)
}

type JobIter interface {
	Next() (*Job, error)
	Close() error
}
