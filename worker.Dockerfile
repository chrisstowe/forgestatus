# This is the build environment.
# Only specify commands here needed for building the app.
FROM golang:1.11-alpine as build-env

# These would normally be factored out into a base build image.
RUN apk update && apk add \
    make \
    git

WORKDIR /go/src/github.com/chrisstowe/forgestatus

COPY Makefile .
COPY common ./common
COPY worker ./worker

RUN make build-worker

# This is the runtime environment.
# Only specify commands here needed for running the app.
FROM alpine

# Set the work dir and user for better security.
WORKDIR /opt/app
RUN addgroup -S appgroup && \
    adduser -S appuser -G appgroup
USER appuser

# Install the built app.
COPY --from=build-env /go/src/github.com/chrisstowe/forgestatus/build/worker .

EXPOSE 80

ENTRYPOINT [ "./worker" ]
