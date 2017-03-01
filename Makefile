TAG ?= onbuild
LDFLAGS += -X main.Version=$(TAG)

build:
	go build -ldflags "$(LDFLAGS)" .

image:
	docker build -t heliostech/go-web-app:$(TAG) .
