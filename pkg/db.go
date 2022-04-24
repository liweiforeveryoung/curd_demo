package pkg

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/dig"
	"gorm.io/gorm"

	"github.com/liweiforeveryoung/curd_demo/config"
	"github.com/liweiforeveryoung/curd_demo/model"
	"github.com/liweiforeveryoung/curd_demo/util"
)

type DB interface {
	Migrate(ctx context.Context) error
}

func NewDB(entry DBEntry) DB {
	return &entry
}

type DBEntry struct {
	dig.In
	*gorm.DB
}

func (entry *DBEntry) Migrate(ctx context.Context) error {
	err := entry.WithContext(ctx).Exec(new(model.Migration).SQL()).Error
	if err != nil {
		return errors.WithStack(err)
	}
	migrations := make([]*model.Migration, 0, 0)
	err = entry.WithContext(ctx).Find(&migrations).Error
	if err != nil {
		return errors.WithStack(err)
	}
	folder, err := util.LoadFolderUnderProject(config.ProjectName, config.MigrationsFolderName)
	if err != nil {
		return errors.WithStack(err)
	}
	filesNotMigrated := model.MigrationSlice(migrations).FilesNotMigrated(folder)
	for _, file := range filesNotMigrated {
		// 执行迁移文件
		if err = entry.WithContext(ctx).Exec(string(file.Content)).Error; err != nil {
			return errors.Errorf("exec err,file[%v],err[%v]", file, err)
		}
		// 插入 migrations 表中一条记录
		m := &model.Migration{FileName: file.Name}
		if err = entry.WithContext(ctx).Create(m).Error; err != nil {
			return errors.Errorf("create err,migration[%v],err[%v]", m, err)
		}
	}
	return nil
}
