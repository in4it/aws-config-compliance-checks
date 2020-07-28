package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, configEvent events.ConfigEvent) {
	fmt.Printf("AWS Config rule: %s\n", configEvent.ConfigRuleName)
	fmt.Printf("Invoking event JSON: %s\n", configEvent.InvokingEvent)
	fmt.Printf("Event version: %s\n", configEvent.Version)

	var status string

	m, err := getInvokingEvent([]byte(configEvent.InvokingEvent))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	if isApplicable(m, configEvent) {
		status = evaluateCompliance(m["configurationItem"])
	} else {
		status = "NOT_APPLICABLE"
	}

	cSession := session.Must(session.NewSession())
	svc := configservice.New(cSession)

	var evaluations []*configservice.Evaluation

	sTime := m["configurationItem"].(map[string]interface{})["configurationItemCaptureTime"]
	layout := "2020-07-27T13:57:15.776Z"
	t, err := time.Parse(layout, sTime.(string))
	if err != nil {
		fmt.Println(err)
	}

	complianceResourceID := m["configurationItem"].(map[string]interface{})["resourceId"]
	complianceResourceType := m["configurationItem"].(map[string]interface{})["resourceType"]

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
	}

	out, err := svc.PutEvaluations(putEvaluations)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Evaluation compleated: %s\n", out)
}

func evaluateCompliance(c interface{}) string {

	if c.(map[string]interface{})["resourceType"] != "AWS::S3::Bucket" {
		return "NOT_APPLICABLE"
	}

	sc := c.(map[string]interface{})["supplementaryConfiguration"]

	if !checkDefined(sc, "PublicAccessBlockConfiguration") {
		return "NON_COMPLIANT"
	}

	pacl := sc.(map[string]interface{})["PublicAccessBlockConfiguration"]

	blockPublicAcls := pacl.(map[string]interface{})["blockPublicAcls"]
	ignorePublicAcls := pacl.(map[string]interface{})["ignorePublicAcls"]
	blockPublicPolicy := pacl.(map[string]interface{})["blockPublicPolicy"]
	restrictPublicBuckets := pacl.(map[string]interface{})["restrictPublicBuckets"]

	fmt.Printf("blockPublicAcls %v\n", blockPublicAcls)
	fmt.Printf("blockPublicAcls %v\n", ignorePublicAcls)
	fmt.Printf("blockPublicAcls %v\n", blockPublicPolicy)
	fmt.Printf("blockPublicAcls %v\n", restrictPublicBuckets)

	if blockPublicAcls == true && ignorePublicAcls == true && blockPublicPolicy == true && restrictPublicBuckets == true {
		return "COMPLIANT"
	}

	return "NON_COMPLIANT"
}

func getInvokingEvent(event []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(event, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return result, nil
}

func checkDefined(m interface{}, k string) bool {
	_, ok := m.(map[string]interface{})[k]
	if ok {
		return true
	}
	return false
}

func isApplicable(s interface{}, e events.ConfigEvent) bool {

	status := s.(map[string]interface{})["configurationItemStatus"]

	if e.EventLeftScope == false && status == "OK" || status == "ResourceDiscovered" {
		return true
	}
	return false

}
