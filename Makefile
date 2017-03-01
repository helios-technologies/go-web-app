TAG ?= onbuild
LDFLAGS += -X main.Version=$(TAG)

build:
	go build -ldflags "$(LDFLAGS)" .

image:
	docker build --build-arg VERSION=$(TAG) -t heliostech/go-web-app:$(TAG) .
