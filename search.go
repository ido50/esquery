package esquery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type SearchRequest struct {
	query Mappable
	aggs  []Aggregation
	size  uint64
}

func Search() *SearchRequest {
	return &SearchRequest{}
}

func (req *SearchRequest) Query(q Mappable) *SearchRequest {
	req.query = q
	return req
}

func (req *SearchRequest) Aggs(aggs ...Aggregation) *SearchRequest {
	req.aggs = aggs
	return req
}

func (req *SearchRequest) Size(size uint64) *SearchRequest {
	req.size = size
	return req
}

func (req *SearchRequest) Map() map[string]interface{} {
	aggs := make(map[string]interface{})

	for _, agg := range req.aggs {
		aggs[agg.Name()] = agg.Map()
	}

	return map[string]interface{}{
		"query": req.query.Map(),
		"aggs":  aggs,
		"size":  req.size,
	}
}

// Run executes the request using the provided ElasticSearch client. Zero or
// more search options can be provided as well. It returns the standard Response
// type of the official Go client.
func (req *SearchRequest) Run(
	api *elasticsearch.Client,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	return req.RunSearch(api.Search, o...)
}

// RunSearch is the same as the Run method, except that it accepts a value of
// type esapi.Search (usually this is the Search field of an elasticsearch.Client
// object). Since the ElasticSearch client does not provide an interface type
// for its API (which would allow implementation of mock clients), this provides
// a workaround. The Search function in the ES client is actually a field of a
// function type.
func (req *SearchRequest) RunSearch(
	search esapi.Search,
	o ...func(*esapi.SearchRequest),
) (res *esapi.Response, err error) {
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(req.Map())
	if err != nil {
		return nil, err
	}

	fmt.Fprintf(os.Stderr, "QUERY: %s\n", b.String())

	opts := append([]func(*esapi.SearchRequest){search.WithBody(&b)}, o...)

	return search(opts...)
}
