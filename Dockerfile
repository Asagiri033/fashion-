FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o fashion ./main.go

ENTRYPOINT ["./fashion"]
