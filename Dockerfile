# -----------------------------------------------------------------------------
# The base image for building the binary

FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod go.sum Makefile ./
COPY internal internal
COPY cmd cmd
RUN BUILD_PATH=/ make build

# -----------------------------------------------------------------------------
# Build the final Docker image

FROM debian:bullseye-slim

WORKDIR /

RUN apt-get update \
	&& apt-get install -y ca-certificates \
	&& apt-get clean \
	&& rm -rf /var/lib/apt/lists/*

COPY --from=build /name-ddns /bin/name-ddns

CMD ["/bin/name-ddns"]
