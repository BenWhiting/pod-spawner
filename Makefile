GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test -v
LDFLAGS=-ldflags="-s -w"
GOBUILDFLAGS=-a $(LDFLAGS)
BUILDOPTIONS=-a -installsuffix nocgo
BUILDDIR=build
GOARCH=amd64

all:  build

apiserver:
	CGO_ENABLE=0 GOOS=linux GOARCH=$(GOARCH) $(GOBUILD) $(GOBUILDFLAGS) $(BUILDOPTIONS) \
	-o $(BUILDDIR)/apiserver cmd/apiserver/server.go

docker:
	mv ./build/apiserver ./DockerFiles/apiserver/apiserver
	docker build ./DockerFiles/apiserver/
	rm ./DockerFiles/apiserver/apiserver