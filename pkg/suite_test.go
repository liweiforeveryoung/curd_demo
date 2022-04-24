package pkg

import (
	"context"
	"fmt"
	"github.com/liweiforeveryoung/curd_demo/config"
	"testing"

	gomysql "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SuiteTest struct {
	suite.Suite
	db        *DBEntry
	userEntry *UserEntry
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuiteTest(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}

// SetupAllSuite has a SetupSuite method, which will run before the
// tests in the suite are run.
func (s *SuiteTest) SetupSuite() {
	config.Initialize()
	dsnCfg, err := gomysql.ParseDSN(config.Hub.DBSetting.MysqlDSN)
	if err != nil {
		panic(err)
	}
	dbName := dsnCfg.DBName
	// 在 dsn 中不指定 dbname
	dsnCfg.DBName = ""

	db, err := gorm.Open(mysql.Open(dsnCfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 开启 debug 模式 方便看到每次执行时的 sql 语句
	db.Logger = db.Logger.LogMode(logger.Info)
	drop := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)
	create := fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;", dbName)
	use := fmt.Sprintf(`USE %s;`, dbName)
	err = db.Exec(drop).Exec(create).Exec(use).Error
	if err != nil {
		panic(err)
	}
	s.db = &DBEntry{DB: db}
	err = s.db.Migrate(context.Background())
	s.Require().NoError(err)
	s.userEntry = &UserEntry{DB: db}
}

// TearDownAllSuite has a TearDownSuite method, which will run after
// all the tests in the suite have been run.
func (s *SuiteTest) TearDownSuite() {

}

// SetupTestSuite has a SetupTest method, which will run before each
// test in the suite.
func (s *SuiteTest) SetupTest() {
}

// TearDownTestSuite has a TearDownTest method, which will run after
// each test in the suite.
func (s *SuiteTest) TearDownTest() {

}

// BeforeTest has a function to be executed right before the test
// starts and receives the suite and test names as input
func (s *SuiteTest) BeforeTest(suiteName, testName string) {

}

// AfterTest has a function to be executed right after the test
// finishes and receives the suite and test names as input
func (s *SuiteTest) AfterTest(suiteName, testName string) {

}

// WithStats implements HandleStats, a function that will be executed
// when a test suite is finished. The stats contain information about
// the execution of that suite and its tests.
func (s *SuiteTest) HandleStats(suiteName string, stats *suite.SuiteInformation) {

}
