FROM golang:1.19.2-alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /bin ./app/order/cmd/order

FROM alpine

ENV TZ=Asia/Shanghai

RUN apk add --no-cache ca-certificates

WORKDIR /usr/src/app/

COPY --from=builder /bin ./
COPY --from=builder /src/configs ./

EXPOSE 9003

CMD [ "./order","-f","/etc/morphix/config.yaml" ]