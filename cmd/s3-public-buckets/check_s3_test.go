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

	var ci configurationItem
	ci = m["configurationItem"].(map[string]interface{})

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
	data, _ := ioutil.ReadFile("test/request_delete.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	var ci configurationItem
	ci = m["configurationItem"].(map[string]interface{})

	e := events.ConfigEvent{
		EventLeftScope: true,
		ResultToken:    "myResultToken",
	}

	if a := isApplicable(ci, e); a {
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
	var ci configurationItem
	ci = m["configurationItem"].(map[string]interface{})

	resp := evaluateCompliance(ci)
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

	var ci configurationItem
	ci = m["configurationItem"].(map[string]interface{})

	resp := evaluateCompliance(ci)

	if resp == "COMPLIANT" {
		t.Errorf("error: Resource COMPLAIANT, should be NOT_COMPLAIANT")
		return
	}

}

func TestTimeParser(t *testing.T) {
	data, _ := ioutil.ReadFile("test/request_update.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	sTime := m["configurationItem"].(map[string]interface{})["configurationItemCaptureTime"]
	pTime, err := parseTime(sTime.(string))

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	fmt.Println(pTime)

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
