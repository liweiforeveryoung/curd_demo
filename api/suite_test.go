package api

import (
	"github.com/liweiforeveryoung/curd_demo/config"
	"github.com/liweiforeveryoung/curd_demo/dep"
	"github.com/liweiforeveryoung/curd_demo/pkg"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite

	userMock *pkg.MockUser
	ctrl     *gomock.Controller
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
}

// TearDownAllSuite has a TearDownSuite method, which will run after
// all the tests in the suite have been run.
func (s *SuiteTest) TearDownSuite() {

}

// SetupTestSuite has a SetupTest method, which will run before each
// test in the suite.
func (s *SuiteTest) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.userMock = pkg.NewMockUser(s.ctrl)
	dep.Hub.User = s.userMock
}

// TearDownTestSuite has a TearDownTest method, which will run after
// each test in the suite.
func (s *SuiteTest) TearDownTest() {
	s.ctrl.Finish()
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
