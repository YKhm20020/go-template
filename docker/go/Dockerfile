FROM golang:1.24.1

WORKDIR /src/app

# install air
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum /src/
RUN cd /src && go mod download

COPY src/app .

CMD ["air"]
