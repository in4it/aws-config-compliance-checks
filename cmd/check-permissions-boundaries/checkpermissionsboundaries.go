package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

type AWSConfigService interface {
	PutEvaluations(*configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
}

type container struct {
	Field customSlice
}

type customSlice []string

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
	if c.ResourceType != "AWS::IAM::Policy" {
		return "NOT_APPLICABLE"
	}

	pn := c.ResourceName
	fmt.Printf("ResourceName: %v \n", pn)
    ms := "permissions-boundary"
	if strings.Contains(pn, ms) {
		encodedValue := c.Configuration.PolicyVersionList[0].Document
		pl, err := url.QueryUnescape(encodedValue)
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println("UNESCAPE ERROR")
			return "NON_COMPLIANT"
		}
		var pdf = new(PolicyDocument)

		err2 := json.Unmarshal([]byte(pl), &pdf)

		fmt.Printf("%v", pl)
		if err2 != nil {
			fmt.Println("Error:", err2)
			fmt.Println("UNMASHALL ERROR")
			return "NON_COMPLIANT"
		}

		var containerStringEquals container
		var containerForallvaluesStringnotequals container
		var containerForanyvalueStringequals container

		StringEquals := pdf.Statement[0].Condition.StringEquals.AwsSourcevpc
		ForallvaluesStringnotequals := pdf.Statement[0].Condition.ForallvaluesArnnotequals.AwsSourcevpc
		ForanyvalueStringequals := pdf.Statement[0].Condition.ForanyvalueStringequals.AwsSourcevpc

        err3 := json.Unmarshal([]byte(StringEquals.(string)), &containerStringEquals)
		if err3 != nil {
			panic(err3)
		}
		err4 := json.Unmarshal([]byte(ForallvaluesStringnotequals.(string)), &containerForallvaluesStringnotequals)
		if err4 != nil {
			panic(err4)
		}
		err5 := json.Unmarshal([]byte(ForanyvalueStringequals.(string)), &containerForanyvalueStringequals)
		if err != nil {
			panic(err5)
		}


		if len(StringEquals.(string)) > 0 || len(ForallvaluesStringnotequals.([]string)) > 0 || len(ForanyvalueStringequals.([]string)) > 0 {
			return "COMPLIANT"
		}
	}
	return "NOT_APPLICABLE"
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

func (c *customSlice) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	slice, ok := tmp.([]interface{})
	if ok {
		for _, item := range slice {
			*c = append(*c, item.(string))
		}
		return nil
	}
	theString, ok := tmp.(string)
	if ok {
		*c = append(*c, theString)
		return nil
	}
	return errors.New("Field neither slice or string")
}
