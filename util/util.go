package util

import (
	"fmt"
	"math/rand"
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

// DirOrFileAbsolutePathFromProject 从项目根路径开始寻找 dir/file, 返回完整的 path
// file 如果有后缀, 需要带上
func DirOrFileAbsolutePathFromProject(projectName, dirOrFileName string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", errors.WithStack(err)
	}
	// 在 learn_ci 下面找到 config 目录
	index := strings.LastIndex(wd, projectName)
	rootPath := wd[:index+len(projectName)]
	path, err := FindDirOrFileNameFullPath([]string{rootPath}, dirOrFileName)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return path, nil
}

// FindDirOrFileNameFullPath 从 parentPaths 开始往下面查询查找 dir/file 的完整路径名
// 如果没有找到, 将会返回 error
func FindDirOrFileNameFullPath(parentPaths []string, dirOrFileName string) (string, error) {
	if len(parentPaths) == 0 {
		return "", errors.Errorf("can't find dirOrFileName: %s", dirOrFileName)
	}
	nextParentPaths := make([]string, 0, 0)
	for _, parentPath := range parentPaths {
		entries, err := os.ReadDir(parentPath)
		if err != nil {
			return "", errors.WithStack(err)
		}
		for _, entry := range entries {
			path := strings.Join([]string{parentPath, entry.Name()}, string(os.PathSeparator))
			if entry.Name() == dirOrFileName {
				return path, nil
			}
			if entry.IsDir() {
				nextParentPaths = append(nextParentPaths, path)
			}
		}
	}
	return FindDirOrFileNameFullPath(nextParentPaths, dirOrFileName)
}

// FileNamesInDir 返回 dir 下面所有的文件名字 (不包括路径)
func FileNamesInDir(dirPath string) ([]string, error) {
	dirs := []string{dirPath}
	fileNames := make([]string, 0, 0)
	for len(dirs) != 0 {
		dir := dirs[0]
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, entry := range entries {
			path := strings.Join([]string{dir, entry.Name()}, string(os.PathSeparator))
			if entry.IsDir() {
				dirs = append(dirs, path)
			} else {
				fileNames = append(fileNames, entry.Name())
			}
		}
		dirs = dirs[1:]
	}
	return fileNames, nil
}

// BindYamlConfig 负责根据 yaml config file unmarshal 出 struct
// cfgFileName shouldn't contain file type suffix
func BindYamlConfig(cfgBaseDir, cfgFileName string, cfgObjPtr interface{}) error {
	vp := viper.New()
	// jww.SetStdoutThreshold(jww.LevelInfo) 开启 viper 的日志
	vp.AddConfigPath(cfgBaseDir)
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
