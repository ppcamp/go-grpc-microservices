FROM golang:1.18-alpine as builder

RUN mkdir /build
WORKDIR /build

ADD src /build/

RUN apk update && \
    apk upgrade && \
    apk add --no-cache musl-dev gcc build-base libc-dev ca-certificates wget

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -race -a -o /build/application cmd/*.go


# ------------------------------------------------------------------------------


FROM alpine:3.14

RUN apk update && \
    apk upgrade && \
    apk add --no-cache && \
    apk add ca-certificates && \
    tzdata ca-certificates && \
    cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime; \
    echo "America/Sao_Paulo" > /etc/timezone

COPY --from=builder /build/application .

# executable
ENTRYPOINT [ "./application" ]

# arguments that can be overridden
CMD [ "./application" ]