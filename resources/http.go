package resources

import (
	"net/http"
	"time"

	"github.com/fgunawan1995/xendit/config"
)

func InitMarvelClient(cfg *config.Config) *http.Client {
	return &http.Client{
		Timeout: time.Duration(cfg.Marvel.TimeoutInMilliSecond) * time.Millisecond,
	}
}
