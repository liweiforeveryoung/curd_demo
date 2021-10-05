package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite
	RootFolderName        string
	Wd                    string
	ProjectName           string
	SubFolderName         string
	File1Name             string
	File1Content          []byte
	File2Name             string
	File2Content          []byte
	ConfigFolderName      string
	ConfigYamlFileName    string
	ConfigYamlFileContent []byte
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSuiteTest(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}

// SetupAllSuite has a SetupSuite method, which will run before the
// tests in the suite are run.
func (s *SuiteTest) SetupSuite() {
	// util/test_for_util_pkg
	// ├── dir1
	// │	└── text2.txt [world]
	// └── text1.txt [hello]
	s.RootFolderName = "test_for_util_pkg"
	s.ProjectName = "util"
	s.SubFolderName = "dir1"
	s.File1Name = "text1.txt"
	s.File1Content = []byte("hello")
	s.File2Name = "text2.txt"
	s.File2Content = []byte("world")

	err := os.Mkdir(s.RootFolderName, os.ModeDir|os.ModePerm)
	s.NoError(err)
	file1, err := os.Create(ConcatPath(".", s.RootFolderName, s.File1Name))
	s.NoError(err)
	defer file1.Close()
	_, err = file1.Write(s.File1Content)
	s.NoError(err)

	err = os.Mkdir(ConcatPath(".", s.RootFolderName, s.SubFolderName), os.ModeDir|os.ModePerm)
	s.NoError(err)
	file2, err := os.Create(ConcatPath(".", s.RootFolderName, s.SubFolderName, s.File2Name))
	s.NoError(err)
	defer file2.Close()
	_, err = file2.Write(s.File2Content)
	s.NoError(err)

	s.Wd, err = os.Getwd()
	s.NoError(err)

	s.ConfigFolderName = "bind_config_folder"
	s.ConfigYamlFileName = "bind_config_test.yaml"
	s.ConfigYamlFileContent = []byte(
		`id: 1
name: "hello"
id_slice:
  - 1
  - 2
name_slice:
  - "hello"
  - "world"
object:
  id: 1
  name: "hello"
object_slice:
  - id: 1
    name: "hello"
  - id: 2
    name: "world"
named_object:
  id: 1
  name: "hello"`)
	err = os.Mkdir(s.ConfigFolderName, os.ModeDir|os.ModePerm)
	s.NoError(err)
	yamlFile, err := os.Create(ConcatPath(".", s.ConfigFolderName, s.ConfigYamlFileName))
	s.NoError(err)
	defer yamlFile.Close()
	_, err = yamlFile.Write(s.ConfigYamlFileContent)
	s.NoError(err)
}

// TearDownAllSuite has a TearDownSuite method, which will run after
// all the tests in the suite have been run.
func (s *SuiteTest) TearDownSuite() {
	err := os.RemoveAll(s.RootFolderName)
	s.NoError(err)
	err = os.RemoveAll(s.ConfigFolderName)
	s.NoError(err)
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
