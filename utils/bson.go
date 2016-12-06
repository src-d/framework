package utils

import (
	"encoding/json"
	"regexp"

	"gopkg.in/mgo.v2/bson"
)

// A un/marshaller of BSON data that just keeps the JSON representation in bytes.
type ByteBSON struct {
	Bytes []byte
}

func (b *ByteBSON) SetBSON(raw bson.Raw) error {
	var out interface{}
	raw.Unmarshal(&out)
	var err error
	b.Bytes, err = json.Marshal(out)
	return err
}

func (b *ByteBSON) GetBSON() (interface{}, error) {
	var ret interface{}
	err := json.Unmarshal(b.Bytes, &ret)
	return ret, err
}

func JSONToBSON(s string) bson.Raw {
	var jsout interface{}
	err := json.Unmarshal([]byte(s), &jsout)
	if err != nil {
		panic(err)
	}
	bsout, err := bson.Marshal(jsout)
	if err != nil {
		panic(err)
	}
	return bson.Raw{Kind: 0x03, Data: bsout}
}

type Regexp struct {
	*regexp.Regexp
}

func (r *Regexp) GetBSON() (interface{}, error) {
	return r.String(), nil
}

func (r *Regexp) SetBSON(raw bson.Raw) error {
	var s string
	err := raw.Unmarshal(&s)
	if err != nil {
		return err
	}
	rgx, err := regexp.Compile(s)
	if err != nil {
		return err
	}
	r.Regexp = rgx
	return nil
}
