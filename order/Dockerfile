FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 12201

CMD ["go", "run", "cmd/main.go"]