FROM golang:1.19.2-alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /bin ./app/payment/cmd/payment

FROM alpine

WORKDIR /usr/src/app/

COPY --from=builder /bin ./
COPY --from=builder /src/configs ./

EXPOSE 9002

CMD [ "./payment","-f","/etc/morphix/config.yaml" ]