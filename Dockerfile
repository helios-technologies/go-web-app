FROM golang

COPY . /go/src/github.com/helios-technologies/go-web-app

RUN go install github.com/helios-technologies/go-web-app

ENTRYPOINT /go/bin/go-web-app

EXPOSE 8080
