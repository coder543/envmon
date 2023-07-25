FROM golang:latest

WORKDIR /root/envmon

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install -v ./...

CMD ["envmon"]