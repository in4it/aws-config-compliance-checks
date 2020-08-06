## IN4IT AWS CONFIG CUSTOM RULES

### s3-public-buckets
Checks if "Block all public access" is set for s3 buckets

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

