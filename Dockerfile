FROM golang:alpine as builder

WORKDIR /digitalgym
COPY go.mod go.sum /cmd/server/.env ./
COPY /cmd/server/.env /digitalgym
COPY . .
RUN GOOS=linux CGO_ENABLED=1 go build -o digitalgym ./cmd/server

FROM scratch
COPY --from=builder /digitalgym/digitalgym .
COPY --from=builder /digitalgym/.env .
CMD ["./digitalgym"]