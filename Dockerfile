# Dockerfile build a docker image with the binary of the application defined by CMD_NAME
# CMD_NAME is a build argument, it is used to copy the cmd folder.
# The resulting image is based on distroless image, which is a minimal image based on debian.
# Build: docker buildx build --build-arg CMD_NAME=./cmd/<cmd_name> -f <path-to-Dockerfile> .
ARG CMD_NAME

FROM golang:1.22-bullseye as dependencies

WORKDIR /usr/src/app

RUN apt-get install -y ca-certificates

COPY go.mod go.sum Makefile ./
RUN make dep

FROM dependencies as builder

ARG CMD_NAME

COPY cmd/${CMD_NAME} cmd/${CMD_NAME}
COPY pkg pkg
COPY internal internal
RUN make build.${CMD_NAME}

# Rename binary to app
RUN mv bin/$(ls ./bin) bin/app

FROM gcr.io/distroless/base-debian11 AS distroless

ARG ROOT_PATH=/usr/src/app

COPY --from=builder ${ROOT_PATH}/bin/app app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./app" ]