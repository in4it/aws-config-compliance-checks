GOARCH = amd64

all: clean build
clean:
	rm -rf *.zip

build: s3-public-buckets sg-public-access sg-public-access-ergess s3-vpc-traffic-only

s3-public-buckets:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o s3-public-buckets cmd/s3-public-buckets/*.go
	zip s3-public-buckets.zip s3-public-buckets
	rm -rf s3-public-buckets

s3-lifecycle:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o s3-lifecycle cmd/s3-lifecycle/*.go
	zip s3-lifecycle.zip s3-lifecycle
	rm -rf s3-lifecycle

sg-public-access:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o sg-public-access cmd/sg-public-access/*.go
	zip sg-public-access.zip sg-public-access
	rm -rf sg-public-access

sg-public-access-ergess:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o sg-public-access-egress cmd/sg-public-access-egress/*.go
	zip sg-public-access-egress.zip sg-public-access-egress
	rm -rf sg-public-access-egress

s3-vpc-traffic-only:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o s3-vpc-traffic-only cmd/s3-vpc-traffic-only/*.go
	zip s3-vpc-traffic-only.zip s3-vpc-traffic-only
	rm -rf s3-vpc-traffic-only