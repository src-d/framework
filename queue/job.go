package queue

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type contentType string

const msgpackContentType contentType = "application/msgpack"

// Job contains the information for a job to be published to a queue.
type Job struct {
	// ID of the job.
	ID string
	// Priority is the priority level.
	Priority Priority
	// Timestamp is the time of creation.
	Timestamp time.Time

	contentType  contentType
	raw          []byte
	acknowledger Acknowledger
	tag          uint64
}

// Acknowledger represents the object in charge of acknowledgement management for a
// job.
type Acknowledger interface {
	// Ack is called when the Job has finished.
	Ack() error
	// Reject is called if the job has errored. The parameter indicates whether the
	// job should be put back in the queue or not.
	Reject(requeue bool) error
}

// NewJob creates a new Job with default values, a new unique ID and current
// timestamp.
func NewJob() *Job {
	return &Job{
		ID:          bson.NewObjectId().Hex(),
		Priority:    PriorityNormal,
		Timestamp:   time.Now(),
		contentType: msgpackContentType,
	}
}

// Encode encodes the payload to the wire format used.
func (j *Job) Encode(payload interface{}) error {
	var err error
	j.raw, err = encode(msgpackContentType, &payload)
	if err != nil {
		return err
	}

	return nil
}

// Decode decodes the payload from the wire format.
func (j *Job) Decode(payload interface{}) error {
	return decode(msgpackContentType, j.raw, &payload)
}

var errCantAck = errors.New("can't acknowledge this message, it does not come from a queue")

// Ack is called when the job is finished.
func (j *Job) Ack() error {
	if j.acknowledger == nil {
		return errCantAck
	}
	return j.acknowledger.Ack()
}

// Reject is called when the job errors. The parameter is true if and only if the
// job should be put back in the queue.
func (j *Job) Reject(requeue bool) error {
	if j.acknowledger == nil {
		return errCantAck
	}
	return j.acknowledger.Reject(requeue)
}

func encode(mime contentType, p interface{}) ([]byte, error) {
	switch mime {
	case msgpackContentType:
		return msgpack.Marshal(p)
	default:
		return nil, fmt.Errorf("unknown content type: %s", mime)
	}
}

func decode(mime contentType, r []byte, p interface{}) error {
	switch mime {
	case msgpackContentType:
		return msgpack.Unmarshal(r, p)
	default:
		return fmt.Errorf("unknown content type: %s", mime)
	}
}
