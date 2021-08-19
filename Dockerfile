FROM golang:1.15-alpine

WORKDIR /app

copy go.mod ./
copy go.sum ./

#Fetch and build all of the components of JACK-AL
RUN go mod download

copy *.go ./

RUN go build -o /jackal

CMD ["/jackal"]