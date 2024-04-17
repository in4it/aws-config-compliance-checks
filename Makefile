GOARCH = arm64

all: clean build
clean:
	rm -rf *.zip

build: s3-public-buckets s3-lifecycle sg-public-access sg-public-access-egress s3-vpc-traffic-only

s3-public-buckets:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -tags lambda.norpc -o bootstrap cmd/s3-public-buckets/*.go
	zip s3-public-buckets.zip bootstrap
	rm -rf bootstrap

s3-lifecycle:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -tags lambda.norpc -o bootstrap cmd/s3-lifecycle/*.go
	zip s3-lifecycle.zip bootstrap
	rm -rf bootstrap

sg-public-access:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -tags lambda.norpc -o bootstrap cmd/sg-public-access/*.go
	zip sg-public-access.zip bootstrap
	rm -rf bootstrap

sg-public-access-egress:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -tags lambda.norpc -o bootstrap cmd/sg-public-access-egress/*.go
	zip sg-public-access-egress.zip bootstrap
	rm -rf bootstrap

s3-vpc-traffic-only:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -tags lambda.norpc -o bootstrap cmd/s3-vpc-traffic-only/*.go
	zip s3-vpc-traffic-only.zip bootstrap
	rm -rf bootstrap