package main

import "time"

type InvokingEvent struct {
	ConfigurationItemDiff    interface{}       `json:"configurationItemDiff"`
	ConfigurationItem        ConfigurationItem `json:"configurationItem"`
	NotificationCreationTime time.Time         `json:"notificationCreationTime"`
	MessageType              string            `json:"messageType"`
	RecordVersion            string            `json:"recordVersion"`
}
type Relationships struct {
	ResourceID   string      `json:"resourceId"`
	ResourceName interface{} `json:"resourceName"`
	ResourceType string      `json:"resourceType"`
	Name         string      `json:"name"`
}
type UserIDGroupPairs struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}
type IPPermissions struct {
	FromPort         int                `json:"fromPort"`
	IPProtocol       string             `json:"ipProtocol"`
	Ipv6Ranges       []interface{}      `json:"ipv6Ranges"`
	PrefixListIds    []interface{}      `json:"prefixListIds"`
	ToPort           int                `json:"toPort"`
	UserIDGroupPairs []UserIDGroupPairs `json:"userIdGroupPairs"`
	Ipv4Ranges       []Ipv4Ranges       `json:"ipv4Ranges"`
	IPRanges         []string           `json:"ipRanges"`
}
type Ipv4Ranges struct {
	CidrIP      string `json:"cidrIp"`
	Description string `json:"description"`
}
type IPPermissionsEgress struct {
	IPProtocol       string        `json:"ipProtocol"`
	Ipv6Ranges       []interface{} `json:"ipv6Ranges"`
	PrefixListIds    []interface{} `json:"prefixListIds"`
	UserIDGroupPairs []interface{} `json:"userIdGroupPairs"`
	Ipv4Ranges       []Ipv4Ranges  `json:"ipv4Ranges"`
	IPRanges         []string      `json:"ipRanges"`
	FromPort         int           `json:"fromPort,omitempty"`
	ToPort           int           `json:"toPort,omitempty"`
}
type Configuration struct {
	Description         string                `json:"description"`
	GroupName           string                `json:"groupName"`
	IPPermissions       []IPPermissions       `json:"ipPermissions"`
	OwnerID             string                `json:"ownerId"`
	GroupID             string                `json:"groupId"`
	IPPermissionsEgress []IPPermissionsEgress `json:"ipPermissionsEgress"`
	Tags                []interface{}         `json:"tags"`
	VpcID               string                `json:"vpcId"`
}
type SupplementaryConfiguration struct {
}
type Tags struct {
}
type ConfigurationItem struct {
	RelatedEvents                []interface{}              `json:"relatedEvents"`
	Relationships                []Relationships            `json:"relationships"`
	Configuration                Configuration              `json:"configuration"`
	SupplementaryConfiguration   SupplementaryConfiguration `json:"supplementaryConfiguration"`
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
	ResourceCreationTime         interface{}                `json:"resourceCreationTime"`
}
