FROM golang:1.24.1

WORKDIR /app

# install air
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum /app/
RUN cd /app && go mod download

COPY app .

CMD ["air"]
