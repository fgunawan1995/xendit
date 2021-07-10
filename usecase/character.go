package usecase

import (
	"fmt"
	"log"

	"github.com/fgunawan1995/xendit/model"
	"github.com/pkg/errors"
)

//GetCharacterByID get single character by id
func (u *impl) GetCharacterByID(id int64) (model.Character, error) {
	result, err := u.marvel.GetCharacterByID(id)
	if err != nil {
		return result.ToCharacter(), errors.WithStack(err)
	}
	return result.ToCharacter(), nil
}

//GetAllCharacterIDs get array of all character id from cache
func (u *impl) GetAllCharacterIDs() ([]int64, error) {
	var result []int64
	progress, err := u.cache.GetProgressAllMarvelCharacterIDs()
	if err != nil {
		return result, errors.WithStack(err)
	}
	if progress.Total == 0 || progress.Current < progress.Total {
		return result, errors.WithStack(fmt.Errorf("currently fetching all marvel characters data, progress=%d/%d", progress.Current, progress.Total))
	}
	result, err = u.cache.GetAllMarvelCharacterIDs()
	if err != nil {
		return result, errors.WithStack(err)
	}
	return result, nil

}

//SaveCharacters hit marvel character api multiple times and save array of character id to cache
func (u *impl) SaveCharacters() error {
	var characterIDs []int64
	log.Print("started fetching all marvel character data")
	progress := model.MarvelAllCharactersCacheProgress{}
	u.cache.SetProgressAllMarvelCharacterIDs(progress)
	// init param for first page
	param := model.MarvelGetCharacterRequest{
		Limit: model.MarvelDefaultLimit,
	}
	for {
		// get all character paginated
		result, err := u.marvel.GetCharacters(param)
		if err != nil {
			return errors.WithStack(err)
		}
		current := result.Offset + result.Count
		total := result.Total
		log.Printf("fetching all marvel character data, progress = %d/%d", current, total)
		// save to array
		for _, data := range result.Results {
			characterIDs = append(characterIDs, data.ID)
		}
		progress.Current = current
		progress.Total = total
		u.cache.SetProgressAllMarvelCharacterIDs(progress)
		// stop when last page
		if current >= total {
			break
		}
		// next page
		param = param.NextPage()
	}
	u.cache.SaveAllMarvelCharacterIDs(characterIDs)
	log.Print("finished fetching all marvel character data")
	return nil
}
