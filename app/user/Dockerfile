FROM golang:1.19.2-alpine AS builder

ENV GOPROXY=https://goproxy.cn

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user ./app/user/cmd/user

FROM alpine

WORKDIR /usr/src/app/

COPY --from=builder /src/user ./
COPY --from=builder /src/configs ./

EXPOSE 9000

CMD [ "./user","-f","./config.yaml" ]