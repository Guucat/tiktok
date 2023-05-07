FROM golang:1.20-alpine AS builder

RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /app

WORKDIR /app/server/comment/
RUN CGO_ENABLED=0 GOOS=linux go build -o comment .

FROM alpine:3.13
COPY --from=builder /app/server/comment/comment /comment

EXPOSE 7010

ENTRYPOINT ["/comment"]