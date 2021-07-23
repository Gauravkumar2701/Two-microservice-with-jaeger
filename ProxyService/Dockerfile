FROM golang:latest as builder
COPY go.mod go.sum /go/src/ProxyService/
WORKDIR /go/src/ProxyService
RUN go mod download
COPY . /go/src/ProxyService
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/ProxyService .
EXPOSE 8090 8090
CMD ["./main"]
