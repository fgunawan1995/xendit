package resources

import (
	"testing"

	"github.com/fgunawan1995/xendit/config"
	cache "github.com/patrickmn/go-cache"
)

func TestInitCache(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want *cache.Cache
	}{

		{
			name: "success",
			args: args{
				cfg: &config.Config{
					Cache: config.CacheConfig{
						DefaultExpirationInMinutes: 60,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitCache(tt.args.cfg)
		})
	}
}
