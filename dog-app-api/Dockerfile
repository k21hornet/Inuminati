FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

EXPOSE 80

CMD ["sh", "-c", "GO_ENV=production go run migrate/migrate.go && ./app"]