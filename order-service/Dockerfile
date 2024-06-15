FROM golang:1.16

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

EXPOSE 8083

CMD ["./main"]

