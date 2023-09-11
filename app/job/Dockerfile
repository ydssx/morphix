FROM golang:1.19.2-alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /bin ./app/job/cmd/job

FROM alpine

ENV TZ=Asia/Shanghai

RUN apk add --no-cache ca-certificates

WORKDIR /usr/src/app/

COPY --from=builder /bin ./
COPY --from=builder /src/configs ./


CMD [ "./job", "-f", "/etc/morphix/config.yaml" ]