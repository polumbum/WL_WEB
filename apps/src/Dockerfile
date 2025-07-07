FROM golang:1.22

WORKDIR /app_spa

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Установите swaggo
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Генерация документации Swagger
RUN swag init -d ./cmd/,./internal/http-server/handlers/,./internal/dto/

RUN go build -o main ./cmd/main.go 

VOLUME [ "/app_spa/log" ]

CMD ["./main"]