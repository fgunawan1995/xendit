package cache

import (
	"github.com/fgunawan1995/xendit/model"
	"github.com/pkg/errors"
)

//GetAllMarvelCharacterIDs get all marvel character ids array
func (dal *impl) GetAllMarvelCharacterIDs() ([]int64, error) {
	var result []int64
	temp, found := dal.cacheObj.Get(marvelCharacterIDs)
	if !found {
		return result, nil
	}
	result, ok := temp.([]int64)
	if !ok {
		return result, errors.WithStack(errors.New("failed conversion"))
	}
	return result, nil
}

//SaveAllMarvelCharacterIDs save all marvel character ids array
func (dal *impl) SaveAllMarvelCharacterIDs(characterIDs []int64) {
	dal.cacheObj.SetDefault(marvelCharacterIDs, characterIDs)
}

//GetProgressAllMarvelCharacterIDs get all marvel character progress cache
func (dal *impl) GetProgressAllMarvelCharacterIDs() (model.MarvelAllCharactersCacheProgress, error) {
	var result model.MarvelAllCharactersCacheProgress
	temp, found := dal.cacheObj.Get(progressMarvelCharacterIDs)
	if !found {
		return result, nil
	}
	result, ok := temp.(model.MarvelAllCharactersCacheProgress)
	if !ok {
		return result, errors.WithStack(errors.New("failed conversion"))
	}
	return result, nil
}

//SetProgressAllMarvelCharacterIDs set cache progress for fetching all marvel characters
func (dal *impl) SetProgressAllMarvelCharacterIDs(progress model.MarvelAllCharactersCacheProgress) {
	dal.cacheObj.SetDefault(progressMarvelCharacterIDs, progress)
}
