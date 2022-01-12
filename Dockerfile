FROM golang:1.17-alpine

WORKDIR /app

RUN go version

COPY . .

RUN go build -o /main cmd/main.go

ENTRYPOINT [ "/main" ]