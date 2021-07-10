package model

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func TestBuildAPIResponseError(t *testing.T) {
	type args struct {
		statusCode int
		err        error
	}
	tests := []struct {
		name string
		args args
		want GeneralAPIResponse
	}{
		{
			name: "internal server error",
			args: args{
				statusCode: 500,
				err:        errors.New("aaa"),
			},
			want: GeneralAPIResponse{
				Status: 500,
				Error:  defaultInternalServerError,
			},
		},
		{
			name: "timeout",
			args: args{
				statusCode: 500,
				err:        errors.New("net/http: request canceled"),
			},
			want: GeneralAPIResponse{
				Status: 500,
				Error:  "RTO",
			},
		},
		{
			name: "as is",
			args: args{
				statusCode: 500,
				err:        errors.New("currently fetching all marvel characters data"),
			},
			want: GeneralAPIResponse{
				Status: 500,
				Error:  "currently fetching all marvel characters data",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildAPIResponseError(tt.args.statusCode, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildAPIResponseError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildAPIResponseSuccess(t *testing.T) {
	type args struct {
		statusCode int
		data       interface{}
	}
	tests := []struct {
		name string
		args args
		want GeneralAPIResponse
	}{
		{
			name: "success",
			args: args{
				statusCode: 200,
				data:       "OK",
			},
			want: GeneralAPIResponse{
				Status: 200,
				Data:   "OK",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildAPIResponseSuccess(tt.args.statusCode, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildAPIResponseSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}
