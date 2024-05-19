FROM golang:1.22

WORKDIR /app

COPY . .
RUN go get -d -v ./...

RUN go build -v -o main .

CMD ["./main"]