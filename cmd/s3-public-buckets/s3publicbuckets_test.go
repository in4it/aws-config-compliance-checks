package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/in4it/aws-config-compliance-checks/pkg/mocks"
)

func TestDataCreate(t *testing.T) {
	data, _ := ioutil.ReadFile("events/create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	fmt.Println(m.ConfigurationItem.ResourceID)

}

func TestDataUpdate(t *testing.T) {
	data, _ := ioutil.ReadFile("events/update.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	fmt.Println(m.ConfigurationItem.ResourceID)

}

func TestDataDelete(t *testing.T) {
	data, _ := ioutil.ReadFile("events/delete.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	fmt.Println(m.ConfigurationItem.ResourceID)

}

func TestIfApplicableOnCreate(t *testing.T) {
	data, _ := ioutil.ReadFile("events/create.json")
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
	data, _ := ioutil.ReadFile("events/update.json")
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
	data, _ := ioutil.ReadFile("events/delete.json")
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
	data, _ := ioutil.ReadFile("events/create.json")
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
	data, _ := ioutil.ReadFile("events/delete.json")
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

	data, _ := ioutil.ReadFile("events/update.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource COMPLIANT, should be NON_COMPLIANT")
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

func TestHandleRequestWithConfigServiceCompliant(t *testing.T) {
	ctx := context.Background()
	data, _ := ioutil.ReadFile("events/compliant.json")
	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeBuckets\":\"testBucket1,testBucket2\"}",
		InvokingEvent:  string(data),
	}

	m := &mocks.MockAWSConfigService{}
	err := handleRequestWithConfigService(ctx, configEvent, m)
	if err != nil {
		t.Error("Error:", err)
		return
	}
}

func TestHandleRequestWithConfigServiceNoncompliant(t *testing.T) {
	ctx := context.Background()
	data, _ := ioutil.ReadFile("events/create.json")
	configEvent := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
		RuleParameters: "{\"excludeBuckets\":\"testBucket1,testBucket2\"}",
		InvokingEvent:  string(data),
	}

	m := &mocks.MockAWSConfigService{}
	err := handleRequestWithConfigService(ctx, configEvent, m)
	if err != nil {
		t.Error("Error:", err)
		return
	}
}
