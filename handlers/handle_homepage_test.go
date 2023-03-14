package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thisdougb/cleango/pkg/usecase/ourpurpose"
)

func TestHomePage(t *testing.T) {

	// create our mock service
	r := ourpurpose.NewMockRepository()
	ourPurposeService := ourpurpose.NewService(r)

	// inject mock service
	env := &Env{OurPurposeService: ourPurposeService}

	// httptest lets us interrogate the http response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.HomePage)

	bodyReader := bytes.NewReader([]byte(nil))

	req, err := http.NewRequest("GET", "/", bodyReader)
	if err != nil {
		assert.Fail(t, "Homepage")
	}

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Homepage")
}
