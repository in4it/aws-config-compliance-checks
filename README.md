## IN4IT AWS CONFIG CUSTOM RULES

### s3-public-buckets
Checks if "Block all public access" is set for s3 buckets

Params:
- excludeBuckets: Skip compliance check for listed buckets
    - key: excludeBuckets
    - value: Comma separated list of buckets, e.g "bucket1, bucket2"

### s3-lifecycle
Checks if "S3 Lifecycle configuration" is set for s3 buckets

Params:
- excludeBuckets: Skip compliance check for listed buckets
    - key: excludeBuckets
    - value: Comma separated list of buckets, e.g "bucket1, bucket2"


### s3-vpc-traffic-only
S3 VPC Traffic Only checks if there is a deny rule present in the S3 policy to block non-VPC traffic

Example policy check: 

```
{
    "Version": "2012-10-17",
    "Id": "S3-test-bucket-policy",
    "Statement": [
        {
            "Sid": "Access-to-specific-VPC-only",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::123456789:root"
            },
            "Action": "s3:*",
            "Resource": [
                "arn:aws:s3:::test-bucket/*",
                "arn:aws:s3:::test-bucket"
            ],
            "Condition": {
                "StringEquals": {
                    "aws:sourceVpc": [
                        "vpc-0e123454566755667",
                        "vpc-0b3k3333443838838"
                    ]
                }
            }
        }
    ]
}
```

Params:
- excludeBuckets: Skip compliance check for listed buckets
    - key: excludeBuckets
    - value: Comma separated list of buckets, e.g "bucket1, bucket2"

### sg-public-access
Checks AWS security groups for rules that allow access from "0.0.0.0/0". A parameter can be added to exclude security groups in the format sg-12345:80+443, sg-45678. The first excludes only specific ports in a security group, the latter excludes the whole security group from compliance checks.

Params:
- excludeSecurityGroups: Skip compliance check for listed Security Groups
    - key: excludeSecurityGroups
    - value: Comma separated list of Security Group Ids, optionally followed by port numbers e.g:  "sg-12345:80+443, sg-67890, sg-112233:22"

### sg-public-access-egress
Checks AWS security groups for rules that allow egress access to "0.0.0.0/0". A parameter can be added to exclude security groups in the format sg-12345:80+443, sg-45678. The first excludes only specific ports in a security group, the latter excludes the whole security group from compliance checks.

Params:
- excludeSecurityGroups: Skip compliance check for listed Security Groups
    - key: excludeSecurityGroups
    - value: Comma separated list of Security Group Ids, optionally followed by port numbers e.g:  "sg-12345:80+443, sg-67890, sg-112233:22"



### Deployment

Build config rules:
```make all```
Upload config rules zipped binaries to your s3 bucket under the `/lambdas` key.
```aws s3 cp s3-public-access.zip s3://my-bucket/lambdas/```

Include terraform module and apply.
Example setup:
```
module "in4it-config-rules" {
    source = "git@github.com:in4it/aws-config-compliance-checks.git//terraform"
    s3_bucket = "my-bucket"
    exclude_buckets = "bucket1"
    exclude_security_groups_ingress = "sg-lb-1:80+443,sg-lb-2"
}
```
