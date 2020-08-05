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

func TestParamiters(t *testing.T) {
	data, _ := ioutil.ReadFile("events/example-compliant.json")

	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeSecurityGroups\":\"sg-1123ffff332212dddd-compliant:50051+443,test-sg-2\"}",
		InvokingEvent:  string(data),
	}

	list := createAllowList(getParams(configEvent, "excludeSecurityGroups"))

	fmt.Println(list)

	if len(list) == 0 {
		t.Errorf("error: List is empty")
		return
	}

	if len(list["test-sg-2"]) != 0 {
		t.Errorf("error: Should be an empty slice")
		return
	}

	if val, ok := list["test-sg-2"]; ok {
		fmt.Println("value of: list[\"test-sg-2\"]", val)
	} else {
		t.Errorf("error: Should be an empty slice")
		return
	}

	if val, ok := list["sg-1123ffff332212dddd-compliant"]; ok {
		fmt.Println("value of: list[\"sg-1123ffff332212dddd-compliant\"]", val)
	} else {
		t.Errorf("error: Should containe values")
		return
	}

}

func TestEvaluateCompliantNonCompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("events/example.json")

	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludedSecurityGroups\":\"test-sg-1,test-sg-2\"}",
		InvokingEvent:  string(data),
	}

	m, err := getInvokingEvent([]byte(configEvent.InvokingEvent))

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	list := createAllowList(getParams(configEvent, "excludeSecurityGroups"))

	resp := evaluateCompliance(ci, list)

	fmt.Println("Resource state:", resp)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource COMPLIANT, should be NON_COMPLIANT")
		return
	}
}

func TestEvaluateCompliantCompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("events/example-compliant-port.json")

	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeSecurityGroups\":\"sg-1123ffff332212dddd-compliant:50051+443,test-sg-2\"}",
		InvokingEvent:  string(data),
	}

	m, err := getInvokingEvent([]byte(configEvent.InvokingEvent))

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	list := createAllowList(getParams(configEvent, "excludeSecurityGroups"))

	resp := evaluateCompliance(ci, list)

	fmt.Println("Resource state:", resp)

	if resp == "NON_COMPLIANT" {
		t.Errorf("error: Resource NON_COMPLIANT, should be COMPLIANT")
		return
	}
}

func TestEvaluateCompliantNonCompliantPort(t *testing.T) {
	data, _ := ioutil.ReadFile("events/example.json")

	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeSecurityGroups\":\"sg-1123ffff332212dddd-noncompliant:80+443,test-sg-2\"}",
		InvokingEvent:  string(data),
	}

	m, err := getInvokingEvent([]byte(configEvent.InvokingEvent))

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	list := createAllowList(getParams(configEvent, "excludeSecurityGroups"))

	resp := evaluateCompliance(ci, list)

	fmt.Println("Resource state:", resp)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource COMPLIANT, should be NON_COMPLIANT")
		return
	}
}

func TestHandleRequestWithConfigServiceNonCompliant(t *testing.T) {
	ctx := context.Background()
	data, _ := ioutil.ReadFile("events/example.json")
	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeSecurityGroups\":\"test-sg-1,test-sg-2,sg-1123ffff332212dddd-noncompliant:80+443\"}",
		InvokingEvent:  string(data),
	}

	m := &MockAWSConfigService{}
	err := handleRequestWithConfigService(ctx, configEvent, m)
	if err != nil {
		t.Error("Error:", err)
		return
	}
}

func TestHandleRequestWithConfigServiceCompliant(t *testing.T) {
	ctx := context.Background()
	data, _ := ioutil.ReadFile("events/example-compliant.json")
	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeSecurityGroups\":\"test-sg-1,test-sg-2\"}",
		InvokingEvent:  string(data),
	}

	m := &MockAWSConfigService{}
	err := handleRequestWithConfigService(ctx, configEvent, m)
	if err != nil {
		t.Error("Error:", err)
		return
	}
}
