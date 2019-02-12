FROM golang:1.11-alpine as build-env
RUN apk add git

WORKDIR /go/src/github.com/chrisstowe/forgestatus

COPY common ./common
COPY server ./server

RUN go get -d ./server

RUN go build -o ./serverApp ./server

FROM alpine

WORKDIR /opt/app

# install app
COPY --from=build-env /go/src/github.com/chrisstowe/forgestatus/serverApp .

EXPOSE 80

ENTRYPOINT [ "./serverApp" ]
