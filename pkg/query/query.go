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
	ID   string
	Data Data
}

//Data describe the data structure of a query
type Data struct {
	Type      string
	Namespace string
	Set       string
	Key       string
	Bins      []Bin
}

//DatabaseGetReply describe a database reply for a get request
type DatabaseGetReply struct {
	ID   string
	Data []Bin
}

//DatabasePutReply describe a database reply for a put request
type DatabasePutReply struct {
	ID   string
	Data string
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
	return dq, nil
}

//CheckValues check the value of the requested query
func (dq DatabaseQuery) CheckValues() error {
	if dq.ID == "" {
		return errors.New(ErrorMalformedQuery)
	}

	if (dq.Data.Type != "get") && (dq.Data.Type != "put") {
		return errors.New(ErrorUnsupportedQueryType)
	}

	if dq.Data.Namespace == "" || dq.Data.Set == "" || dq.Data.Key == "" {
		return errors.New(ErrorMalformedQuery)
	}

	if dq.Data.Type == "put" && len(dq.Data.Bins) == 0 {
		return errors.New(ErrorMalformedQuery)
	}
	return nil
}

//ExecuteGetQuery execute the get query on the db
func (dq DatabaseQuery) ExecuteGetQuery() (string, error) {
	var bin Bin
	var dgr DatabaseGetReply
	var bins []Bin

	client, err := aero.NewClient("127.0.0.1", 3000)
	if err != nil {
		return "", err
	}

	key, err := aero.NewKey(dq.Data.Namespace, dq.Data.Set, dq.Data.Key)
	if err != nil {
		return "", err
	}

	policy := aero.NewPolicy()
	policy.SocketTimeout = 50 * time.Millisecond

	rec, err := client.Get(policy, key)
	if err != nil {
		return "", err
	}

	for k, b := range rec.Bins {
		bin.BinName = k
		bin.BinValue = b.(string)
		bins = append(bins, bin)
	}

	dgr = DatabaseGetReply{
		ID:   dq.ID,
		Data: bins,
	}

	res, err := json.Marshal(dgr)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

//ExecutePutQuery execute the put query on the db
func (dq DatabaseQuery) ExecutePutQuery() (string, error) {

	client, err := aero.NewClient("127.0.0.1", 3000)
	if err != nil {
		return "", err
	}

	key, err := aero.NewKey(dq.Data.Namespace, dq.Data.Set, dq.Data.Key)
	if err != nil {
		return "", err
	}

	policy := aero.NewWritePolicy(0, 0)
	policy.SocketTimeout = 50 * time.Millisecond

	var bins = make(aero.BinMap, len(dq.Data.Bins))

	for _, b := range dq.Data.Bins {
		bins[b.BinName] = b.BinValue
	}

	err = client.Put(policy, key, bins)
	if err != nil {
		panic(err)
	}

	dpr := DatabasePutReply{
		ID:   dq.ID,
		Data: key.String(),
	}

	res, err := json.Marshal(dpr)
	if err != nil {
		return "", err
	}
	return string(res), nil

}
