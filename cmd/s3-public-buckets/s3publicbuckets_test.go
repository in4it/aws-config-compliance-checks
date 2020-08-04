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

func TestDataCreate(t *testing.T) {
	data, _ := ioutil.ReadFile("test/create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	fmt.Println(m.ConfigurationItem.ResourceID)

}

func TestDataUpdate(t *testing.T) {
	data, _ := ioutil.ReadFile("test/update.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	fmt.Println(m.ConfigurationItem.ResourceID)

}

func TestDataDelete(t *testing.T) {
	data, _ := ioutil.ReadFile("test/delete.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	fmt.Println(m.ConfigurationItem.ResourceID)

}

func TestIfApplicableOnCreate(t *testing.T) {
	data, _ := ioutil.ReadFile("test/create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	e := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
	}

	if a := isApplicable(ci, e); !a {
		t.Errorf("error: Resource NOT_APPLICABLE should be APPLICABLE")
		return
	}

}

func TestIfApplicableOnUpdate(t *testing.T) {
	data, _ := ioutil.ReadFile("test/update.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	e := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
	}

	if a := isApplicable(ci, e); !a {
		t.Errorf("error: Resource NOT_APPLICABLE should be APPLICABLE")
		return
	}

}

func TestIfNotApplicable(t *testing.T) {
	data, _ := ioutil.ReadFile("test/delete.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	e := events.ConfigEvent{
		EventLeftScope: true,
		ResultToken:    "myResultToken",
	}

	if a := isApplicable(ci, e); a {
		t.Errorf("error: Resource NOT_APPLICABLE should be APPLICABLE")
		return
	}

}

func TestEvaluateComplianceNotcompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("test/create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)
	fmt.Println(resp)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource compliant, should be NOT_compliant")
		return
	}

}

func TestEvaluateComplianceNotApplicable(t *testing.T) {
	data, _ := ioutil.ReadFile("test/delete.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)
	fmt.Println(resp)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource compliant, should be NOT_compliant")
		return
	}

}

func TestEvaluateCompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("test/update.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource compliant, should be NOT_compliant")
		return
	}

}

func TestParams(t *testing.T) {
	e := events.ConfigEvent{
		EventLeftScope: true,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeBuckets\":\"testBucket1,testBucket2\"}",
	}

	status := ""
	params := getParams(e, "excludeBuckets")

	if len(params) != 2 {
		t.Errorf("Error: expected 2 results")
		return
	}

	fmt.Println("Ignored buckets:", params)

	if params := getParams(e, "excludeBuckets"); params != nil {
		for _, v := range params {
			if v == "testBucket1" {
				status = "NOT_APPLICABLE"
			}
		}
	}

	if status != "NOT_APPLICABLE" {
		t.Errorf("Error: Wrong status should be NOT_APPLICABLE")
		return
	}

}

func TestHandleRequestWithConfigServicecompliant(t *testing.T) {
	ctx := context.Background()
	data, _ := ioutil.ReadFile("test/compliant.json")
	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeBuckets\":\"testBucket1,testBucket2\"}",
		InvokingEvent:  string(data),
	}

	m := &MockAWSConfigService{}
	err := handleRequestWithConfigService(ctx, configEvent, m)
	if err != nil {
		t.Error("Error:", err)
		return
	}
}

func TestHandleRequestWithConfigServiceNoncompliant(t *testing.T) {
	ctx := context.Background()
	data, _ := ioutil.ReadFile("test/create.json")
	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeBuckets\":\"testBucket1,testBucket2\"}",
		InvokingEvent:  string(data),
	}

	m := &MockAWSConfigService{}
	err := handleRequestWithConfigService(ctx, configEvent, m)
	if err != nil {
		t.Error("Error:", err)
		return
	}
}

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
