package resources

import (
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/fgunawan1995/xendit/config"
)

func TestInitMarvelClient(t *testing.T) {
	type args struct {
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want *http.Client
	}{
		{
			name: "success",
			args: args{
				cfg: &config.Config{
					Marvel: config.MarvelAPIConfig{
						TimeoutInMilliSecond: 1000,
					},
				},
			},
			want: &http.Client{
				Timeout: 1 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitMarvelClient(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitMarvelClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
