FROM golang:alpine

RUN mkdir -p /app
WORKDIR /app

COPY . .
RUN GOOS=linux   GOARCH=amd64  go build

ENTRYPOINT ["./bbmri-fhir-gen"]