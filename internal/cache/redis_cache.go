package cache

import (
	"strings"
	"time"

	"github.com/denlipov/com-request-api/internal/config"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

// NewRedisCache ...
func NewRedisCache(cfg config.Redis) *cache.Cache {
	log.Debug().Msgf("Using Redis config: %+v", cfg)

	addrMap := make(map[string]string)
	for _, addr := range cfg.Addrs {
		hostPort := strings.FieldsFunc(addr, func(ch rune) bool {
			return ch == ':'
		})
		addrMap[hostPort[0]] = ":" + hostPort[1]
	}
	opts := redis.RingOptions{
		Addrs: addrMap,
	}
	ring := redis.NewRing(&opts)
	redisCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return redisCache
}
