package sopsearch

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type ElasticClient struct {
	esclient *elasticsearch.Client
}

func NewElasticClient(addresses []string, username, password string) (ElasticClient, error) {

	cfg := elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return ElasticClient{}, err
	}

	return ElasticClient{client}, nil
}

func (ec *ElasticClient) Info() (*esapi.Response, error) {
	res, err := ec.esclient.Info()
	return res, err
}

func (ec *ElasticClient) CreateOrUpdateIndex(index, documentID, body string) error {

	// Instantiate a request object
	req := esapi.IndexRequest{
		Index:      index,
		DocumentID: documentID,
		Body:       strings.NewReader(body),
		Refresh:    "true",
	}
	ctx := context.Background()
	// Return an API response object from request
	res, err := req.Do(ctx, ec.esclient)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
