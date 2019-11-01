FROM golang:1.13.3 AS build

ENV GO111MODULE on
WORKDIR /go/src/timer

ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://goproxy.io go build -a timer.go

FROM alpine AS prod

COPY --from=build /go/src/timer/timer /data/timer/
CMD ["/data/timer/timer"]
