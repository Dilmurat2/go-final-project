FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 50052

CMD ["go", "run", "cmd/main.go"]