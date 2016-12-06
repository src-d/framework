package configurable

// General initializators of differente connections, sessions, etc.

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/olivere/elastic.v3"
)

/*
	Creates a new session to db given by the url with a given timeout

	Panics if an error occurs
*/
func CreateMgoSession(url string, timeout time.Duration) *mgo.Session {
	session, err := mgo.DialWithTimeout(
		url,
		timeout,
	)
	checkError(err)

	session.SetMode(mgo.Monotonic, true)
	return session
}

type authTransport struct {
	user     string
	password string
}

func (t *authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(t.user, t.password)
	return http.DefaultTransport.RoundTrip(r)
}

func (t *authTransport) CancelRequest(r *http.Request) {
	http.DefaultTransport.(*http.Transport).CancelRequest(r)
}

/*
	Creates a new Elasticsearch client given the url, user and password

	Panics if an error occurs
*/
func CreateElasticsearchClient(url, user, password string) *elastic.Client {
	httpClient := &http.Client{
		Transport: &authTransport{
			user:     user,
			password: password,
		},
	}

	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHttpClient(httpClient),
		elastic.SetHealthcheckTimeoutStartup(time.Second*10),
		elastic.SetMaxRetries(3),
	)
	checkError(err)

	return client
}

// TODO: Should we make this public so we can share a way to deal with errors?
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
