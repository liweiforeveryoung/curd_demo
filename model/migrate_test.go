package model

import (
	"reflect"
	"testing"

	"curd_demo/util"
)

func TestMigrationSlice_GroupByFileName(t *testing.T) {
	tests := []struct {
		name  string
		slice MigrationSlice
		want  map[string]*Migration
	}{
		{
			name: "ok",
			slice: MigrationSlice{
				{Id: 1, FileName: "hello"},
				{Id: 2, FileName: "world"},
			},
			want: map[string]*Migration{
				"hello": {Id: 1, FileName: "hello"},
				"world": {Id: 2, FileName: "world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.GroupByFileName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMigrationSlice_FilesNotMigrated(t *testing.T) {
	tests := []struct {
		name   string
		slice  MigrationSlice
		folder *util.Folder
		want   []*util.File
	}{
		{
			name:   "[hello world] && [] = []",
			slice:  MigrationSlice{{FileName: "hello"}, {FileName: "world"}},
			folder: new(util.Folder),
			want:   []*util.File{},
		},
		{
			name:   "[hello world] && [hello] = []",
			slice:  MigrationSlice{{FileName: "hello"}, {FileName: "world"}},
			folder: &util.Folder{Files: []*util.File{{Name: "hello"}}},
			want:   []*util.File{},
		},
		{
			name:   "[hello world] && [world] = []",
			slice:  MigrationSlice{{FileName: "hello"}, {FileName: "world"}},
			folder: &util.Folder{Files: []*util.File{{Name: "world"}}},
			want:   []*util.File{},
		},
		{
			name:   "[hello world] && [XXX] = [XXX]",
			slice:  MigrationSlice{{FileName: "hello"}, {FileName: "world"}},
			folder: &util.Folder{Files: []*util.File{{Name: "XXX"}}},
			want:   []*util.File{{Name: "XXX"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.FilesNotMigrated(tt.folder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilesNotMigrated() = %v, want %v", got, tt.want)
			}
		})
	}
}
