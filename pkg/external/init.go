// Package external is function to request 3rd party api
package external

import (
	"context"
	"time"
)

type apiCaller struct{}

// APICaller is interface
type APICaller interface {
	Consume(ctx context.Context, request, recv interface{}, header map[string]string, method string, uri string, timeout *time.Duration) (int, error)
}

// InitAPICaller : initial struct
func InitAPICaller() APICaller {
	return &apiCaller{}
}

func (s *apiCaller) Consume(ctx context.Context, request, recv interface{}, header map[string]string, method string, uri string, timeout *time.Duration) (int, error) {
	return s.caller(ctx, request, recv, header, method, uri, timeout)
}
