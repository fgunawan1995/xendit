package usecase

import (
	"testing"

	"github.com/fgunawan1995/xendit/config"
	marveldal "github.com/fgunawan1995/xendit/dal/api/marvel"
	cachedal "github.com/fgunawan1995/xendit/dal/cache"
)

func TestNew(t *testing.T) {
	type args struct {
		cfg    *config.Config
		marvel marveldal.MarvelDAL
		cache  cachedal.CacheDAL
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.cfg, tt.args.marvel, tt.args.cache)
		})
	}
}
