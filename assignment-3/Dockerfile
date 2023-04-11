FROM golang:1.20

WORKDIR /app

COPY ./ ./

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]