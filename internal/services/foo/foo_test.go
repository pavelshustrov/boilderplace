package foo

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var (
	errorResponse = func(req *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("error happened")
	}
	statusOKResponse = func(req *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
		}, nil
	}
)

func TestErrorReturn(t *testing.T) {
	testHttpClient := NewTestClient(errorResponse)

	service := New(testHttpClient)
	foo, err := service.Foo(context.Background())
	assert.Nil(t, foo)
	assert.NotNil(t, err)
}

func TestStatusOKReturn(t *testing.T) {
	testHttpClient := NewTestClient(statusOKResponse)

	service := New(testHttpClient)
	foo, err := service.Foo(context.Background())
	require.NotNil(t, foo)
	resp := foo.(*http.Response)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
	require.NoError(t, err)
}
