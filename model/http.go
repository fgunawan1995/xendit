package model

import (
	"log"
	"strings"

	"github.com/pkg/errors"
)

type GeneralAPIResponse struct {
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func BuildAPIResponseError(statusCode int, err error) GeneralAPIResponse {
	errString := defaultInternalServerError
	for k, v := range specialErrors {
		if strings.Contains(err.Error(), k) {
			errString = v
			if v == asIs {
				errString = errors.Cause(err).Error()
			}
		}
	}
	if errString == defaultInternalServerError {
		log.Printf("error = %+v", err)
	}
	return GeneralAPIResponse{
		Status: statusCode,
		Error:  errString,
	}
}

func BuildAPIResponseSuccess(statusCode int, data interface{}) GeneralAPIResponse {
	return GeneralAPIResponse{
		Status: statusCode,
		Data:   data,
	}
}
