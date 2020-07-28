package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

type invokingEvent map[string]interface{}
type configurationItem map[string]interface{}
type supplementaryConfiguration map[string]interface{}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, configEvent events.ConfigEvent) {
	fmt.Printf("AWS Config rule: %s\n", configEvent.ConfigRuleName)
	fmt.Printf("Invoking event JSON: %s\n", configEvent.InvokingEvent)
	fmt.Printf("Event version: %s\n", configEvent.Version)
	fmt.Printf("Params: %s\n", configEvent.RuleParameters)

	var status string
	var m invokingEvent
	var ci configurationItem

	m, err := getInvokingEvent([]byte(configEvent.InvokingEvent))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	ci = m["configurationItem"].(map[string]interface{})

	if params := getParams(configEvent); params != nil {
		for _, v := range params {
			if v == ci["resourceName"] {
				fmt.Println("Skiping over Compliance check for resource", v, "Params: ignored")
				status = "NOT_APPLICABLE"
			}
		}
	}

	if isApplicable(ci, configEvent) && status == "" {
		fmt.Println("Resource APPLICABLE for Compliance check")
		status = evaluateCompliance(ci)
	} else {
		fmt.Println("Resource NOT_APPLICABLE for Compliance check")
		status = "NOT_APPLICABLE"
	}

	cSession := session.Must(session.NewSession())
	svc := configservice.New(cSession)

	var evaluations []*configservice.Evaluation

	sTime := ci["configurationItemCaptureTime"]
	t, err := parseTime(sTime.(string))

	if err != nil {
		fmt.Println(err)
		return
	}

	complianceResourceID := ci["resourceId"]
	complianceResourceType := ci["resourceType"]

	evaluation := &configservice.Evaluation{
		ComplianceResourceId:   aws.String(complianceResourceID.(string)),
		ComplianceResourceType: aws.String(complianceResourceType.(string)),
		ComplianceType:         aws.String(status),
		OrderingTimestamp:      aws.Time(t),
	}

	evaluations = append(evaluations, evaluation)
	putEvaluations := &configservice.PutEvaluationsInput{
		Evaluations: evaluations,
		ResultToken: aws.String(configEvent.ResultToken),
		TestMode:    aws.Bool(false),
	}

	fmt.Printf("Evaluation: %s\n%s\n", evaluations, configEvent.ResultToken)

	out, err := svc.PutEvaluations(putEvaluations)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Evaluation compleated: %s\n", out)
}

func evaluateCompliance(c configurationItem) string {

	fmt.Println("Starting Evaluation Complaiance")

	var sc supplementaryConfiguration

	if c["resourceType"] != "AWS::S3::Bucket" {
		fmt.Println("Resource NOT_APPLICABLE")
		return "NOT_APPLICABLE"
	}

	if !checkDefined(c["supplementaryConfiguration"], "PublicAccessBlockConfiguration") {
		fmt.Println("Resource NON_COMPLIANT")
		return "NON_COMPLIANT"
	}

	sc = c["supplementaryConfiguration"].(map[string]interface{})
	pacl := sc["PublicAccessBlockConfiguration"].(map[string]interface{})

	blockPublicAcls := pacl["blockPublicAcls"]
	ignorePublicAcls := pacl["ignorePublicAcls"]
	blockPublicPolicy := pacl["blockPublicPolicy"]
	restrictPublicBuckets := pacl["restrictPublicBuckets"]

	fmt.Printf("blockPublicAcls %v\n", blockPublicAcls)
	fmt.Printf("blockPublicAcls %v\n", ignorePublicAcls)
	fmt.Printf("blockPublicAcls %v\n", blockPublicPolicy)
	fmt.Printf("blockPublicAcls %v\n", restrictPublicBuckets)

	if blockPublicAcls == true && ignorePublicAcls == true && blockPublicPolicy == true && restrictPublicBuckets == true {
		fmt.Println("Resource COMPLIANT")
		return "COMPLIANT"
	}

	fmt.Println("Resource COMPLIANT")
	return "NON_COMPLIANT"
}

func getInvokingEvent(event []byte) (invokingEvent, error) {
	var result invokingEvent
	err := json.Unmarshal(event, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return result, nil
}

func checkDefined(m interface{}, k string) bool {
	fmt.Println("Checking if defined:", k)
	_, ok := m.(map[string]interface{})[k]
	if ok {
		fmt.Println(k, "is defined")
		return true
	}
	fmt.Println(k, "is not defined")
	return false
}

func isApplicable(s configurationItem, e events.ConfigEvent) bool {

	status := s["configurationItemStatus"]
	fmt.Println("Resource status:", status)
	if e.EventLeftScope == false && status == "OK" || status == "ResourceDiscovered" {
		fmt.Println("Returning status:", status)
		return true
	}
	return false

}

func getParams(p events.ConfigEvent) []string {
	if p.RuleParameters != "" {
		var result map[string]string
		err := json.Unmarshal([]byte(p.RuleParameters), &result)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		if _, ok := result["ignored"]; ok {
			ignoredBuckets := strings.Split(result["ignored"], ",")
			return ignoredBuckets
		}
		return nil
	}
	return nil
}

func parseTime(s string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		fmt.Println(err)
		return t, err
	}

	return t, nil
}
