TAG ?= latest

image:
	docker build -t heliostech/go-web-app:$(TAG) .
