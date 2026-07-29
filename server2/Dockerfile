FROM golang:tip-trixie AS base

WORKDIR /ServerTwoDir

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "./server2/server2.go"]
