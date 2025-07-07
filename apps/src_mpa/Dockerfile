FROM golang:1.22

WORKDIR /app_mpa

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go 

VOLUME [ "/app_mpa/log" ]

CMD ["./main"]