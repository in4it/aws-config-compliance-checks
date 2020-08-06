package mocks

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/configservice"
)

type MockAWSConfigService struct {
}

func (m *MockAWSConfigService) PutEvaluations(input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error) {
	if strings.HasSuffix(*input.Evaluations[0].ComplianceResourceId, "-noncompliant") && *input.Evaluations[0].ComplianceType == "NON_COMPLIANT" {
		return &configservice.PutEvaluationsOutput{}, nil
	}

	if strings.HasSuffix(*input.Evaluations[0].ComplianceResourceId, "-compliant") && *input.Evaluations[0].ComplianceType == "COMPLIANT" {
		return &configservice.PutEvaluationsOutput{}, nil
	}

	return &configservice.PutEvaluationsOutput{}, fmt.Errorf("Resource should be in a different state have: %v want: %v", *input.Evaluations[0].ComplianceResourceId, *input.Evaluations[0].ComplianceType)
}
