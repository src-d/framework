package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestMemorySuite(t *testing.T) {
	suite.Run(t, new(MemorySuite))
}

type MemorySuite struct {
	QueueSuite
}

func (s *MemorySuite) SetupSuite() {
	s.BrokerURI = testMemoryURI
}

func (s *MemorySuite) TestIntegration() {
	assert := assert.New(s.T())

	qName := newName()
	q, err := s.Broker.Queue(qName)
	assert.NoError(err)
	assert.NotNil(q)

	j := NewJob()
	j.Encode(true)
	err = q.Publish(j)
	assert.NoError(err)

	for i := 0; i < 100; i++ {
		job := NewJob()
		job.Encode(true)
		err = q.Publish(job)
		assert.NoError(err)
	}

	iter, err := q.Consume()
	assert.NoError(err)

	retrievedJob, err := iter.Next()
	assert.NoError(err)
	assert.NoError(retrievedJob.Ack())

	var payload bool
	err = retrievedJob.Decode(&payload)
	assert.NoError(err)
	assert.True(payload)

	assert.Equal(j.tag, retrievedJob.tag)
	assert.Equal(j.Priority, retrievedJob.Priority)
	assert.Equal(j.Timestamp.Second(), retrievedJob.Timestamp.Second())

	err = iter.Close()
	assert.NoError(err)
}
