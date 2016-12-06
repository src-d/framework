package client

import (
	"fmt"
	"net/http"

	"github.com/gregjones/httpcache"
	"github.com/mcuadros/go-mgo-cache"

	"gopkg.in/mgo.v2"
)

type Client struct {
	http.Client
}

func New(c *mgo.Collection, cache int) *Client {
	t := httpcache.NewTransport(mgocache.New(c))
	t.Transport = &cacheAdder{http.DefaultTransport, fmt.Sprintf("max-age=%d", cache)}
	return &Client{http.Client{Transport: t}}
}

type cacheAdder struct {
	Transport http.RoundTripper
	cache     string
}

func (t *cacheAdder) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	resp.Header.Set("cache-control", t.cache)
	return resp, nil
}
