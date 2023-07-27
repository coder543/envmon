FROM golang:latest

WORKDIR /root/envmon

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install -v ./...

# yes, committing this is not best practice.
ENV INFLUXDB_TOKEN=fxndV7rpLBcD3A5IDGz5hiyYElvQW_Cjf1XcNye9Mr_Te_7gNBUQTCvMuqskUT2IWPWjmlybWoln86ql5PykeQ==

CMD ["envmon"]