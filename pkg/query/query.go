package query

import (
	"encoding/json"
	"errors"
	"time"

	aero "github.com/aerospike/aerospike-client-go"
)

const (

	//ErrorUnsupportedQueryType is returned when the query type is not recognized
	ErrorUnsupportedQueryType = "[ERROR] Unsupported Query type"

	//ErrorMalformedQuery is returned when a malformed query is received
	ErrorMalformedQuery = "[ERROR] Malformed Query"
)

//DatabaseQuery describe a database query
type DatabaseQuery struct {
	Type      string
	Namespace string
	Set       string
	Key       string
	Bins      []Bin
}

//Bin describe a bin
type Bin struct {
	BinName  string
	BinValue string
}

//NewDatabaseQuery create a new database query
func NewDatabaseQuery(queryJSON string) (DatabaseQuery, error) {

	dq := DatabaseQuery{}
	err := json.Unmarshal([]byte(queryJSON), &dq)
	if err != nil {
		return dq, err
	}

	if (dq.Type != "get") && (dq.Type != "put") {
		return dq, errors.New(ErrorUnsupportedQueryType)
	}

	if dq.Namespace == "" || dq.Set == "" || dq.Key == "" {
		return dq, errors.New(ErrorMalformedQuery)
	}

	if dq.Type == "put" && len(dq.Bins) == 0 {
		return dq, errors.New(ErrorMalformedQuery)
	}

	return dq, nil
}

//ExecuteGetQuery execute the get query on the db
func (dq DatabaseQuery) ExecuteGetQuery() (aero.BinMap, error) {

	client, err := aero.NewClient("127.0.0.1", 3000)
	if err != nil {
		panic(err)
	}

	key, err := aero.NewKey(dq.Namespace, dq.Set, dq.Key)
	if err != nil {
		panic(err)
	}

	policy := aero.NewPolicy()
	policy.SocketTimeout = 50 * time.Millisecond

	rec, err := client.Get(policy, key)
	if err != nil {
		panic(err)
	}

	return rec.Bins, nil
}

//ExecutePutQuery execute the put query on the db
func (dq DatabaseQuery) ExecutePutQuery() (string, error) {

	client, err := aero.NewClient("127.0.0.1", 3000)
	if err != nil {
		panic(err)
	}

	key, err := aero.NewKey(dq.Namespace, dq.Set, dq.Key)
	if err != nil {
		panic(err)
	}

	policy := aero.NewWritePolicy(0, 0)
	policy.SocketTimeout = 50 * time.Millisecond

	var bins = make(aero.BinMap, len(dq.Bins))

	for _, b := range dq.Bins {
		bins[b.BinName] = b.BinValue
	}

	err = client.Put(policy, key, bins)
	if err != nil {
		panic(err)
	}

	return key.String(), nil

}
