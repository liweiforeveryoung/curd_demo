package util

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	var randomStr string

	randomStr = RandomString(0)
	assert.Equal(t, "", randomStr)

	randomStr = RandomString(-1)
	assert.Equal(t, "", randomStr)

	// 测 5 次吧
	const randomTimes = 5
	for i := 0; i < randomTimes; i++ {
		length := rand.Intn(100)
		randomStr = RandomString(length)
		assert.Equal(t, length, len(randomStr))
	}
}

func (s *SuiteTest) TestLoadFolderUnderProject() {
	res, err := LoadFolderUnderProject(s.ProjectName, s.RootFolderName)
	s.NoError(err)
	s.Equal(ConcatPath(s.Wd, s.RootFolderName), res.AbsolutePath)

	res, err = LoadFolderUnderProject(s.ProjectName, s.SubFolderName)
	s.NoError(err)
	s.Equal(ConcatPath(s.Wd, s.RootFolderName, s.SubFolderName), res.AbsolutePath)
}

func Test_newFolder(t *testing.T) {
	tests := []struct {
		name                 string
		absolutePathOfFolder string
		want                 *Folder
	}{
		{
			name: "a/b", absolutePathOfFolder: "a/b",
			want: &Folder{AbsolutePath: "a/b", Name: "b"},
		}, {
			name: "a", absolutePathOfFolder: "a",
			want: &Folder{AbsolutePath: "a", Name: "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFolder(tt.absolutePathOfFolder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newFile(t *testing.T) {
	tests := []struct {
		name               string
		absolutePathOfFile string
		want               *File
	}{
		{
			name: "a/b", absolutePathOfFile: "a/b",
			want: &File{AbsolutePath: "a/b", Name: "b"},
		}, {
			name: "a", absolutePathOfFile: "a",
			want: &File{AbsolutePath: "a", Name: "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFile(tt.absolutePathOfFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *SuiteTest) TestFile_apply() {
	// case1: 一个存在的文件
	absolutePathOfFile := ConcatPath(s.Wd, s.RootFolderName, s.File1Name)
	file := newFile(absolutePathOfFile)
	err := file.apply()
	s.NoError(err)

	s.Equal(&File{
		AbsolutePath: absolutePathOfFile,
		Name:         s.File1Name,
		Content:      s.File1Content,
	}, file)

	// case2: 一个不存在的文件
	absolutePathOfFile = ConcatPath(s.Wd, s.RootFolderName, RandomString(32))
	file = newFile(absolutePathOfFile)
	err = file.apply()
	s.Error(err)
}

func (s *SuiteTest) Test_newFolder1() {
	// case1: 一个存在的文件夹
	absolutePathOfFolder := ConcatPath(s.Wd, s.RootFolderName)
	folder := newFolder(absolutePathOfFolder)
	err := folder.apply()
	s.NoError(err)

	expectedFolder := &Folder{
		AbsolutePath: absolutePathOfFolder,
		Name:         s.RootFolderName,
		SubFolders: []*Folder{
			{
				AbsolutePath: ConcatPath(absolutePathOfFolder, s.SubFolderName),
				Name:         s.SubFolderName,
				Files: []*File{
					{
						AbsolutePath: ConcatPath(absolutePathOfFolder, s.SubFolderName, s.File2Name),
						Name:         s.File2Name,
						Content:      s.File2Content,
					},
				},
			},
		},
		Files: []*File{
			{
				AbsolutePath: ConcatPath(absolutePathOfFolder, s.File1Name),
				Name:         s.File1Name,
				Content:      s.File1Content,
			},
		},
	}
	s.Equal(expectedFolder, folder)
	// case2: 一个不存在的文件夹
	absolutePathOfFolder = ConcatPath(s.Wd, RandomString(32))
	folder = newFolder(absolutePathOfFolder)
	err = folder.apply()
	s.Error(err)
}

func TestFolder_AllFiles(t *testing.T) {
	tests := []struct {
		Name   string
		Folder *Folder
		Files  FileSlice
	}{
		{
			// folder1
			//		1
			// 		2
			// 		folder2
			// 			3
			// 			4
			// 			folder4
			// 				7
			// 		folder3
			// 			5
			// 			6
			Name: "case1",
			Folder: &Folder{
				Name: "folder1",
				SubFolders: []*Folder{
					{
						Name:  "folder2",
						Files: []*File{{Name: "3"}, {Name: "4"}},
						SubFolders: []*Folder{{
							Name:  "folder4",
							Files: []*File{{Name: "7"}},
						}},
					},
					{
						Name:  "folder3",
						Files: []*File{{Name: "5"}, {Name: "6"}},
					},
				},
				Files: []*File{{Name: "1"}, {Name: "2"}},
			},
			Files: []*File{{Name: "1"}, {Name: "2"}, {Name: "3"}, {Name: "4"}, {Name: "7"}, {Name: "5"}, {Name: "6"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			m := tt.Folder
			if got := m.AllFiles(); !reflect.DeepEqual(got, tt.Files) {
				t.Errorf("AllFiles() = %v, want %v", got, tt.Files)
			}
		})
	}
}

func Test_absolutePathOfCurProject(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		wd          string
		want        string
		wantErr     bool
	}{
		{
			name:        "a in a/b",
			projectName: "a",
			wd:          "a/b",
			want:        "a",
			wantErr:     false,
		},
		{
			name:        "b in a/b",
			projectName: "b",
			wd:          "a/b",
			want:        "a/b",
			wantErr:     false,
		}, {
			name:        "c in a/b",
			projectName: "c",
			wd:          "a/b",
			want:        "",
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := absolutePathOfCurProject(tt.projectName, tt.wd)
			if (err != nil) != tt.wantErr {
				t.Errorf("absolutePathOfCurProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("absolutePathOfCurProject() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *SuiteTest) TestBindYamlConfig() {
	type NamedObject struct {
		Id   int
		Name string
	}

	type BindYamlTestCfg struct {
		Id          int      `yaml:"id"`
		Name        string   `yaml:"name"`
		IdSlice     []int    `yaml:"id_slice"`
		StringSlice []string `yaml:"name_slice"`
		Object      struct {
			Id   int    `yaml:"id"`
			Name string `yaml:"name"`
		} `yaml:"object"`
		ObjectSlice []struct {
			Id   int    `yaml:"id"`
			Name string `yaml:"name"`
		} `yaml:"object_slice"`
		NamedObject NamedObject `yaml:"named_object"`
	}

	cfgObjPtr := new(BindYamlTestCfg)
	err := BindYamlConfig(s.ConfigFolderName, s.ConfigYamlFileName, cfgObjPtr)
	s.NoError(err)
	dstConfig := &BindYamlTestCfg{
		Id:   1,
		Name: "hello",
		IdSlice: []int{
			1, 2,
		},
		StringSlice: []string{
			"hello",
			"world",
		},
		Object: struct {
			Id   int    `yaml:"id"`
			Name string `yaml:"name"`
		}{
			Id:   1,
			Name: "hello",
		},
		ObjectSlice: []struct {
			Id   int    `yaml:"id"`
			Name string `yaml:"name"`
		}{
			{
				Id:   1,
				Name: "hello",
			},
			{
				Id:   2,
				Name: "world",
			},
		},
		NamedObject: NamedObject{
			Id:   1,
			Name: "hello",
		},
	}
	s.Equal(dstConfig, cfgObjPtr)
}
