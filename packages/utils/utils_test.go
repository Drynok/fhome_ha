// Package utils_test
package utils_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type utilsTestSuite struct {
	suite.Suite
}

func (s *utilsTestSuite) SetupSuite() {
}

func (s *utilsTestSuite) SetupTest() {
}

func (s *utilsTestSuite) TestNew() {
}
func TestEnv(t *testing.T) {
	suite.Run(t, new(utilsTestSuite))
}
