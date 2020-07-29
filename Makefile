GOARCH = amd64

all: clean build 
clean:
	rm -rf *.zip

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o s3-public-buckets cmd/s3-public-buckets/*.go
