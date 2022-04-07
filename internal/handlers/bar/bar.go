package bar

import (
	"context"
	"net/http"
)

type service interface {
	Bar(ctx context.Context, bar string) ([]string, error)
}

type barHandler struct {
	barService service
}

func New(service service) *barHandler {
	return &barHandler{
		barService: service,
	}
}

func (b *barHandler) Handle(w http.ResponseWriter, req *http.Request) {
	// read vars from request

	_, _ = b.barService.Bar(context.Background(), "some_value_from_request")

	// return results
}
