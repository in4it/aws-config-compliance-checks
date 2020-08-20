package main

import "time"

type InvokingEvent struct {
	ConfigurationItemDiff    ConfigurationItemDiff `json:"configurationItemDiff"`
	ConfigurationItem        ConfigurationItem     `json:"configurationItem"`
	NotificationCreationTime time.Time             `json:"notificationCreationTime"`
	MessageType              string                `json:"messageType"`
	RecordVersion            string                `json:"recordVersion"`
}
type ChangedProperties struct {
}
type ConfigurationItemDiff struct {
	ChangedProperties ChangedProperties `json:"changedProperties"`
	ChangeType        string            `json:"changeType"`
}
type Owner struct {
	DisplayName interface{} `json:"displayName"`
	ID          string      `json:"id"`
}
type Configuration struct {
	Name         string    `json:"name"`
	Owner        Owner     `json:"owner"`
	CreationDate time.Time `json:"creationDate"`
}
type BucketAccelerateConfiguration struct {
	Status interface{} `json:"status"`
}
type Predicate struct {
	Type   string `json:"type"`
	Prefix string `json:"prefix"`
}
type Filter struct {
	Predicate Predicate `json:"predicate"`
}
type Rules struct {
	ID                                string      `json:"id"`
	Prefix                            interface{} `json:"prefix"`
	Status                            string      `json:"status"`
	Filter                            Filter      `json:"filter"`
	ExpirationInDays                  int         `json:"expirationInDays"`
	ExpiredObjectDeleteMarker         bool        `json:"expiredObjectDeleteMarker"`
	NoncurrentVersionExpirationInDays int         `json:"noncurrentVersionExpirationInDays"`
	ExpirationDate                    interface{} `json:"expirationDate"`
	Transitions                       interface{} `json:"transitions"`
	NoncurrentVersionTransitions      interface{} `json:"noncurrentVersionTransitions"`
	AbortIncompleteMultipartUpload    interface{} `json:"abortIncompleteMultipartUpload"`
}
type BucketLifecycleConfiguration struct {
	Rules []Rules `json:"rules"`
}
type BucketLoggingConfiguration struct {
	DestinationBucketName interface{} `json:"destinationBucketName"`
	LogFilePrefix         interface{} `json:"logFilePrefix"`
}
type Configurations struct {
}
type BucketNotificationConfiguration struct {
	Configurations Configurations `json:"configurations"`
}
type BucketPolicy struct {
	PolicyText interface{} `json:"policyText"`
}
type BucketVersioningConfiguration struct {
	Status             string      `json:"status"`
	IsMfaDeleteEnabled interface{} `json:"isMfaDeleteEnabled"`
}
type PublicAccessBlockConfiguration struct {
	BlockPublicAcls       bool `json:"blockPublicAcls"`
	IgnorePublicAcls      bool `json:"ignorePublicAcls"`
	BlockPublicPolicy     bool `json:"blockPublicPolicy"`
	RestrictPublicBuckets bool `json:"restrictPublicBuckets"`
}
type SupplementaryConfiguration struct {
	AccessControlList               string                          `json:"AccessControlList"`
	BucketAccelerateConfiguration   BucketAccelerateConfiguration   `json:"BucketAccelerateConfiguration"`
	BucketLifecycleConfiguration    BucketLifecycleConfiguration    `json:"BucketLifecycleConfiguration"`
	BucketLoggingConfiguration      BucketLoggingConfiguration      `json:"BucketLoggingConfiguration"`
	BucketNotificationConfiguration BucketNotificationConfiguration `json:"BucketNotificationConfiguration"`
	BucketPolicy                    BucketPolicy                    `json:"BucketPolicy"`
	BucketVersioningConfiguration   BucketVersioningConfiguration   `json:"BucketVersioningConfiguration"`
	IsRequesterPaysEnabled          bool                            `json:"IsRequesterPaysEnabled"`
	PublicAccessBlockConfiguration  PublicAccessBlockConfiguration  `json:"PublicAccessBlockConfiguration"`
}
type Tags struct {
}
type ConfigurationItem struct {
	RelatedEvents                []interface{}              `json:"relatedEvents"`
	Relationships                []interface{}              `json:"relationships"`
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
	ResourceCreationTime         time.Time                  `json:"resourceCreationTime"`
}
