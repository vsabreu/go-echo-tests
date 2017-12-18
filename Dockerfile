FROM golang:latest

ENV GOPATH $GOPATH
LABEL maintainer="Vinicius Abreu <vsabreu.dev@gmail.com>"
WORKDIR /go/src/github.com/vsabreu/go-echo-tests
ADD . $GOPATH/src/github.com/vsabreu/go-echo-tests
RUN ./installdeps.sh && env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

FROM alpine:latest

COPY --from=0 /go/src/github.com/vsabreu/go-echo-tests/go-echo-tests .
EXPOSE 8111
CMD ["./go-echo-tests"]
