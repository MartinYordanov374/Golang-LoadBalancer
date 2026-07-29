FROM golang:tip-trixie AS base

WORKDIR /loadbalancerdir

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD ["go", "run", "./loadbalancer/loadbalancer.go"]
