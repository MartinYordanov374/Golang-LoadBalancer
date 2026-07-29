FROM golang:tip-trixie AS base

WORKDIR /ServerOneDir

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "./server1/server1.go"]
