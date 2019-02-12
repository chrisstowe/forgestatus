FROM golang:1.11-alpine as build-env

WORKDIR /go/src/github.com/chrisstowe/forgestatus

COPY common ./common
COPY worker ./worker

RUN go build -o ./workerApp ./worker

FROM alpine

WORKDIR /opt/app

# install app
COPY --from=build-env /go/src/github.com/chrisstowe/forgestatus/workerApp .

EXPOSE 80

ENTRYPOINT [ "./workerApp" ]
