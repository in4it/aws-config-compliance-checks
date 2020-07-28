package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestData(t *testing.T) {
	data, _ := ioutil.ReadFile("test/request_create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if len(m) == 0 {
		t.Errorf("error: Key not found")
		return
	}
	fmt.Println(m["configurationItem"].(map[string]interface{})["resourceId"])

}

func TestCheckDefined(t *testing.T) {
	data, _ := ioutil.ReadFile("test/request_create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}
	if len(m) == 0 {
		t.Errorf("error: Key not found")
		return
	}
	ok := checkDefined(m["configurationItem"], "resourceName")

	if !ok {
		t.Errorf("error: Key not found: resourceName")
		return
	}
	fmt.Println(m["configurationItem"].(map[string]interface{})["resourceName"])

	ok = checkDefined(m["configurationItem"], "something")

	if ok {
		t.Errorf("error: Not existing fey found")
		return
	}
}

func TestIfApplicable(t *testing.T) {
	data, _ := ioutil.ReadFile("test/request_create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	e := events.ConfigEvent{
		EventLeftScope: false,
		ResultToken:    "myResultToken",
	}

	if a := isApplicable(m["configurationItem"], e); !a {
		t.Errorf("error: Resource NOT_APPLICABLE should be APPLICABLE")
		return
	}

}

func TestIfNotApplicable(t *testing.T) {
	data, _ := ioutil.ReadFile("test/request_delete.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	e := events.ConfigEvent{
		EventLeftScope: true,
		ResultToken:    "myResultToken",
	}

	if a := isApplicable(m["configurationItem"], e); a {
		t.Errorf("error: Resource NOT_APPLICABLE should be APPLICABLE")
		return
	}

}

func TestEvaluateComplianceNotComplaiant(t *testing.T) {
	data, _ := ioutil.ReadFile("test/request_create.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	resp := evaluateCompliance(m["configurationItem"])
	fmt.Println(resp)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource COMPLAIANT, should be NOT_COMPLAIANT")
		return
	}

}

func TestEvaluateComplianceComplaiant(t *testing.T) {
	data, _ := ioutil.ReadFile("test/request_update.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	resp := evaluateCompliance(m["configurationItem"])

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource COMPLAIANT, should be NOT_COMPLAIANT")
		return
	}

}
