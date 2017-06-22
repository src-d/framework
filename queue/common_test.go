package queue

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var testRand *rand.Rand

func init() {
	testRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func newName() string {
	return fmt.Sprintf("queue_tests_%d", testRand.Int())
}

const (
	testAMQPURI   = "amqp://localhost:5672"
	testMemoryURI = "memory://"
)

func TestNewBroker(t *testing.T) {
	assert := assert.New(t)

	b, err := NewBroker(testAMQPURI)
	assert.NoError(err)
	assert.IsType(&AMQPBroker{}, b)
	assert.NoError(b.Close())

	b, err = NewBroker("amqp://badurl")
	assert.Error(err)

	b, err = NewBroker(testMemoryURI)
	assert.NoError(err)
	assert.IsType(&memoryBroker{}, b)
	assert.NoError(b.Close())

	b, err = NewBroker("memory://anything")
	assert.NoError(err)
	assert.IsType(&memoryBroker{}, b)
	assert.NoError(b.Close())

	b, err = NewBroker("badproto://badurl")
	assert.Equal(ErrUnsupportedProtocol, err)

	b, err = NewBroker("foo://host%10")
	assert.Error(err)
}

type QueueSuite struct {
	suite.Suite
	r rand.Rand

	TxNotSupported bool
	BrokerURI      string

	Broker Broker
}

func (s *QueueSuite) SetupTest() {
	b, err := NewBroker(s.BrokerURI)
	s.NoError(err)
	s.Broker = b
}

func (s *QueueSuite) TearDownTest() {
	s.NoError(s.Broker.Close())
}

func (s *QueueSuite) TestConsume_empty() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	iter, err := q.Consume()
	assert.NoError(err)
	assert.NotNil(iter)

	assert.NoError(iter.Close())
}

func (s *QueueSuite) TestJobIter_Next_empty() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	iter, err := q.Consume()
	assert.NoError(err)
	assert.NotNil(iter)

	done := s.checkNextClosed(iter)
	assert.NoError(iter.Close())
	<-done
}

func (s *QueueSuite) TestJob_Reject_no_requeue() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	j := NewJob()
	err = j.Encode(1)
	assert.NoError(err)

	err = q.Publish(j)
	assert.NoError(err)

	iter, err := q.Consume()
	assert.NoError(err)
	assert.NotNil(iter)

	j, err = iter.Next()
	assert.NoError(err)
	assert.NotNil(j)

	err = j.Reject(false)
	assert.NoError(err)

	done := s.checkNextClosed(iter)
	<-time.After(50 * time.Millisecond)
	assert.NoError(iter.Close())
	<-done
}

func (s *QueueSuite) TestJob_Reject_requeue() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	j := NewJob()
	err = j.Encode(1)
	assert.NoError(err)

	err = q.Publish(j)
	assert.NoError(err)

	iter, err := q.Consume()
	assert.NoError(err)
	assert.NotNil(iter)

	j, err = iter.Next()
	assert.NoError(err)
	assert.NotNil(j)

	err = j.Reject(true)
	assert.NoError(err)

	j, err = iter.Next()
	assert.NoError(err)
	assert.NotNil(j)

	assert.NoError(iter.Close())
}

func (s *QueueSuite) TestPublish_nil() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	err = q.Publish(nil)
	assert.Equal(ErrEmptyJob, err)
}

func (s *QueueSuite) TestPublish_empty() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	err = q.Publish(&Job{})
	assert.Equal(ErrEmptyJob, err)
}

func (s *QueueSuite) TestPublishDelayed_nil() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	err = q.PublishDelayed(nil, time.Second)
	assert.Equal(ErrEmptyJob, err)
}

func (s *QueueSuite) TestPublishDelayed_empty() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	err = q.PublishDelayed(&Job{}, time.Second)
	assert.Equal(ErrEmptyJob, err)
}

func (s *QueueSuite) TestPublishAndConsume_immediate_ack() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	var (
		ids        []string
		priorities []Priority
		timestamps []time.Time
	)
	for i := 0; i < 100; i++ {
		j := NewJob()
		err = j.Encode(i)
		assert.NoError(err)
		err = q.Publish(j)
		assert.NoError(err)
		ids = append(ids, j.ID)
		priorities = append(priorities, j.Priority)
		timestamps = append(timestamps, j.Timestamp)
	}

	iter, err := q.Consume()
	assert.NoError(err)
	assert.NotNil(iter)

	for i := 0; i < 100; i++ {
		j, err := iter.Next()
		assert.NoError(err)
		assert.NoError(j.Ack())

		var payload int
		assert.NoError(j.Decode(&payload))
		assert.Equal(i, payload)

		assert.Equal(ids[i], j.ID)
		assert.Equal(priorities[i], j.Priority)
		assert.Equal(timestamps[i].Unix(), j.Timestamp.Unix())
	}

	done := s.checkNextClosed(iter)
	assert.NoError(iter.Close())
	<-done
}

