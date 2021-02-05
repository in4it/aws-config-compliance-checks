data "aws_region" "current" {}
data "aws_caller_identity" "current" {}

resource "aws_iam_role" "sg-public-access-egress" {
  name = "sg-public-access-egress"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role" "config-s3-lifecycle" {
  name = "config-s3-lifecycle"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role" "s3-public-buckets" {
  name = "s3-public-buckets"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}


resource "aws_iam_role" "s3-vpc-traffic-only" {
  name = "s3-vpc-traffic-only"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}


resource "aws_iam_role" "sg-public-access" {
  name = "sg-public-access"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "sg-public-access-egress" {
  name        = "sg-public-access-egress"
  path        = "/"
  description = "IAM policy for logging and config from a lambda"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-sg-public-access-egress"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-sg-public-access-egress:log-stream:*"
            ]
        },
        {
            "Sid": "putEvaluations",
            "Effect": "Allow",
            "Action": [
                "config:PutEvaluations"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_policy" "s3-public-buckets" {
  name        = "s3-public-buckets"
  path        = "/"
  description = "IAM policy for logging and config from a lambda"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-s3-public-buckets"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-s3-public-buckets:log-stream:*"
            ]
        },
        {
            "Sid": "putEvaluations",
            "Effect": "Allow",
            "Action": [
                "config:PutEvaluations"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_policy" "config-s3-lifecycle" {
  name        = "config-s3-lifecycle"
  path        = "/"
  description = "IAM policy for logging and config from a lambda"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-config-s3-lifecycle"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-config-s3-lifecycle:log-stream:*"
            ]
        },
        {
            "Sid": "putEvaluations",
            "Effect": "Allow",
            "Action": [
                "config:PutEvaluations"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}


resource "aws_iam_policy" "" {
  name        = "s3-public-buckets"
  path        = "/"
  description = "IAM policy for logging and config from a lambda"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-s3-public-buckets"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-s3-public-buckets:log-stream:*"
            ]
        },
        {
            "Sid": "putEvaluations",
            "Effect": "Allow",
            "Action": [
                "config:PutEvaluations"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}





resource "aws_iam_policy" "sg-public-access" {
  name        = "sg-public-access"
  path        = "/"
  description = "IAM policy for logging and config from a lambda"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-sg-public-access"
            ]
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
              "arn:aws:logs:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:log-group:/aws/lambda/${var.resource_name_prefix}-sg-public-access:*"
            ]
        },
        {
            "Sid": "putEvaluations",
            "Effect": "Allow",
            "Action": [
                "config:PutEvaluations"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "sg-public-access-egress" {
  role       = aws_iam_role.sg-public-access-egress.name
  policy_arn = aws_iam_policy.sg-public-access-egress.arn
}

resource "aws_iam_role_policy_attachment" "sg-public-access" {
  role       = aws_iam_role.sg-public-access.name
  policy_arn = aws_iam_policy.sg-public-access.arn
}

resource "aws_iam_role_policy_attachment" "s3-public-buckets" {
  role       = aws_iam_role.s3-public-buckets.name
  policy_arn = aws_iam_policy.s3-public-buckets.arn
}

resource "aws_iam_role_policy_attachment" "config-s3-lifecycle" {
  role       = aws_iam_role.config-s3-lifecycle.name
  policy_arn = aws_iam_policy.config-s3-lifecycle.arn
}

resource "aws_lambda_function" "sg-public-access" {
  s3_bucket     = var.s3_bucket
  kms_key_arn   = var.s3_bucket_kms_key_arn
  s3_key        = "lambdas/sg-public-access.zip"
  function_name = "${var.resource_name_prefix}-sg-public-access"
  role          = aws_iam_role.sg-public-access.arn
  handler       = "sg-public-access"

  runtime = "go1.x"

}

resource "aws_lambda_function" "s3-public-buckets" {
  s3_bucket     = var.s3_bucket
  kms_key_arn   = var.s3_bucket_kms_key_arn
  s3_key        = "lambdas/s3-public-buckets.zip"
  function_name = "${var.resource_name_prefix}-s3-public-buckets"
  role          = aws_iam_role.s3-public-buckets.arn
  handler       = "s3-public-buckets"

  runtime = "go1.x"

}

resource "aws_lambda_function" "s3-lifecycle" {
  s3_bucket     = var.s3_bucket
  kms_key_arn   = var.s3_bucket_kms_key_arn
  s3_key        = "lambdas/s3-lifecycle.zip"
  function_name = "${var.resource_name_prefix}-s3-lifecycle"
  role          = aws_iam_role.config-s3-lifecycle.arn
  handler       = "s3-lifecycle"

  runtime = "go1.x"

}

resource "aws_lambda_function" "sg-public-access-egress" {
  s3_bucket     = var.s3_bucket
  kms_key_arn   = var.s3_bucket_kms_key_arn
  s3_key        = "lambdas/sg-public-access-egress.zip"
  function_name = "${var.resource_name_prefix}-sg-public-access-egress"
  role          = aws_iam_role.sg-public-access-egress.arn
  handler       = "sg-public-access-egress"

  runtime = "go1.x"

}

resource "aws_lambda_permission" "sg-public-access" {
  statement_id   = "AllowConfigToInvoke"
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.sg-public-access.function_name
  principal      = "config.amazonaws.com"
  source_account = data.aws_caller_identity.current.account_id
}

resource "aws_lambda_permission" "sg-public-access-egress" {
  statement_id   = "AllowConfigToInvoke"
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.sg-public-access-egress.function_name
  principal      = "config.amazonaws.com"
  source_account = data.aws_caller_identity.current.account_id
}

resource "aws_lambda_permission" "s3-public-buckets" {
  statement_id   = "AllowConfigToInvoke"
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.s3-public-buckets.function_name
  principal      = "config.amazonaws.com"
  source_account = data.aws_caller_identity.current.account_id
}

resource "aws_lambda_permission" "config-s3-lifecycle" {
  statement_id   = "AllowConfigToInvoke"
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.s3-lifecycle.function_name
  principal      = "config.amazonaws.com"
  source_account = data.aws_caller_identity.current.account_id
}

