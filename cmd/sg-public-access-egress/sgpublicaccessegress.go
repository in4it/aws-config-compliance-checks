package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
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

	list := createAllowList(getParams(configEvent, "excludeSecurityGroups"))

	if isApplicable(configurationItem, configEvent) && status == "" {
		status = evaluateCompliance(configurationItem, list)
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

func evaluateCompliance(c ConfigurationItem, list map[string][]string) string {
	if c.ResourceType != "AWS::EC2::SecurityGroup" {
		return "NOT_APPLICABLE"
	}

	for _, s := range c.Configuration.IPPermissionsEgress {
		if f := findInSlice(s.IPRanges, "0.0.0.0/0"); f == true {
			if val, ok := list[c.ResourceID]; ok {
				if len(list[c.ResourceID]) == 0 {
					fmt.Println("Skipping over Compliance check for resource", c.ResourceID, "Params: excludeSecurityGroups")
					return "NOT_APPLICABLE"
				}
				if !findInSlice(val, strconv.Itoa(s.ToPort)) && !findInSlice(val, strconv.Itoa(s.FromPort)) {
					return "NON_COMPLIANT"
				}

			} else {
				return "NON_COMPLIANT"
			}
		}
	}

	return "COMPLIANT"

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

func createAllowList(params []string) map[string][]string {
	list := make(map[string][]string)
	if params != nil {
		for _, v := range params {
			param := strings.Split(v, ":")
			if len(param) > 1 {
				ports := strings.Split(param[1], "+")
				for _, port := range ports {
					list[param[0]] = append(list[param[0]], port)
				}
			} else {
				list[param[0]] = []string{}
			}
		}
	}
	return list
}

func findInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
