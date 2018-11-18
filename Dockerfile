FROM golang:alpine AS builder
LABEL maintainer="Erin Schnabel <erinschnabel@gmail.com> (@ebullientworks)"

WORKDIR /go/src
RUN apk add bash && ls /go/src
COPY docker/build.sh /usr/local/bin/build.sh
COPY ./src /go/src/advent
RUN cd /go/src && build.sh && find /go


FROM alpine
LABEL maintainer="Erin Schnabel <erinschnabel@gmail.com> (@ebullientworks)"
WORKDIR /app
COPY --from=builder /go/bin /app/

ENTRYPOINT ./advent
