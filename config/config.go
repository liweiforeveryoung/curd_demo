package config

import (
	"fmt"
	"os"

	"curd_demo/util"
)

const (
	ProjectName          = "curd_demo"
	CfgFolderName        = "config"
	MigrationsFolderName = "migrations"
)

var Hub struct {
	HttpSetting *HttpSetting
	DBSetting   *DBSetting
}

func Initialize() {
	env := os.Getenv("APP_ENVIRONMENT")
	Hub.HttpSetting = httpConfigInit(env)
	Hub.DBSetting = dbConfigInit(env)
}

var CfgAbsolutePath = func() string {
	dir, err := util.DirOrFileAbsolutePathFromProject(ProjectName, CfgFolderName)
	if err != nil {
		panic(err)
	}
	return dir
}()

type DBSetting struct {
	MysqlDSN string `yaml:"mysql_dsn"`
}
type HttpSetting struct {
	Addr string `yaml:"addr"`
}

func dbConfigInit(env string) *DBSetting {
	cfg := new(DBSetting)
	var err error
	switch env {
	default:
		err = util.BindYamlConfig(CfgAbsolutePath, "db.test", cfg)
	case "TEST":
		err = util.BindYamlConfig(CfgAbsolutePath, "db.test", cfg)
	case "PRODUCTION":
		err = util.BindYamlConfig(CfgAbsolutePath, "db.production", cfg)
	}
	if err != nil {
		panic(fmt.Errorf("BindYamlConfig(),err[%w]", err))
	}
	return cfg
}

func httpConfigInit(env string) *HttpSetting {
	cfg := new(HttpSetting)
	var err error
	switch env {
	default:
		err = util.BindYamlConfig(CfgAbsolutePath, "http.test", cfg)
	case "TEST":
		err = util.BindYamlConfig(CfgAbsolutePath, "http.test", cfg)
	case "PRODUCTION":
		err = util.BindYamlConfig(CfgAbsolutePath, "http.production", cfg)
	}
	if err != nil {
		panic(fmt.Errorf("BindYamlConfig(),err[%w]", err))
	}
	return cfg
}
