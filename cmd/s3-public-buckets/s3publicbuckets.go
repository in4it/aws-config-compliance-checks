package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

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
	fmt.Printf("Params: %s\n", configEvent.RuleParameters)

	var status string
	var invokingEvent InvokingEvent
	var configurationItem ConfigurationItem

	invokingEvent, err := getInvokingEvent([]byte(configEvent.InvokingEvent))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	configurationItem = invokingEvent.ConfigurationItem

	if params := getParams(configEvent, "excludeBuckets"); params != nil {
		for _, v := range params {
			if v == configurationItem.ResourceName {
				fmt.Println("Skiping over Compliance check for resource", v, "Params: ignored")
				status = "NOT_APPLICABLE"
			}
		}
	}

	if isApplicable(configurationItem, configEvent) && status == "" {
		fmt.Println("Resource APPLICABLE for Compliance check")
		status = evaluateCompliance(configurationItem)
	} else {
		fmt.Println("Resource NOT_APPLICABLE for Compliance check")
		status = "NOT_APPLICABLE"
	}

	cSession := session.Must(session.NewSession())
	svc := configservice.New(cSession)

	var evaluations []*configservice.Evaluation

	if err != nil {
		fmt.Println(err)
		return
	}

	evaluation := &configservice.Evaluation{
		ComplianceResourceId:   aws.String(configurationItem.ResourceID),
		ComplianceResourceType: aws.String(configurationItem.ResourceType),
		ComplianceType:         aws.String(status),
		OrderingTimestamp:      aws.Time(configurationItem.ConfigurationItemCaptureTime),
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

	fmt.Printf("Evaluation completed: %s\n", out)
}

func evaluateCompliance(c ConfigurationItem) string {

	fmt.Println("Starting Evaluation Complaiance")

	if c.ResourceType != "AWS::S3::Bucket" {
		fmt.Println("Resource NOT_APPLICABLE")
		return "NOT_APPLICABLE"
	}

	blockPublicAcls := c.SupplementaryConfiguration.PublicAccessBlockConfiguration.BlockPublicAcls
	ignorePublicAcls := c.SupplementaryConfiguration.PublicAccessBlockConfiguration.IgnorePublicAcls
	blockPublicPolicy := c.SupplementaryConfiguration.PublicAccessBlockConfiguration.BlockPublicPolicy
	restrictPublicBuckets := c.SupplementaryConfiguration.PublicAccessBlockConfiguration.RestrictPublicBuckets

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

func getInvokingEvent(event []byte) (InvokingEvent, error) {
	var result InvokingEvent
	err := json.Unmarshal(event, &result)
	if err != nil {
		fmt.Println("Error:", err)
		return result, err
	}

	return result, nil
}

func isApplicable(c ConfigurationItem, e events.ConfigEvent) bool {

	status := c.ConfigurationItemStatus
	fmt.Println("Resource status:", status)
	if e.EventLeftScope == false && status == "OK" || status == "ResourceDiscovered" {
		fmt.Println("Returning status:", status)
		return true
	}
	return false

}

func getParams(p events.ConfigEvent, param string) []string {
	if p.RuleParameters != "" {
		var result map[string]string
		err := json.Unmarshal([]byte(p.RuleParameters), &result)
		if err != nil {
			fmt.Println("Error:", err)
			return nil
		}
		if _, ok := result[param]; ok {
			value := strings.Split(result[param], ",")
			return value
		}
		return nil
	}
	return nil
}
