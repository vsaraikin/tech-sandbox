FROM golang:1.20

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY kafka .
COPY main.go .

CMD ["go", "run", "main.go"]
