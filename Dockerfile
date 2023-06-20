# Start from golang base image
FROM golang:alpine as builder

RUN apk update
WORKDIR /src/digitalgym
COPY go.mod go.sum ./
COPY . .
RUN go build -o digitalgym ./cmd/server

FROM alpine:3.16 as binary
COPY --from=builder /src/digitalgym/digitalgym .
EXPOSE 3000
CMD ["./digitalgym"]