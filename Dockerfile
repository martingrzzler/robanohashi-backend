FROM golang:1.18-alpine as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest

COPY --from=builder /src/main .

EXPOSE 4000

CMD ["./main"]
