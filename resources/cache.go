package resources

import (
	"time"

	"github.com/fgunawan1995/xendit/config"
	cache "github.com/patrickmn/go-cache"
)

const (
	cacheCleanup = 5 * time.Minute
)

//InitCache create new cache
func InitCache(cfg *config.Config) *cache.Cache {
	return cache.New(time.Duration(cfg.Cache.DefaultExpirationInMinutes)*time.Minute, cacheCleanup)
}
