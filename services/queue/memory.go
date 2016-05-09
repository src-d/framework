package queue

import (
	"sync"
	"time"
)

type memoryBroker struct {
}

func NewMemoryBroker() Broker {
	return &memoryBroker{}
}

func (b *memoryBroker) Queue(name string) (Queue, error) {
	return &memoryQueue{jobs: make([]*Job, 0, 10)}, nil
}

func (b *memoryBroker) Close() error {
	return nil
}

type memoryQueue struct {
	jobs []*Job
	sync.RWMutex
	idx int
}

func (q *memoryQueue) Publish(job *Job) error {
	q.Lock()
	defer q.Unlock()
	q.jobs = append(q.jobs, job)
	return nil
}

func (q *memoryQueue) PublishDelayed(job *Job, delay time.Duration) error {
	go func() {
		<-time.After(delay)
		q.Publish(job)
	}()
	return nil
}

func (q *memoryQueue) Consume() (JobIter, error) {
	return &memoryJobIter{&q.jobs, &q.idx, &q.RWMutex}, nil
}

type memoryJobIter struct {
	jobs *[]*Job
	idx  *int
	*sync.RWMutex
}

func (i *memoryJobIter) Next() (*Job, error) {
	i.Lock()
	defer i.Unlock()
	if len(*i.jobs) <= *i.idx {
		return nil, nil
	}
	j := (*i.jobs)[*i.idx]
	(*i.idx)++
	return j, nil
}

func (i *memoryJobIter) Close() error {
	return nil
}
