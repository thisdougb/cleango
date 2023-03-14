//go:build dev || test
// +build dev test

package ourpurpose

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOurPurpose(t *testing.T) {

	mockDatastore := NewMockRepository()
	s := NewService(mockDatastore)

	msg := s.OurPurpose()
	assert.Equal(t, "Hello World!", msg, "Checks the message")

}
