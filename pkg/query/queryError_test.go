package query

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorToString(t *testing.T) {
	errorString := "Test error"
	id := "0"
	json := fmt.Sprintf(`{"ID":"%s","Data":{"Error":"%s"}}`, id, errorString)
	e := NewError(id, errorString)
	s := e.ToString()
	assert.Equal(t, s, json)
}
