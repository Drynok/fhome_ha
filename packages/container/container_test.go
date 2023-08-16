package container_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type conainerTestSuite struct {
	suite.Suite
}

func (s *conainerTestSuite) SetupSuite() {
}

func (s *conainerTestSuite) SetupTest() {
}

func (s *conainerTestSuite) TestNew() {
}
func TestEnv(t *testing.T) {
	suite.Run(t, new(conainerTestSuite))
}
