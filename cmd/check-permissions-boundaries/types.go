package main

import (
	"time"
)

type InvokingEvent struct {
	ConfigurationItemDiff    ConfigurationItemDiff `json:"configurationItemDiff"`
	ConfigurationItem        ConfigurationItem     `json:"configurationItem"`
	NotificationCreationTime time.Time             `json:"notificationCreationTime"`
	MessageType              string                `json:"messageType"`
	RecordVersion            string                `json:"recordVersion"`
}
type SupplementaryConfigurationPublicAccessBlockConfigurationBlockPublicAcls struct {
	PreviousValue bool   `json:"previousValue"`
	UpdatedValue  bool   `json:"updatedValue"`
	ChangeType    string `json:"changeType"`
}
type SupplementaryConfigurationPublicAccessBlockConfigurationBlockPublicPolicy struct {
	PreviousValue bool   `json:"previousValue"`
	UpdatedValue  bool   `json:"updatedValue"`
	ChangeType    string `json:"changeType"`
}
type SupplementaryConfigurationPublicAccessBlockConfigurationRestrictPublicBuckets struct {
	PreviousValue bool   `json:"previousValue"`
	UpdatedValue  bool   `json:"updatedValue"`
	ChangeType    string `json:"changeType"`
}
type SupplementaryConfigurationPublicAccessBlockConfigurationIgnorePublicAcls struct {
	PreviousValue bool   `json:"previousValue"`
	UpdatedValue  bool   `json:"updatedValue"`
	ChangeType    string `json:"changeType"`
}
type ChangedProperties struct {
	SupplementaryConfigurationPublicAccessBlockConfigurationBlockPublicAcls       SupplementaryConfigurationPublicAccessBlockConfigurationBlockPublicAcls       `json:"SupplementaryConfiguration.PublicAccessBlockConfiguration.BlockPublicAcls"`
	SupplementaryConfigurationPublicAccessBlockConfigurationBlockPublicPolicy     SupplementaryConfigurationPublicAccessBlockConfigurationBlockPublicPolicy     `json:"SupplementaryConfiguration.PublicAccessBlockConfiguration.BlockPublicPolicy"`
	SupplementaryConfigurationPublicAccessBlockConfigurationRestrictPublicBuckets SupplementaryConfigurationPublicAccessBlockConfigurationRestrictPublicBuckets `json:"SupplementaryConfiguration.PublicAccessBlockConfiguration.RestrictPublicBuckets"`
	SupplementaryConfigurationPublicAccessBlockConfigurationIgnorePublicAcls      SupplementaryConfigurationPublicAccessBlockConfigurationIgnorePublicAcls      `json:"SupplementaryConfiguration.PublicAccessBlockConfiguration.IgnorePublicAcls"`
}
type ConfigurationItemDiff struct {
	ChangedProperties ChangedProperties `json:"changedProperties"`
	ChangeType        string            `json:"changeType"`
}
type Owner struct {
	DisplayName interface{} `json:"displayName"`
	ID          string      `json:"id"`
}

type Tags struct {
}
type ConfigurationItem struct {
	RelatedEvents                []interface{}              `json:"relatedEvents"`
	Relationships                []interface{}              `json:"relationships"`
	Configuration                Configuration              `json:"configuration"`
	SupplementaryConfiguration   interface{}                `json:"supplementaryConfiguration"`
	Tags                         Tags                       `json:"tags"`
	ConfigurationItemVersion     string                     `json:"configurationItemVersion"`
	ConfigurationItemCaptureTime time.Time                  `json:"configurationItemCaptureTime"`
	ConfigurationStateID         int64                      `json:"configurationStateId"`
	AwsAccountID                 string                     `json:"awsAccountId"`
	ConfigurationItemStatus      string                     `json:"configurationItemStatus"`
	ResourceType                 string                     `json:"resourceType"`
	ResourceID                   string                     `json:"resourceId"`
	ResourceName                 string                     `json:"resourceName"`
	ARN                          string                     `json:"ARN"`
	AwsRegion                    string                     `json:"awsRegion"`
	AvailabilityZone             string                     `json:"availabilityZone"`
	ConfigurationStateMd5Hash    string                     `json:"configurationStateMd5Hash"`
	ResourceCreationTime         time.Time                  `json:"resourceCreationTime"`
}

type Configuration struct {
	Name         string    `json:"name"`
	Owner        Owner     `json:"owner"`
	CreationDate time.Time `json:"creationDate"`
	PolicyVersionList []PolicyVersion `json:"policyVersionList"`
}

type PolicyVersion struct {
	CreateDate time.Time  `json:"createDate"`
	Document string  `json:"document"`
	IsDefaultVersion bool  `json:"isDefaultVersion"`
	VersionId string  `json:"versionId"`
}

type PolicyDocument struct {
	Version   string         `json:"Version"`
	ID        string         `json:"Id"`
	Statement StatementEntry `json:"Statement"`
}

type StatementEntry []struct {
	Sid       string         `json:"Sid"`
	Effect    string         `json:"Effect"`
	Principal interface{}    `json:"Principal"`
	Action    interface{}    `json:"Action"`
	Resource  interface{}   `json:"Resource"`
	Condition ConditionEntry `json:"Condition,omitempty"`
}
type ConditionEntry struct {
	StringEquals                StringEqualsEntry                `json:"StringEquals"`
	ForallvaluesStringnotequals ForallvaluesArnnotequalsentry    `json:"ForAllValues:StringNotEquals"`
	ForallvaluesArnnotequals    ForallvaluesStringnotequalsentry `json:"ForAllValues:ArnNotEquals"`
	ForanyvalueStringequals     ForanyvalueStringequalsentry	 `json:"ForAnyValue:StringEquals"`
	ArnNotLike                  ArnNotLikeEntry                  `json:"ArnNotLike"`
}

type ForallvaluesArnnotequalsentry struct {
	AwsPrincipalArn string `json:"aws:PrincipalArn"`
}

type StringEqualsEntry struct {
	AwsSourcevpc string `json:"aws:SourceVpc"`
}

type ForanyvalueStringequalsentry  struct {
	AwsSourcevpc interface{} `json:"aws:SourceVpc"`
}

type ForallvaluesStringnotequalsentry struct {
	AwsSourcevpc interface{} `json:"aws:SourceVpc"`
	AwsCalledvia string   `json:"aws:CalledVia"`
}
type ArnNotLikeEntry struct {
	AwsUsername     string   `json:"aws:Username"`
	AwsPrincipalarn []string `json:"aws:PrincipalArn"`
}

type container struct {
	Field customSlice `json:"aws:SourceVpc"`
}

type customSlice []string
