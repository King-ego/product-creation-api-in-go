FROM golang:1.22

WORKDIR /usr/app/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8090

RUN go build -o main cmd/main.go

CMD ["./main"]
