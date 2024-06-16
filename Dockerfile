FROM golang:1.22.2-alpine as builder 

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./application

FROM scratch 
WORKDIR /app
COPY --from=builder /build/application ./application
COPY --from=builder /build/assets/. ./assets/.
COPY --from=builder /build/lib/. ./lib/.
COPY --from=builder /build/templates/. ./templates/.
COPY --from=builder /build/PALfiles/. ./PALfiles/.
CMD ["/app/application"]
