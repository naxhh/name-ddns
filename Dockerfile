FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
#COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /name-ddns ./cmd/ddns

FROM debian:bullseye-slim

ENV NAME_DDNS_UPDATE_CRON */10 * * * *

WORKDIR /

RUN apt-get update && apt-get install -y ca-certificates cron

COPY --from=build /name-ddns /bin/name-ddns
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh
RUN touch /cron.log

ENTRYPOINT ["/entrypoint.sh"]
