FROM golang:1.20-alpine AS builder

RUN go env -w GO111MODULE=auto
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /app

WORKDIR /app/server/user/
RUN CGO_ENABLED=0 GOOS=linux go build -o user .

FROM alpine:3.13
COPY --from=builder /app/server/user/user /user

EXPOSE 7030

ENTRYPOINT ["/user"]