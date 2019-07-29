package query

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGetQuery(t *testing.T) {
	queryID := "123"
	queryType := "get"
	queryNameSpace := "on-disk-db"
	querySet := "test"
	queryKey := "7157762ecb437e34e5770667d086887532a304d9152b171a974d100f2fcfb687"
	queryJSON := fmt.Sprintf("{\"ID\" : \"%s\", \"Data\": { \"Type\" : \"%s\", \"Namespace\" : \"%s\", \"Set\" : \"%s\", \"Key\" : \"%s\"}}", queryID, queryType, queryNameSpace, querySet, queryKey)
	dq, err := NewDatabaseQuery(queryJSON)
	assert.Nil(t, err)
	assert.Equal(t, queryID, dq.ID)
	assert.Equal(t, queryType, dq.Data.Type)
	assert.Equal(t, queryNameSpace, dq.Data.Namespace)
	assert.Equal(t, querySet, dq.Data.Set)
	assert.Equal(t, queryKey, dq.Data.Key)
}

func TestNewQueryWorngType(t *testing.T) {
	queryJSON := "{\"ID\" : \"123\", \"Data\": { \"Type\" : \"unkown\", \"Namespace\" : \"on-disk-db\", \"Set\" : \"test\", \"Key\" : \"7157762ecb437e34e5770667d086887532a304d9152b171a974d100f2fcfb687\"}}"
	_, err := NewDatabaseQuery(queryJSON)
	assert.Equal(t, err.Error(), ErrorUnsupportedQueryType)
}

func TestNewQueryWithoutID(t *testing.T) {
	queryJSON := "{\"Data\": { \"Type\" : \"unkown\", \"Namespace\" : \"on-disk-db\", \"Set\" : \"test\", \"Key\" : \"7157762ecb437e34e5770667d086887532a304d9152b171a974d100f2fcfb687\"}}"
	_, err := NewDatabaseQuery(queryJSON)
	assert.Equal(t, err.Error(), ErrorMalformedQuery)
}

func TestNewPutQueryWithoutBins(t *testing.T) {
	queryJSON := "{\"Data\": { \"Type\" : \"put\", \"Namespace\" : \"on-disk-db\", \"Set\" : \"test\", \"Key\" : \"7157762ecb437e34e5770667d086887532a304d9152b171a974d100f2fcfb687\"}}"
	_, err := NewDatabaseQuery(queryJSON)
	assert.Equal(t, err.Error(), ErrorMalformedQuery)
}
