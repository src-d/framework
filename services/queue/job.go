package queue

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/vmihailenco/msgpack.v2"
)

type contentType string

const msgpackContentType contentType = "application/msgpack"

type Job struct {
	Id        string
	Priority  Priority
	Timestamp time.Time

	contentType contentType
	raw         []byte
}

func NewJob() *Job {
	return &Job{
		Id:          bson.NewObjectId().Hex(),
		Priority:    PriorityNormal,
		Timestamp:   time.Now(),
		contentType: msgpackContentType,
	}
}

func (j *Job) Encode(payload interface{}) error {
	var err error
	j.raw, err = encode(msgpackContentType, &payload)
	if err != nil {
		return err
	}

	return nil
}

func (j *Job) Decode(payload interface{}) error {
	return decode(msgpackContentType, j.raw, &payload)
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
