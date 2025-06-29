FROM golang:1.22

WORKDIR /usr/app/src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8090

CMD ["go", "run", "./cmd/main.go"]
