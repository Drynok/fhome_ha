package env_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type envTestSuite struct {
	suite.Suite
}

func (s *envTestSuite) SetupSuite() {
}

func (s *envTestSuite) SetupTest() {
}

func (s *envTestSuite) TestNew() {
}

func TestEnv(t *testing.T) {
	suite.Run(t, new(envTestSuite))
}
