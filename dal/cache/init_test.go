package cache

import (
	"testing"

	cache "github.com/patrickmn/go-cache"
)

func TestNew(t *testing.T) {
	type args struct {
		cache *cache.Cache
	}
	tests := []struct {
		name string
		args args
		want CacheDAL
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.cache)
		})
	}
}
