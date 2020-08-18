variable "aws_region" {}

variable "aws_account_id" {}

variable "resource_name_prefix" {
  description = "All the resources will be prefixed with this varible"
  default     = "in4it"
}

variable "s3_bucket" {
  description = "S3 bucket with lambda packages"
}

variable "rule_s3_public_buckets_enabled" {
  description = "Enable rule s3-public-buckets"
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
