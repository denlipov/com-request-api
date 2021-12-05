package api

import (
	"github.com/go-redis/cache/v8"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/denlipov/com-request-api/internal/repo"

	pb "github.com/denlipov/com-request-api/pkg/com-request-api"
)

var (
	totalRequestNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "com_request_api_request_not_found_total",
		Help: "Total number of requests that were not found",
	})
)

type requestAPI struct {
	pb.UnimplementedComRequestApiServiceServer
	repo  repo.Repo
	cache *cache.Cache
}

// NewRequestAPI returns api of com-request-api service
func NewRequestAPI(r repo.Repo, cache *cache.Cache) pb.ComRequestApiServiceServer {
	return &requestAPI{
		repo:  r,
		cache: cache,
	}
}
