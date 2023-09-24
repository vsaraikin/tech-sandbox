FROM golang:1.20

WORKDIR /app

COPY app/go.mod .
COPY app/go.sum .

RUN go mod download

COPY app .

CMD ["go", "run", "main.go"]
