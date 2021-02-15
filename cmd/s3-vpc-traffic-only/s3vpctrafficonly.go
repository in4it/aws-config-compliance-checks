package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"strings"
)

type AWSConfigService interface {
	PutEvaluations(*configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, configEvent events.ConfigEvent) error {
	cSession := session.Must(session.NewSession())
	svc := configservice.New(cSession)
	err := handleRequestWithConfigService(ctx, configEvent, svc)
	if err != nil {
		return err
	}
	return nil
}

func handleRequestWithConfigService(ctx context.Context, configEvent events.ConfigEvent, svc AWSConfigService) error {
	var status string
	var invokingEvent InvokingEvent
	var configurationItem ConfigurationItem

	invokingEvent, err := getInvokingEvent([]byte(configEvent.InvokingEvent))

	if err != nil {
		panic(err)
	}

	configurationItem = invokingEvent.ConfigurationItem

	if params := getParams(configEvent, "excludeBuckets"); params != nil {
		for _, v := range params {
			if v == configurationItem.ResourceName {
				fmt.Println("Skipping over Compliance check for resource", v, "Params: excludeBuckets")
				status = "NOT_APPLICABLE"
			}
		}
	}

	if isApplicable(configurationItem, configEvent) && status == "" {
		status = evaluateCompliance(configurationItem)
	} else {
		status = "NOT_APPLICABLE"
	}

	var evaluations []*configservice.Evaluation

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

	_, err = svc.PutEvaluations(putEvaluations)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

func evaluateCompliance(c ConfigurationItem) string {
	if c.ResourceType != "AWS::S3::Bucket" {
		return "NOT_APPLICABLE"
	}

	pd := c.SupplementaryConfiguration.BucketPolicy.PolicyText

	if pd != nil {
		pds := pd.(string)
		var pdf = new(PolicyDocument)
		err := json.Unmarshal([]byte(pds), &pdf)
		if err != nil {
			fmt.Println("Error Unmarshal :", err)
			return "NON_COMPLIANT"
		}
		if pdf.Statement[0].Condition.StringEquals.AwsSourceVpc != nil || pdf.Statement[0].Condition.ForAllValuesStringNotEquals.AwsSourceVpc != nil {
			return "COMPLIANT"
		}
	}
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
	if e.EventLeftScope == false && status == "OK" || status == "ResourceDiscovered" {
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
