output "sg_public_access_arn" {
  value = aws_config_config_rule.in4it-sg-public-access[0].arn
}

output "sg_public_access_egress_arn" {
  value = aws_config_config_rule.in4it-sg-public-access-egress[0].arn
}

output "s3_public_buckets_arn" {
  value = aws_config_config_rule.in4it-s3-public-buckets[0].arn
}

output "sg_public_access_id" {
  value = aws_config_config_rule.in4it-sg-public-access[0].id
}

output "sg_public_access_egress_id" {
  value = aws_config_config_rule.in4it-sg-public-access-egress[0].id
}

output "s3_public_buckets_id" {
  value = aws_config_config_rule.in4it-s3-public-buckets[0].id
}
