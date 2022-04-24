package model

import "github.com/liweiforeveryoung/curd_demo/util"

type MigrationSlice []*Migration

func (slice MigrationSlice) GroupByFileName() map[string]*Migration {
	groupByFileName := make(map[string]*Migration)
	for _, migration := range slice {
		groupByFileName[migration.FileName] = migration
	}
	return groupByFileName
}

// FilesNotMigrated 返回还没有被 migrated 的文件
func (slice MigrationSlice) FilesNotMigrated(folder *util.Folder) []*util.File {
	groupByFileName := slice.GroupByFileName()
	unMigratedFiles := make([]*util.File, 0, 0)
	for _, file := range folder.AllFiles() {
		if _, exist := groupByFileName[file.Name]; !exist {
			// 说明之前该文件没被 migrate
			unMigratedFiles = append(unMigratedFiles, file)
		}
	}
	return unMigratedFiles
}

type Migration struct {
	Id int64 `gorm:"column:id" json:"-"`
	// 已经迁移的 sql 文件名
	FileName  string `gorm:"column:file_name" json:"file_name"`
	DeletedAt int64  `gorm:"column:deleted_at" json:"-"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"-"`
}

func (m *Migration) TableName() string {
	return "migrations"
}

func (m *Migration) SQL() string {
	return `CREATE TABLE IF NOT EXISTS migrations
(
    id         BIGINT(20)   NOT NULL AUTO_INCREMENT,
    file_name  VARCHAR(191) NOT NULL DEFAULT '' COMMENT '已经迁移的文件名',
    deleted_at BIGINT(20)   NOT NULL DEFAULT 0,
    created_at BIGINT(20)   NOT NULL DEFAULT 0,
    updated_at BIGINT(20)   NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    UNIQUE INDEX udx_file_name (file_name)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT '迁移信息表';
	`
}
