GOARCH = amd64

all: clean build 
clean:
	rm -rf *.zip

s3-public-buckets:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o s3-public-buckets cmd/s3-public-buckets/*.go

sg-public-access:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o sg-public-access cmd/sg-public-access/*.go
