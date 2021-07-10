package cache

import (
	"reflect"
	"testing"

	"github.com/fgunawan1995/xendit/model"
	cache "github.com/patrickmn/go-cache"
)

func Test_impl_GetAllMarvelCharacterIDs(t *testing.T) {
	mCache := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	tests := []struct {
		name    string
		want    []int64
		wantErr bool
		mock    func()
	}{
		{
			name: "not found",
			mock: func() {
				mCache.Delete(marvelCharacterIDs)
			},
		},
		{
			name: "success",
			mock: func() {
				mCache.SetDefault(marvelCharacterIDs, []int64{1, 2, 3})
			},
			want: []int64{
				1, 2, 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				cacheObj: mCache,
			}
			tt.mock()
			got, err := dal.GetAllMarvelCharacterIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetAllMarvelCharacterIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetAllMarvelCharacterIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_SaveAllMarvelCharacterIDs(t *testing.T) {
	mCache := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	type args struct {
		characterIDs []int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				cacheObj: mCache,
			}
			dal.SaveAllMarvelCharacterIDs(tt.args.characterIDs)
		})
	}
}

func Test_impl_GetProgressAllMarvelCharacterIDs(t *testing.T) {
	mCache := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	tests := []struct {
		name    string
		want    model.MarvelAllCharactersCacheProgress
		wantErr bool
		mock    func()
	}{
		{
			name: "not found",
			mock: func() {
				mCache.Delete(progressMarvelCharacterIDs)
			},
		},
		{
			name: "success",
			mock: func() {
				mCache.SetDefault(progressMarvelCharacterIDs, model.MarvelAllCharactersCacheProgress{
					Total: 1000,
				})
			},
			want: model.MarvelAllCharactersCacheProgress{
				Total: 1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				cacheObj: mCache,
			}
			tt.mock()
			got, err := dal.GetProgressAllMarvelCharacterIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetProgressAllMarvelCharacterIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetProgressAllMarvelCharacterIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_SetProgressAllMarvelCharacterIDs(t *testing.T) {
	mCache := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	type args struct {
		progress model.MarvelAllCharactersCacheProgress
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				cacheObj: mCache,
			}
			dal.SetProgressAllMarvelCharacterIDs(tt.args.progress)
		})
	}
}
