package usecase

import (
	"reflect"
	"testing"

	"github.com/fgunawan1995/xendit/mocks"
	"github.com/fgunawan1995/xendit/model"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

func Test_impl_GetCharacterByID(t *testing.T) {
	marvelDAL := new(mocks.MarvelDAL)
	cacheDAL := new(mocks.CacheDAL)
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		args    args
		want    model.Character
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				id: 1,
			},
			mock: func() {
				marvelDAL.On("GetCharacterByID", mock.Anything).Return(model.MarvelGetCharactersResponseResult{}, nil).Once()
			},
		},
		{
			name: "error",
			args: args{
				id: 1,
			},
			mock: func() {
				marvelDAL.On("GetCharacterByID", mock.Anything).Return(model.MarvelGetCharactersResponseResult{}, errors.New("aaa")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				marvel: marvelDAL,
				cache:  cacheDAL,
			}
			tt.mock()
			got, err := u.GetCharacterByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetCharacterByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetCharacterByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_GetAllCharacterIDs(t *testing.T) {
	marvelDAL := new(mocks.MarvelDAL)
	cacheDAL := new(mocks.CacheDAL)
	tests := []struct {
		name    string
		want    []int64
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				cacheDAL.On("GetProgressAllMarvelCharacterIDs").Return(model.MarvelAllCharactersCacheProgress{
					Total:   1,
					Current: 1,
				}, nil).Once()
				cacheDAL.On("GetAllMarvelCharacterIDs").Return([]int64{1}, nil).Once()
			},
			want: []int64{1},
		},
		{
			name: "in progress",
			mock: func() {
				cacheDAL.On("GetProgressAllMarvelCharacterIDs").Return(model.MarvelAllCharactersCacheProgress{
					Total:   1,
					Current: 0,
				}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "error progress",
			mock: func() {
				cacheDAL.On("GetProgressAllMarvelCharacterIDs").Return(model.MarvelAllCharactersCacheProgress{}, errors.New("aaa")).Once()
			},
			wantErr: true,
		},
		{
			name: "error data",
			mock: func() {
				cacheDAL.On("GetProgressAllMarvelCharacterIDs").Return(model.MarvelAllCharactersCacheProgress{
					Total:   1,
					Current: 1,
				}, nil).Once()
				cacheDAL.On("GetAllMarvelCharacterIDs").Return([]int64{}, errors.New("aaa")).Once()
			},
			want:    []int64{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				marvel: marvelDAL,
				cache:  cacheDAL,
			}
			tt.mock()
			got, err := u.GetAllCharacterIDs()
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetAllCharacterIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetAllCharacterIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_SaveCharacters(t *testing.T) {
	marvelDAL := new(mocks.MarvelDAL)
	cacheDAL := new(mocks.CacheDAL)
	tests := []struct {
		name    string
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				cacheDAL.On("SetProgressAllMarvelCharacterIDs", mock.Anything)
				marvelDAL.On("GetCharacters", mock.Anything).Return(model.MarvelGetCharactersResponseData{
					Total: 10,
					Count: 5,
				}, nil).Once()
				cacheDAL.On("SetProgressAllMarvelCharacterIDs", mock.Anything)
				marvelDAL.On("GetCharacters", mock.Anything).Return(model.MarvelGetCharactersResponseData{
					Total:  10,
					Count:  5,
					Offset: 5,
				}, nil).Once()
				cacheDAL.On("SetProgressAllMarvelCharacterIDs", mock.Anything)
				cacheDAL.On("SaveAllMarvelCharacterIDs", mock.Anything)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				marvel: marvelDAL,
				cache:  cacheDAL,
			}
			tt.mock()
			if err := u.SaveCharacters(); (err != nil) != tt.wantErr {
				t.Errorf("impl.SaveCharacters() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
