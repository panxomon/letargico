FROM golang:latest

WORKDIR /letargico

COPY . .

RUN go build -o /go/bin/letargico server/main.go

CMD ["webgo"]