FROM golang:latest as builder
COPY go.mod go.sum /go/src/StudentEndpoints/
WORKDIR /go/src/StudentEndpoints
RUN go mod download
COPY . /go/src/StudentEndpoints
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/StudentEndpoints .
EXPOSE 8080 8080
CMD ["./main"]
