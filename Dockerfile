# golang:1.24.1-alpine3.21
FROM golang@sha256:c3c72c53c324f5c4b48b3918bf98351b0003216710717233d5a7aae539441e48 AS builder

# environment variables
ENV CGO_ENABLED=0 \
    GOOS=linux    \
    GOARCH=amd64  \
    GO111MODULE=on

WORKDIR /app

# copy source code
COPY . .

# download dependencies with cache mount for better performance
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download && \
    go mod verify

# -ldflags optimize binary size
# -trimpath remove absolute path from binary
RUN go build -ldflags="-s -w -extldflags '-static'" -trimpath -o service ./cmd/service/main.go
RUN go build -ldflags="-s -w -extldflags '-static'" -trimpath -o cli ./cmd/cli/main.go

# final image
# alpine:3.21
FROM alpine@sha256:a8560b36e8b8210634f77d9f7f9efd7ffa463e380b75e2e74aff4511df3ef88c

LABEL maintainer="Javier Aguilera"

ARG USER=app
ARG UID=10001
ARG GID=10001
ENV HOME=/home/${USER}

RUN apk update && apk upgrade && \
    # ca-certificates is needed for SSL/TLS connections and libcap for capabilities
    apk add --no-cache ca-certificates=20241121-r1 libcap=2.71-r0 && \
    rm -rf /var/cache/apk/* && \
    # create non-root user and group
    addgroup -g ${GID} ${USER} && \
    adduser -D -g "" -h "${HOME}" -s "/sbin/nologin" -G ${USER} -u ${UID} ${USER} && \
    # setup app directory and permissions
    mkdir -p /app && \
    chown -R ${UID}:${GID} /app

WORKDIR /app

# copy binaries from builder
COPY --from=builder --chown=${UID}:${GID} /app/server .
COPY --from=builder --chown=${UID}:${GID} /app/cli .

# drop root privileges
USER ${USER}
