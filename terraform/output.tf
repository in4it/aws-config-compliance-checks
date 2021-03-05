output "sg_public_access_arn" {
  value = aws_config_config_rule.sg-public-access[0].arn
}

output "sg_public_access_egress_arn" {
  value = aws_config_config_rule.sg-public-access-egress[0].arn
}

output "s3_public_buckets_arn" {
  value = aws_config_config_rule.s3-public-buckets[0].arn
}

output "s3_vpc_traffic_only" {
  value = aws_config_config_rule.s3-vpc-traffic-only[0].arn
}

output "sg_public_access_id" {
  value = aws_config_config_rule.sg-public-access[0].id
}

output "sg_public_access_egress_id" {
  value = aws_config_config_rule.sg-public-access-egress[0].id
}

output "s3_public_buckets_id" {
  value = aws_config_config_rule.s3-public-buckets[0].id
}

output "check_permissions_boundaries_id" {
  value = aws_config_config_rule.check-permissions-boundaries[0].id
}
