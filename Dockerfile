FROM golang:latest

WORKDIR /root/envmon
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["envmon"]