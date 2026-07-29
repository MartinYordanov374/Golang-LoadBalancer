FROM golang:tip-trixie AS base

WORKDIR /ServerThreeDir

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "./server3/server3.go"]
