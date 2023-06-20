FROM golang:1.19.2-alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN go build -o /bin ./app/gateway/cmd/gateway

FROM alpine

WORKDIR /usr/src/app/

COPY --from=builder /bin ./
COPY --from=builder /src/configs ./

EXPOSE 8000

CMD [ "./gateway","-f","./config.yaml" ]