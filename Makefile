BUILD_PATH := deployment/docker
TAG := jdxj/hello:v0.2.0

.PHONY: all
all: build docker

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(BUILD_PATH)/main.run *.go

.PHONY: docker
docker:
	docker build -t $(TAG) $(BUILD_PATH)
	docker push $(TAG)

.PHONY: clean
clean:
	@rm -rf $(BUILD_PATH)/main.run
