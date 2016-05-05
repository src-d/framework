package queue

import (
	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

var _ = Suite(&MemorySuite{})

type MemorySuite struct{}

func (s *MemorySuite) TestIntegration(c *C) {
	b := NewMemoryBroker()

	q, err := b.Queue(bson.NewObjectId().Hex())
	c.Assert(err, IsNil)

	job := NewJob()
	job.Encode(true)
	err = q.Publish(job)
	c.Assert(err, IsNil)

	for i := 0; i < 100; i++ {
		job := NewJob()
		job.Encode(true)
		err = q.Publish(job)
		c.Assert(err, IsNil)
	}

	i, err := q.Consume()
	c.Assert(err, IsNil)

	retrievedJob, err := i.Next()
	c.Assert(err, IsNil)

	var payload bool
	err = retrievedJob.Decode(&payload)
	c.Assert(err, IsNil)
	c.Assert(payload, Equals, true)

	c.Assert(retrievedJob.ID, Equals, job.ID)
	c.Assert(retrievedJob.Priority, Equals, job.Priority)
	c.Assert(retrievedJob.Timestamp.Second(), Equals, job.Timestamp.Second())

	err = i.Close()
	c.Assert(err, IsNil)

	err = b.Close()
	c.Assert(err, IsNil)
}
