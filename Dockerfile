FROM golang:latest

WORKDIR /app

COPY Pr .

RUN go mod download

RUN go build -o main.go

CMD ["./main.go"]