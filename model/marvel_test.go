package model

import (
	"reflect"
	"testing"
)

func TestMarvelGetCharacterRequest_NextPage(t *testing.T) {
	type fields struct {
		Limit  int
		Offset int
	}
	tests := []struct {
		name   string
		fields fields
		want   MarvelGetCharacterRequest
	}{
		{
			name: "success",
			fields: fields{
				Limit: 10,
			},
			want: MarvelGetCharacterRequest{
				Offset: 10,
				Limit:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MarvelGetCharacterRequest{
				Limit:  tt.fields.Limit,
				Offset: tt.fields.Offset,
			}
			if got := p.NextPage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarvelGetCharacterRequest.NextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarvelGetCharactersResponseResult_ToCharacter(t *testing.T) {
	type fields struct {
		ID          int64
		Name        string
		Description string
		Modified    string
		ResourceURI string
	}
	tests := []struct {
		name   string
		fields fields
		want   Character
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := MarvelGetCharactersResponseResult{
				ID:          tt.fields.ID,
				Name:        tt.fields.Name,
				Description: tt.fields.Description,
				Modified:    tt.fields.Modified,
				ResourceURI: tt.fields.ResourceURI,
			}
			if got := p.ToCharacter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarvelGetCharactersResponseResult.ToCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
