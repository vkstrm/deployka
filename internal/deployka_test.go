package internal

import (
	"github.com/stretchr/testify/suite"
	"os"
	"strings"
	"testing"
)

// Suite

type DeploykaTestSuite struct {
	suite.Suite
	tmpDir string
}

func (suite *DeploykaTestSuite) SetupSuite() {
	tmpDir := "/tmp/deployka/config"
	_ = os.Setenv(ConfigDirEnv, tmpDir)

	// Create a temporary directory
	err := os.MkdirAll(tmpDir, os.ModePerm)
	suite.Require().NoError(err)

	suite.tmpDir = "/tmp/deployka"
}

func (suite *DeploykaTestSuite) SetupTest() {
}

func (suite *DeploykaTestSuite) TearDownSuite() {
	err := os.RemoveAll(suite.tmpDir)
	suite.Require().NoError(err)
}

// Required in order to run the tests
func TestDeployka(t *testing.T) {
	suite.Run(t, new(DeploykaTestSuite))
}

// Tests

func (suite *DeploykaTestSuite) TestConfig() {
	inputLines := []string{
		"TV5742zYaBcfKJgo5LQRojsfeSiOEj8643563", // API key
		"https://api.example.com",
	}

	input := strings.Join(inputLines, "\n")

	r := strings.NewReader(input)
	err := Config(r)
	suite.NoError(err)
}

