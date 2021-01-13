FROM golang:latest AS builder

WORKDIR /app
COPY src .
RUN go build -o main

FROM golang:latest

COPY --from=builder /app/main /entrypoint
EXPOSE 8000

ENTRYPOINT [ "/entrypoint" ]
