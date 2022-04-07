package foo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const URL = "https://cafu.com/someService"

type service struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) *service {
	return &service{
		httpClient: httpClient,
	}
}

// Foo is external service call
func (c *service) Foo(ctx context.Context) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	resp, err := c.httpClient.Get(URL + "/foo")
	if err != nil {
		return nil, fmt.Errorf("someService/foo error: %w", err)
	}

	return resp, nil
}
