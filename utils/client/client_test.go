package client

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	. "gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ClientSuite struct {
	session *mgo.Session
	ch      chan bool
}

const serverAddr = "127.0.0.1:6060"

var _ = Suite(&ClientSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *ClientSuite) getCollection() *mgo.Collection {
	return s.session.DB("test").C("server-test")
}

func (s *ClientSuite) SetUpSuite(c *C) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic("A local Mongo is required in order to run client tests.")
	}
	s.session = session.Clone()
	s.ch = make(chan bool, 10)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c.Log("Call received")
		s.ch <- true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test"))
	})
	go func() {
		if err := http.ListenAndServe(serverAddr, nil); err != nil {
			panic(err)
		}
	}()
	s.clean()
}

func (s *ClientSuite) clean() {
	if _, err := s.getCollection().RemoveAll(bson.M{}); err != nil {
		panic("Cannot clean connection")
	}
}

func (s *ClientSuite) TearDownSuite(c *C) {
	defer s.session.Close()
	s.clean()
}

func (s *ClientSuite) TestServer(c *C) {
	var addr = fmt.Sprintf("http://%s/", serverAddr)
	client := New(s.getCollection(), 1000000)
	for i := 0; i < 5; i++ {
		c.Logf("Calling to %s, %d attempt/s", addr, i+1)
		_, err := client.Get(addr)
		c.Assert(err, IsNil)
		var cached bool
		select {
		case <-time.After(time.Second / 4):
			cached = true
		case <-s.ch:
			cached = false
		}
		c.Assert(cached, Equals, i > 0)
	}

}
