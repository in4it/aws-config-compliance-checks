package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestEvaluateComplianceNoNcompliant(t *testing.T) {
	data, _ := ioutil.ReadFile("events/case1.json")
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

	data, _ := ioutil.ReadFile("events/case1.json")
	m, err := getInvokingEvent(data)

	if err != nil {
		t.Errorf("error: %s", err)
		return
	}

	ci := m.ConfigurationItem

	resp := evaluateCompliance(ci)
	fmt.Println(resp)

	if resp == "NON_COMPLIANT" {
		t.Errorf("error: Resource NON_COMPLIANT, should be COMPLIANT")
		return
	}

}
