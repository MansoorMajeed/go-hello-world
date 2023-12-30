FROM golang:latest AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Second stage: Create a smaller image
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

