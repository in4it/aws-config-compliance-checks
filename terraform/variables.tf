variable "resource_name_prefix" {
  description = "All the resources will be prefixed with this varible"
  default     = "in4it"
}

variable "s3_bucket" {
  description = "S3 bucket with lambda packages"
}

variable "s3_bucket_kms_key_arn" {
  description = "S3 bucket key with lambda packages"
  default = ""
}

variable "rule_s3_public_buckets_enabled" {
  description = "Enable rule s3-public-buckets"
  default = true
}

variable "rule_s3_lifecycle_enabled" {
  description = "Enable rule s3-lifecycle"
  default = true
}

variable "rule_s3_vpc_traffic_only_enabled" {
  description = "Enable rule s3-vpc-traffic-only"
  default = true
}

variable "rule_sg_public_access_enabled" {
  description = "Enable rule sg-public-access"
  default = true
}

variable "rule_sg_public_access_egress_enabled" {
  description = "Enable rule sg-public-access-egress"
  default = true
}

variable "exclude_buckets" {
  description = "Skip compliance check for listed buckets"
  default = ""
}

variable "exclude_security_groups_ingress" {
  description = "Skip ingress compliance check for listed Security Groups"
  default = ""
}

variable "exclude_security_groups_egress" {
  description = "Skip ingress compliance check for listed Security Groups"
  default = ""
}

variable "cloudwatch_log_retention_period" {
  description = "cloudwatch retention period in days"
  default     = "0"
}