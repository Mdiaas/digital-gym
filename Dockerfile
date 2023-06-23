FROM golang:alpine as builder

WORKDIR /digitalgym
COPY go.mod go.sum  ./
COPY . .
RUN GOOS=linux go build -o digitalgym ./cmd/server

FROM scratch
COPY --from=builder /digitalgym/digitalgym .
CMD ["./digitalgym"]