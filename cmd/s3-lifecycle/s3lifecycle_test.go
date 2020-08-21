package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/in4it/aws-config-compliance-checks/pkg/mocks"
)

func TestEvaluateComplianceCompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("events/compliant.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)
	fmt.Println(resp)

	if resp != "COMPLIANT" {
		t.Errorf("error: Resource compliant, should be compliant")
		return
	}

}

func TestEvaluateComplianceNonCompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("events/noncompliant.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)
	fmt.Println(resp)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource compliant, should be NON_compliant")
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
	data, _ := ioutil.ReadFile("events/noncompliant.json")
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
