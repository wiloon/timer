FROM golang:1.13.6 AS build

ENV GO111MODULE on
WORKDIR /go/src/workdir

ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.io go build -a timer.go

FROM alpine AS prod

COPY --from=build /go/src/workdir/timer /data/
ADD config.toml /data/
CMD ["/data/timer"]
