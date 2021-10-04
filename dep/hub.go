package dep

import (
	"curd_demo/config"
	"curd_demo/pkg"
	"go.uber.org/dig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Hub hub

type hub struct {
	dig.In
	DB   *gorm.DB
	User pkg.User
}

var diContainer = dig.New()

func NewGormDB() (*gorm.DB, error) {
	return gorm.Open(mysql.Open(config.Hub.DBSetting.MysqlDSN), &gorm.Config{})
}

func Prepare() {
	_ = diContainer.Provide(NewGormDB)
	_ = diContainer.Provide(pkg.NewUser)

	err := diContainer.Invoke(func(h hub) {
		Hub = h
	})
	if err != nil {
		panic(err)
	}
}