func (s *QueueSuite) TestDelayed() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	j := NewJob()
	err = j.Encode("hello")
	assert.NoError(err)
	err = q.PublishDelayed(j, 1*time.Second)
	assert.NoError(err)

	iter, err := q.Consume()
	assert.NoError(err)

	start := time.Now()
	var since time.Duration
	for {
		j, err := iter.Next()
		assert.NoError(err)
		if j == nil {
			<-time.After(300 * time.Millisecond)
			continue
		}

		since = time.Since(start)

		var payload string
		assert.NoError(j.Decode(&payload))
		assert.Equal("hello", payload)
		break
	}

	assert.True(since >= 1*time.Second)
}

func (s *QueueSuite) TestTransaction_Error() {
	if s.TxNotSupported {
		s.T().Skip("transactions not supported")
	}

	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	err = q.Transaction(func(qu Queue) error {
		job := NewJob()
		assert.NoError(job.Encode("goodbye"))
		assert.NoError(qu.Publish(job))
		return errors.New("foo")
	})
	assert.Error(err)

	i, err := q.Consume()
	assert.NoError(err)
	done := s.checkNextClosed(i)
	<-time.After(50 * time.Millisecond)
	assert.NoError(i.Close())
	<-done
}

func (s *QueueSuite) TestTransaction() {
	if s.TxNotSupported {
		s.T().Skip("transactions not supported")
	}

	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	err = q.Transaction(func(q Queue) error {
		job := NewJob()
		assert.NoError(job.Encode("hello"))
		assert.NoError(q.Publish(job))
		return nil
	})
	assert.NoError(err)

	iter, err := q.Consume()
	assert.NoError(err)
	j, err := iter.Next()
	assert.NoError(err)
	assert.NotNil(j)
	var payload string
	assert.NoError(j.Decode(&payload))
	assert.Equal("hello", payload)
	assert.NoError(iter.Close())
}

func (s *QueueSuite) TestTransaction_not_supported() {
	assert := assert.New(s.T())

	if !s.TxNotSupported {
		s.T().Skip("transactions supported")
	}

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	err = q.Transaction(nil)
	assert.Equal(ErrTxNotSupported, err)
}

func (s *QueueSuite) TestRetryQueue() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	// 1: Publish jobs to the main queue.
	j1 := NewJob()
	err = j1.Encode(1)
	assert.NoError(err)

	err = q.Publish(j1)
	assert.NoError(err)

	j2 := NewJob()
	err = j2.Encode(2)
	assert.NoError(err)
	err = q.Publish(j2)
	assert.NoError(err)

	// 2: consume and reject them.
	iterMain, err := q.Consume()
	assert.NoError(err)
	assert.NotNil(iterMain)

	jReject1, err := iterMain.Next()
	assert.NoError(err)
	assert.NotNil(jReject1)
	// Jobs should go to the retry queue when rejected with requeue = false
	err = jReject1.Reject(false)
	assert.NoError(err)

	jReject2, err := iterMain.Next()
	assert.NoError(err)
	assert.NotNil(jReject2)
	err = jReject2.Reject(false)
	assert.NoError(err)

	// 3. republish the jobs in the retry queue.
	err = q.RepublishBuried()
	assert.NoError(err)

	// 4. re-read the jobs on the main queue.
	var payload int
	jRepub1, err := iterMain.Next()
	assert.NoError(jRepub1.Decode(&payload))
	assert.Equal(1, payload)
	assert.NoError(jRepub1.Ack())

	jRepub2, err := iterMain.Next()
	assert.NoError(jRepub2.Decode(&payload))
	assert.Equal(2, payload)
	assert.NoError(jRepub2.Ack())

	done := s.checkNextClosed(iterMain)
	assert.NoError(iterMain.Close())
	iterMain.Close()
	<- done
}

func (s *QueueSuite) checkNextClosed(iter JobIter) chan struct{} {
	assert := assert.New(s.T())

	done := make(chan struct{})
	go func() {
		j, err := iter.Next()
		assert.Equal(ErrAlreadyClosed, err)
		assert.Nil(j)
		done <- struct{}{}
	}()
	return done
}
