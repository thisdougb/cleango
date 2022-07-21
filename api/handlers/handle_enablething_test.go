package handlers

import (
	"bytes"
	"github.com/thisdougb/cleango/pkg/usecase/enablething"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// we can't really test the html output here
var TestEnableThingItems = []struct {
	comment      string
	httpURL      string
	httpMethod   string
	bodyData     string
	expectStatus int
}{
	{
		comment:      "valid request",
		httpURL:      "http://localhost/thing/enable/",
		httpMethod:   "POST",
		bodyData:     `{"thing_id":1}`,
		expectStatus: 200,
	},
	{
		comment:      "invalid http method",
		httpURL:      "http://localhost/thing/enable/",
		httpMethod:   "GET",
		bodyData:     `{"thing_id":1}`,
		expectStatus: 405,
	},
	{
		comment:      "incorrect json data",
		httpURL:      "http://localhost/thing/enable/",
		httpMethod:   "POST",
		bodyData:     `{"thing_id":"two"}`,
		expectStatus: 400,
	},
	{
		comment:      "thing does not exist",
		httpURL:      "http://localhost/thing/enable/",
		httpMethod:   "POST",
		bodyData:     `{"thing_id":2}`,
		expectStatus: 404,
	},
	{
		comment:      "trigger datastore error",
		httpURL:      "http://localhost/thing/enable/",
		httpMethod:   "POST",
		bodyData:     `{"thing_id":3}`,
		expectStatus: 500,
	},
}

func TestMagicLinkRequestWeb(t *testing.T) {

	// create our mock service
	r := enablething.NewMockRepository()
	enableThingService := enablething.NewService(r)

	// inject mock service
	env := &Env{EnableThingService: enableThingService}

	for _, item := range TestEnableThingItems {

		// httptest lets us interrogate the http response
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(env.EnableThing)

		bodyReader := bytes.NewReader([]byte(item.bodyData))

		req, err := http.NewRequest(item.httpMethod, item.httpURL, bodyReader)
		if err != nil {
			assert.Fail(t, item.comment)
		}
		req.Header.Add("Content-Type", "encoding/json")

		handler.ServeHTTP(rr, req)

		assert.Equal(t, item.expectStatus, rr.Code, item.comment)
	}
}
