FROM golang:latest AS builder

LABEL maintainer="darmiel <hi@d2a.io>"
LABEL org.opencontainers.image.source = "https://github.com/darmiel/yaxc"

WORKDIR /usr/src/app

# Install dependencies
# Thanks to @montanaflynn
# https://github.com/montanaflynn/golang-docker-cache
COPY go.mod go.sum ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

# Copy remaining source
COPY . .

# Build from sources
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o yaxc .

# Output Image
FROM alpine
COPY --from=builder /usr/src/app/yaxc .

# Copy assets
RUN mkdir /assets
COPY --from=builder /usr/src/app/assets/ /assets
RUN ls -larth /assets

ENTRYPOINT ["/yaxc"]