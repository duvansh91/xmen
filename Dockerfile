FROM golang:1.17 as go-builder

WORKDIR /xmen

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY cmd cmd
COPY pkg pkg
COPY internal internal

WORKDIR /xmen/cmd

RUN go build -ldflags="-s -w" -o /main main.go

EXPOSE 8080

CMD ["/main"]
