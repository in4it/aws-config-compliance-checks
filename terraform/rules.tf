resource "aws_config_config_rule" "in4it-s3-public-buckets" {
  count = var.rule_s3_public_buckets_enabled ? 1 : 0
  description = "Checks if \"Block all public access\" is set for s3 buckets"
  input_parameters = jsonencode(
    {
      excludeBuckets = var.exclude_buckets
    }
  )
  name = "${var.resource_name_prefix}-s3-public-buckets"

  scope {
    compliance_resource_types = [
      "AWS::S3::Bucket",
    ]
  }

  source {
    owner             = "CUSTOM_LAMBDA"
    source_identifier = "${aws_lambda_function.in4it-s3-public-buckets.arn}"

    source_detail {
      event_source = "aws.config"
      message_type = "ConfigurationItemChangeNotification"
    }
    source_detail {
      event_source = "aws.config"
      message_type = "OversizedConfigurationItemChangeNotification"
    }
  }
}

resource "aws_config_config_rule" "in4it-sg-public-access" {
  count = var.rule_sg_public_access_enabled ? 1 : 0
  description = "Checks AWS security groups for rules that allow access from \"0.0.0.0/0\""
  name        = "${var.resource_name_prefix}-sg-public-access"

  input_parameters = jsonencode(
    {
      excludeSecurityGroups = var.exclude_security_groups_ingress
    }
  )

  scope {
    compliance_resource_types = [
      "AWS::EC2::SecurityGroup",
    ]
  }

  source {
    owner             = "CUSTOM_LAMBDA"
    source_identifier = "${aws_lambda_function.in4it-sg-public-access.arn}"

    source_detail {
      event_source = "aws.config"
      message_type = "ConfigurationItemChangeNotification"
    }
    source_detail {
      event_source = "aws.config"
      message_type = "OversizedConfigurationItemChangeNotification"
    }
  }
}

resource "aws_config_config_rule" "in4it-sg-public-access-egress" {
  count = var.rule_sg_public_access_egress_enabled ? 1 : 0
  description = "Checks AWS security groups for rules that allow egress access to \"0.0.0.0/0\""
  name        = "${var.resource_name_prefix}-sg-public-access-egress"

  input_parameters = jsonencode(
    {
      excludeSecurityGroups = var.exclude_security_groups_egress
    }
  )

  scope {
    compliance_resource_types = [
      "AWS::EC2::SecurityGroup",
    ]
  }

  source {
    owner             = "CUSTOM_LAMBDA"
    source_identifier = "${aws_lambda_function.in4it-sg-public-access-egress.arn}"

    source_detail {
      event_source = "aws.config"
      message_type = "ConfigurationItemChangeNotification"
    }
    source_detail {
      event_source = "aws.config"
      message_type = "OversizedConfigurationItemChangeNotification"
    }
  }
}
