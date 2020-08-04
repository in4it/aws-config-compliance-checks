package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/aws/aws-lambda-go/events"
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

func TestEvaluateCompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("events/example.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)

	fmt.Println("Resource state:", resp)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource COMPLIANT, should be NON_COMPLIANT")
		return
	}
}

func TestParams(t *testing.T) {
	e := events.ConfigEvent{
		EventLeftScope: true,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludedSecurityGroups\":\"test-sg-noncompliant,testBucket2\"}",
	}

	status := ""
	params := getParams(e, "excludedSecurityGroups")

	if len(params) != 2 {
		t.Errorf("Error: expected 2 results")
		return
	}

	fmt.Println("Excluded SecurityGroups:", params)

	if params := getParams(e, "excludedSecurityGroups"); params != nil {
		for _, v := range params {
			if v == "test-sg-noncompliant" {
				status = "NOT_APPLICABLE"
			}
		}
	}

	if status != "NOT_APPLICABLE" {
		t.Errorf("Error: Wrong status should be NOT_APPLICABLE")
		return
	}

}

func TestHandleRequestWithConfigServiceNonCompliant(t *testing.T) {
	ctx := context.Background()
	data, _ := ioutil.ReadFile("events/example.json")
	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludedSecurityGroups\":\"test-sg-1,test-sg-2\"}",
		InvokingEvent:  string(data),
	}

	m := &MockAWSConfigService{}
	err := handleRequestWithConfigService(ctx, configEvent, m)
	if err != nil {
		t.Error("Error:", err)
		return
	}
}
