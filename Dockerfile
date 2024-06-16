FROM golang:1.22.2-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./application
CMD ["/app/application"]

