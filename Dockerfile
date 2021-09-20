#################
# Base image
#################
FROM alpine:3.12 as echo-mongodb-base

USER root

RUN addgroup -g 10001 echo-mongodb && \
    adduser --disabled-password --system --gecos "" --home "/home/echo-mongodb" --shell "/sbin/nologin" --uid 10001 echo-mongodb && \
    mkdir -p "/home/echo-mongodb" && \
    chown echo-mongodb:0 /home/echo-mongodb && \
    chmod g=u /home/echo-mongodb && \
    chmod g=u /etc/passwd

ENV USER=echo-mongodb
USER 10001
WORKDIR /home/echo-mongodb

#################
# Builder image
#################
FROM golang:1.16-alpine AS echo-mongodb-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build

#################
# Final image
#################
FROM echo-mongodb-base

COPY --from=echo-mongodb-builder /app/bin/echo-mongodb /usr/local/bin

# Command to run the executable
ENTRYPOINT ["echo-mongodb"]
