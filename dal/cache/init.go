package cache

import (
	"github.com/fgunawan1995/xendit/model"
	cache "github.com/patrickmn/go-cache"
)

type CacheDAL interface {
	//GetAllMarvelCharacterIDs get all marvel character ids array
	GetAllMarvelCharacterIDs() ([]int64, error)
	//SaveAllMarvelCharacterIDs save all marvel character ids array
	SaveAllMarvelCharacterIDs(characterIDs []int64)
	//GetProgressAllMarvelCharacterIDs get all marvel character progress cache
	GetProgressAllMarvelCharacterIDs() (model.MarvelAllCharactersCacheProgress, error)
	//SetProgressAllMarvelCharacterIDs set cache progress for fetching all marvel characters
	SetProgressAllMarvelCharacterIDs(progress model.MarvelAllCharactersCacheProgress)
}

type impl struct {
	cacheObj *cache.Cache
}

func New(cache *cache.Cache) CacheDAL {
	return &impl{
		cacheObj: cache,
	}
}
