FROM golang:1.20-alpine AS builder

RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /app

WORKDIR /app/server/route/
RUN CGO_ENABLED=0 GOOS=linux go build -o route .

FROM alpine:3.13
COPY --from=builder /app/server/route/route /route

EXPOSE 7070

ENTRYPOINT ["/route"]