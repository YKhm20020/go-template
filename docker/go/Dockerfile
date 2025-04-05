FROM golang:1.24.1

WORKDIR /app

# install air
RUN go install github.com/air-verse/air@latest

COPY app/ .

RUN go mod download

CMD ["air"]
