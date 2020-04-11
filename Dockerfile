FROM golang:1.14.2-alpine3.11

WORKDIR /src

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=8080

RUN (cd  ./cmd && go build -o ../app)

CMD ["./app"]