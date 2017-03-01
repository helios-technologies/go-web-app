FROM golang

COPY . /go/src/github.com/helios-technologies/go-web-app

ARG VERSION
RUN go install -ldflags "-X main.Version=${VERSION}" github.com/helios-technologies/go-web-app

ENTRYPOINT /go/bin/go-web-app

EXPOSE 8080
