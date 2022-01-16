FROM golang:1.17-alpine

WORKDIR /app

RUN go version

COPY . .

RUN go build -o /server cmd/server/main.go && go build -o /client cmd/client/main.go

ENTRYPOINT [ "/server" ]