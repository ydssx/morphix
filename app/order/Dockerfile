FROM golang:1.21.4-alpine3.18 AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /bin ./app/order/cmd/order

FROM alpine


RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Shanghai

WORKDIR /usr/src/app/

COPY --from=builder /bin ./

EXPOSE 9003

CMD [ "./order","-f","/etc/morphix/config.yaml" ]