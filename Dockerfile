# ==========================
# Stage 1: Build binary
# ==========================
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Copy module files dulu (cache dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./main.go

# ==========================
# Stage 2: Runtime
# ==========================
FROM alpine:3.20

WORKDIR /app

# Copy binary dari stage builder
COPY --from=builder /app/server /app/server

# Optional: copy migrations folder
COPY --from=builder /app/migrations /app/migrations

# Install certificates & set executable permission
RUN apk add --no-cache ca-certificates && chmod +x /app/server

# Set environment & expose port
ENV PORT=8000
EXPOSE 8000

# Run the binary
CMD ["/app/server"]
