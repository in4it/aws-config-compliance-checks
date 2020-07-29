## IN4IT AWS CONFIG CUSTOM RULES

### s3-public-buckets
Checks if "Block all public access" is set for s3 buckets

Params:
- excludeBuckets: Skip compliance check for listed buckets
    - key: excludeBuckets
    - value: Comma separated list of buckets, e.g "bucket1, bucket2"
