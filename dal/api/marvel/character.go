package marvel

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fgunawan1995/xendit/model"
	"github.com/pkg/errors"
)

//generateAuthParamGet helper func to generate auth param
func (c *impl) generateAuthParamGet() string {
	ts := time.Now().UTC().Unix()
	stringToHash := fmt.Sprintf("%d%s%s", ts, c.cfg.Marvel.PrivateKey, c.cfg.Marvel.PublicKey)
	hash := md5.Sum([]byte(stringToHash))
	return fmt.Sprintf("apikey=%s&ts=%d&hash=%s", c.cfg.Marvel.PublicKey, ts, hex.EncodeToString(hash[:]))
}

//GetCharacterByID from marvel api and return single data of character
func (c *impl) GetCharacterByID(id int64) (model.MarvelGetCharactersResponseResult, error) {
	var result model.MarvelGetCharactersResponseResult
	var apiResp model.MarvelGetCharactersResponse
	url := fmt.Sprintf("%s%s/%d?%s", c.cfg.Marvel.Host, c.cfg.Marvel.EndpointGetCharacters, id, c.generateAuthParamGet())
	resp, err := c.client.Get(url)
	if err != nil {
		return result, errors.WithStack(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, errors.WithStack(err)
	}
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return result, errors.WithStack(err)
	}
	if apiResp.Code != http.StatusOK {
		return result, errors.WithStack(fmt.Errorf("marvelStatus:%d,error:%s", apiResp.Code, apiResp.Status))
	}
	result = apiResp.Data.Results[0]
	return result, nil
}

//GetCharacters get all character from marvel api (paginated)
func (c *impl) GetCharacters(param model.MarvelGetCharacterRequest) (model.MarvelGetCharactersResponseData, error) {
	var result model.MarvelGetCharactersResponseData
	var apiResp model.MarvelGetCharactersResponse
	url := fmt.Sprintf("%s%s?limit=%d&offset=%d&%s", c.cfg.Marvel.Host, c.cfg.Marvel.EndpointGetCharacters, param.Limit, param.Offset, c.generateAuthParamGet())
	resp, err := c.client.Get(url)
	if err != nil {
		return result, errors.WithStack(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, errors.WithStack(err)
	}
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		return result, errors.WithStack(err)
	}
	if apiResp.Code != http.StatusOK {
		return result, errors.WithStack(fmt.Errorf("marvelStatus:%d,error:%s", apiResp.Code, apiResp.Status))
	}
	result = apiResp.Data
	return result, nil
}
