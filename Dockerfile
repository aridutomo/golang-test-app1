# Gunakan gambar resmi Golang untuk membuat artefak build
FROM golang:1.18-alpine AS builder

# Setel direktori kerja di dalam container
WORKDIR /app

# Salin file go mod dan sum
COPY go.mod go.sum ./

# Unduh semua dependensi. Dependensi akan di-cache jika file go.mod dan go.sum tidak berubah
RUN go mod download

# Salin source code dari direktori saat ini ke direktori kerja di dalam container
COPY . .

# Bangun aplikasi Go
RUN go build -o main .

# Mulai tahap baru dari awal
FROM alpine:latest  

# Setel direktori kerja di dalam container
WORKDIR /app

# Salin file binary yang sudah dibangun dari tahap sebelumnya
COPY --from=builder /app/main .

# Ekspos port 8080 ke dunia luar
EXPOSE 8080

# Perintah untuk menjalankan executable
CMD ["./main"]