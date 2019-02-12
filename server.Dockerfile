FROM golang:1.11-alpine as build-env

# These would normally be factored out into a base build image.
RUN apk add make
RUN apk add git

WORKDIR /go/src/github.com/chrisstowe/forgestatus

COPY Makefile .
COPY common ./common
COPY server ./server

RUN make build-server

FROM alpine

WORKDIR /opt/app

# Install app.
COPY --from=build-env /go/src/github.com/chrisstowe/forgestatus/build/server .

EXPOSE 80

ENTRYPOINT [ "./server" ]
