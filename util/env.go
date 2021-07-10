package util

import (
	"os"

	"github.com/fgunawan1995/xendit/model"
)

func GetEnv() string {
	env := os.Getenv(model.EnvKey)
	if env == "" {
		env = model.LocalEnv
	}
	return env
}
