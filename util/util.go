package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func RandomString(length int) string {
	if length <= 0 {
		return ""
	}
	set := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	setCap := len(set)

	res := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		index := rand.Intn(setCap)
		res = append(res, set[index])
	}
	return string(res)
}

// File 表示 文件
type File struct {
	// 完整路径(包含 Name)
	AbsolutePath string
	// 文件名字(包含后缀)
	Name string
	// 内容
	Content []byte
}

// Folder 表示文件夹
type Folder struct {
	// 完整路径(包含 Name)
	AbsolutePath string
	// 文件夹的名字
	Name string
	// 文件夹下面的子文件夹
	SubFolders []*Folder
	// 文件夹下面的文件
	Files []*File
}

type FileSlice []*File

func (slice FileSlice) Names() []string {
	names := make([]string, 0, len(slice))
	for _, file := range slice {
		names = append(names, file.Name)
	}
	return names
}

func (slice FileSlice) Contents() []string {
	contents := make([]string, 0, len(slice))
	for _, file := range slice {
		contents = append(contents, string(file.Content))
	}
	return contents
}

// AllFiles 返回文件夹下面的所有文件
func (m *Folder) AllFiles() FileSlice {
	files := make(FileSlice, 0, 0)
	files = append(files, m.Files...)
	for _, folder := range m.SubFolders {
		files = append(files, folder.AllFiles()...)
	}
	return files
}

// FindFile 根据文件名查找该文件夹下面的问题, 如果没有找到将会返回 nil
func (m *Folder) FindFile(fileName string) *File {
	for _, file := range m.Files {
		if file.Name == fileName {
			return file
		}
	}
	for _, folder := range m.SubFolders {
		if file := folder.FindFile(fileName); file != nil {
			return file
		}
	}
	return nil
}

// FindFolder 根据文件夹名查找文件夹, 如果没有找到将会返回 nil
func (m *Folder) FindFolder(folderName string) *Folder {
	if m.Name == folderName {
		return m
	}
	for _, subFolder := range m.SubFolders {
		if folder := subFolder.FindFolder(folderName); folder != nil {
			return folder
		}
	}
	return nil
}

func newFolder(absolutePathOfFolder string) *Folder {
	folder := new(Folder)
	folder.AbsolutePath = absolutePathOfFolder
	idx := strings.LastIndex(absolutePathOfFolder, string(os.PathSeparator))
	folder.Name = absolutePathOfFolder[idx+1:]
	return folder
}

func newFile(absolutePathOfFile string) *File {
	file := new(File)
	file.AbsolutePath = absolutePathOfFile
	idx := strings.LastIndex(absolutePathOfFile, string(os.PathSeparator))
	file.Name = absolutePathOfFile[idx+1:]
	return file
}

func (m *Folder) apply() error {
	entries, err := os.ReadDir(m.AbsolutePath)
	if err != nil {
		return errors.WithStack(err)
	}
	for _, entry := range entries {
		absolutePath := ConcatPath(m.AbsolutePath, entry.Name())
		if entry.IsDir() {
			subFolder := newFolder(absolutePath)
			if err := subFolder.apply(); err != nil {
				return errors.WithStack(err)
			}
			m.SubFolders = append(m.SubFolders, subFolder)
		} else {
			file := newFile(absolutePath)
			if err := file.apply(); err != nil {
				return errors.WithStack(err)
			}
			m.Files = append(m.Files, file)
		}
	}
	return nil
}

func (m *File) apply() error {
	content, err := os.ReadFile(m.AbsolutePath)
	if err != nil {
		return errors.WithStack(err)
	}
	m.Content = content
	return nil
}

// LoadFolder 根据文件夹的绝对路径 load 出一个文件夹
// 如果文件夹不存在, 将会返回一个 error
func LoadFolder(absolutePathOfFolder string) (*Folder, error) {
	folder := newFolder(absolutePathOfFolder)
	err := folder.apply()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return folder, nil
}

// AbsolutePathOfCurProject 获取当前项目的绝对路径
// projectName 必须在当前的 working directory 下面
func AbsolutePathOfCurProject(projectName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return absolutePathOfCurProject(projectName, wd)
}

func absolutePathOfCurProject(projectName, wd string) (string, error) {
	idx := strings.LastIndex(wd, projectName)
	if idx == -1 {
		return "", errors.Errorf("can't find project name[%s] in wd[%s]", projectName, wd)
	}
	return wd[:idx+len(projectName)], nil
}

// LoadFolderUnderProject 从项目根路径开始寻找 folder, 返回 folder 对象, 如果找不到, nil 将非空
func LoadFolderUnderProject(projectName, folderName string) (*Folder, error) {
	absolutePath, err := AbsolutePathOfCurProject(projectName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	folder, err := LoadFolder(absolutePath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if dst := folder.FindFolder(folderName); dst != nil {
		return dst, nil
	}

	return nil, errors.New("can't find")
}

// BindYamlConfig 负责根据 yaml config file unmarshal 出 struct
func BindYamlConfig(cfgBaseDir, cfgFileName string, cfgObjPtr interface{}) error {
	vp := viper.New()
	// jww.SetStdoutThreshold(jww.LevelInfo) 开启 viper 的日志
	vp.AddConfigPath(cfgBaseDir)
	cfgFileName = strings.TrimSuffix(cfgFileName, ".yaml")
	vp.SetConfigName(cfgFileName)
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return fmt.Errorf("ReadInConfig(),err[%w]", err)
	}
	if err = vp.Unmarshal(cfgObjPtr, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		// 不能多出不用的配置项
		config.ErrorUnused = true
	}); err != nil {
		return fmt.Errorf("unmarshal(),err[%w]", err)
	}
	return nil
}

func BindResp(resp *http.Response, obj interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ReadAll err[%w]", err)
	}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return fmt.Errorf("unmarshal err[%w]", err)
	}
	return nil
}

func ConcatPath(paths ...string) string {
	return strings.Join(paths, string(os.PathSeparator))
}
