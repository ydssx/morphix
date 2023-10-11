FROM golang:1.19.2-alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /bin ./app/{{.appName}}/cmd/{{.appName}}

FROM alpine

RUN apk add --no-cache ca-certificates tzdata

ENV TZ=Asia/Shanghai

WORKDIR /usr/src/app/

COPY --from=builder /bin ./

EXPOSE {{.port}}

CMD [ "./{{.appName}}","-f","/etc/morphix/config.yaml" ]