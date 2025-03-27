# Stage 1: Build Stage
FROM golang:1.24.1 AS builder

# Set working directory dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Unduh dependensi
RUN go mod tidy

# Build aplikasi
RUN go build -o main ./cmd/main.go

# Stage 2: Runtime Stage
FROM debian:bookworm-slim

WORKDIR /app

# Copy binary dari stage pertama
COPY --from=builder /app/main .

# Copy file konfigurasi jika ada
COPY config ./config

# Install Redis CLI jika diperlukan untuk debugging
RUN apt-get update && apt-get install -y redis-tools && rm -rf /var/lib/apt/lists/*

# Expose port 9000 (sesuai dengan yang kamu gunakan)
EXPOSE 9000

# Jalankan aplikasi
CMD ["./main"]
