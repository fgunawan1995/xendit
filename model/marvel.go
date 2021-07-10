package model

const (
	MarvelDefaultLimit = 100
)

type MarvelAllCharactersCacheProgress struct {
	Total   int `json:"total"`
	Current int `json:"current"`
}

type MarvelGetCharactersResponse struct {
	Code   int                             `json:"code"`
	Status string                          `json:"status"`
	Data   MarvelGetCharactersResponseData `json:"data"`
}

type MarvelGetCharactersResponseData struct {
	Offset  int                                 `json:"offset"`
	Limit   int                                 `json:"limit"`
	Total   int                                 `json:"total"`
	Count   int                                 `json:"count"`
	Results []MarvelGetCharactersResponseResult `json:"results"`
}

type MarvelGetCharactersResponseResult struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Modified    string `json:"modified"`
	ResourceURI string `json:"resourceURI"`
}

type MarvelGetCharacterRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func (p MarvelGetCharacterRequest) NextPage() MarvelGetCharacterRequest {
	return MarvelGetCharacterRequest{
		Limit:  p.Limit,
		Offset: p.Offset + p.Limit,
	}
}

func (p MarvelGetCharactersResponseResult) ToCharacter() Character {
	return Character{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
	}
}

type Character struct {
	ID          int64  `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}
