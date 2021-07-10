package config

import (
	"testing"
)

func TestGetConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want *Config
	}{
		{
			name: "success",
			args: args{
				path: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetConfig(tt.args.path)
		})
	}
}
