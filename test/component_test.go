package test

import (
	"testing"

	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
	"github.com/stretchr/testify/assert"
)

type ComponentSuite struct {
	helper.TestSuite
}

func (s *ComponentSuite) TestBasic() {
	const component = "tgw-spoke/basic"
	const stack = "default-ue2-test"
	const awsRegion = "us-east-2"

	defer s.DestroyAtmosComponent(s.T(), component, stack, nil)
	options, _ := s.DeployAtmosComponent(s.T(), component, stack, nil)
	assert.NotNil(s.T(), options)

    s.DriftTest(component, stack, nil)
}

func (s *ComponentSuite) TestEnabledFlag() {
	const component = "tgw-spoke/disabled"
	const stack = "default-ue2-test"
	s.VerifyEnabledFlag(component, stack, nil)
}

func (s *ComponentSuite) SetupSuite() {
	s.TestSuite.InitConfig()
	s.TestSuite.Config.ComponentDestDir = "components/terraform/tgw/spoke"
	s.TestSuite.SetupSuite()
  }

func TestRunSuite(t *testing.T) {
	suite := new(ComponentSuite)

	suite.AddDependency(t, "vpc", "default-ue2-test", nil)
	suite.AddDependency(t, "tgw/hub", "default-ue2-test", nil)

	suite.AddDependency(t, "vpc", "default-ue1-test", nil)
	suite.AddDependency(t, "tgw/hub", "default-ue1-test", nil)
	helper.Run(t, suite)
}
