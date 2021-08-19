FROM golang:1.15-alpine

WORKDIR /app

ENV GOPATH /go

COPY go.mod ./
COPY go.sum ./

#Fetch the needed components to get started on building JACK-AL
RUN apk update
RUN apk add git
RUN go mod download

copy * ./

RUN go build -o /jackal

CMD ["/jackal"]

