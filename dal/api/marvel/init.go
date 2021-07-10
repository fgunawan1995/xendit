package marvel

import (
	"net/http"

	"github.com/fgunawan1995/xendit/config"
	"github.com/fgunawan1995/xendit/model"
)

type MarvelDAL interface {
	//GetCharacterByID from marvel api and return single data of character
	GetCharacterByID(id int64) (model.MarvelGetCharactersResponseResult, error)
	//GetCharacters get all character from marvel api (paginated)
	GetCharacters(param model.MarvelGetCharacterRequest) (model.MarvelGetCharactersResponseData, error)
}

type impl struct {
	client *http.Client
	cfg    *config.Config
}

func New(cfg *config.Config, httpClient *http.Client) MarvelDAL {
	return &impl{
		cfg:    cfg,
		client: httpClient,
	}
}
