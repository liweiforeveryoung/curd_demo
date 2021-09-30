package dep

import (
	"curd_demo/config"
	"curd_demo/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Hub struct {
	DB   *gorm.DB
	User pkg.User
}

func Prepare() {
	db, err := gorm.Open(mysql.Open(config.Hub.DBSetting.MysqlDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Hub.DB = db
	Hub.User = pkg.NewUser(db)
}
