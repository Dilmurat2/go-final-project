FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8081

CMD ["go", "run", "cmd/main.go"]