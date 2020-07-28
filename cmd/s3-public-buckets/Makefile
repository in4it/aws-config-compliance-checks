GOARCH = amd64

all: clean build-linux zip

clean:
	cd bin && rm -rf check_s3 check_s3.zip

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o bin/check_s3 check_s3.go

zip:
	cd bin && zip check_s3.zip check_s3
