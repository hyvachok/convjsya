FROM golang:alpine AS build

RUN apk --no-cache add git

ADD . "$GOPATH/src/github.com/hyvachok/convjsya"

RUN go get github.com/ghodss/yaml
RUN go install github.com/hyvachok/convjsya

FROM alpine:latest

WORKDIR /bin/

COPY --from=build /go/bin/convjsya .

ENTRYPOINT ["/bin/convjsya"]