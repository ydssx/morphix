FROM golang:1.19.2-alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN go build -o /bin ./app/user/cmd/user

FROM alpine

ENV TZ=Asia/Shanghai

RUN apk add --no-cache ca-certificates

WORKDIR /usr/src/app/

COPY --from=builder /bin ./
COPY --from=builder /src/configs ./

EXPOSE 9000

CMD [ "./user","-f","/etc/morphix/config.yaml" ]