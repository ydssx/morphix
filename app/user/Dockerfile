FROM golang:1.21.4-alpine3.18 AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN go build -o /bin ./app/user/cmd/user

FROM alpine

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Shanghai

WORKDIR /usr/src/app/

COPY --from=builder /bin ./

EXPOSE 9000

CMD [ "./user","-f","/etc/morphix/config.yaml" ]