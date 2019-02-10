FROM golang:1.11-alpine as build-env

WORKDIR /go/src/github.com/chrisstowe/forgestatus

COPY common ./common
COPY forgestatus-server ./forgestatus-server

RUN go build -o ./goApp ./forgestatus-server

FROM alpine

WORKDIR /opt/app

# install app
COPY --from=build-env /go/src/github.com/chrisstowe/forgestatus/goApp .

ENTRYPOINT [ "./goApp" ]
