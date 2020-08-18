resource "aws_iam_role" "iam_for_lambda" {
  name = "custom_config_lambda"

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

resource "aws_iam_policy" "lambda_role" {
  name        = "config_lambda_role"
  path        = "/"
  description = "IAM policy for logging and config from a lambda"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "logs:CreateLogGroup",
            "Resource": "arn:aws:logs:${var.aws_region}:${var.aws_account_id}:*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "logs:CreateLogStream",
                "logs:PutLogEvents"
            ],
            "Resource": [
                "arn:aws:logs:${var.aws_region}:${var.aws_account_id}:log-group:*:*"
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

resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = aws_iam_policy.lambda_role.arn
}

resource "aws_lambda_function" "in4it-sg-public-access" {
  s3_bucket     = var.s3_bucket
  s3_key        = "lambdas/sg-public-access.zip"
  function_name = "${var.resource_name_prefix}-sg-public-access"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "sg-public-access"

  runtime = "go1.x"

}

resource "aws_lambda_function" "in4it-s3-public-buckets" {
  s3_bucket     = var.s3_bucket
  s3_key        = "lambdas/s3-public-buckets.zip"
  function_name = "${var.resource_name_prefix}-s3-public-buckets"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "s3-public-buckets"

  runtime = "go1.x"

}

resource "aws_lambda_function" "in4it-sg-public-access-egress" {
  s3_bucket     = var.s3_bucket
  s3_key        = "lambdas/sg-public-access-egress.zip"
  function_name = "${var.resource_name_prefix}-sg-public-access-egress"
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = "sg-public-access-egress"

  runtime = "go1.x"

}

resource "aws_lambda_permission" "in4it-sg-public-access" {
  statement_id   = "AllowConfigToInvoke"
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.in4it-sg-public-access.function_name
  principal      = "config.amazonaws.com"
  source_account = var.aws_account_id
}

resource "aws_lambda_permission" "in4it-sg-public-access-egress" {
  statement_id   = "AllowConfigToInvoke"
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.in4it-sg-public-access-egress.function_name
  principal      = "config.amazonaws.com"
  source_account = var.aws_account_id
}

resource "aws_lambda_permission" "in4it-s3-public-buckets" {
  statement_id   = "AllowConfigToInvoke"
  action         = "lambda:InvokeFunction"
  function_name  = aws_lambda_function.in4it-s3-public-buckets.function_name
  principal      = "config.amazonaws.com"
  source_account = var.aws_account_id
}

