FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o fenlie

EXPOSE 8080

ENTRYPOINT ["/app/fenlie"]