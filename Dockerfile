FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /name-ddns ./cmd/ddns


FROM debian:stretch-slim

WORKDIR /

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=build /name-ddns /bin/name-ddns

CMD ["/bin/name-ddns"]
