FROM golang:1.12

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
