package usecase

import (
	"github.com/fgunawan1995/xendit/config"
	marveldal "github.com/fgunawan1995/xendit/dal/api/marvel"
	cachedal "github.com/fgunawan1995/xendit/dal/cache"
	"github.com/fgunawan1995/xendit/model"
)

type impl struct {
	cfg    *config.Config
	marvel marveldal.MarvelDAL
	cache  cachedal.CacheDAL
}

type Usecase interface {
	//GetCharacterByID get single character by id
	GetCharacterByID(id int64) (model.Character, error)
	//GetAllCharacterIDs get array of all character id from cache
	GetAllCharacterIDs() ([]int64, error)
	//SaveCharacters hit marvel character api multiple times and save array of character id to cache
	SaveCharacters() error
}

func New(cfg *config.Config, marvel marveldal.MarvelDAL, cache cachedal.CacheDAL) Usecase {
	return &impl{
		cfg:    cfg,
		marvel: marvel,
		cache:  cache,
	}
}
