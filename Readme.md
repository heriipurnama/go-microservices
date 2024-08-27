# Go Microservices

Go Microservices adalah proyek contoh yang terdiri dari beberapa service mikro yang dibangun menggunakan bahasa Go. Setiap service dirancang untuk berkomunikasi satu sama lain melalui gRPC, dengan API Gateway yang bertindak sebagai titik masuk tunggal untuk semua request dari klien.

## Layanan yang Tersedia

Proyek ini terdiri dari empat layanan utama:

1. **Auth Service** - Bertanggung jawab untuk autentikasi pengguna, termasuk pendaftaran, login, dan pengelolaan token JWT.
2. **Product Service** - Mengelola data produk, termasuk pembuatan dan pengambilan informasi produk.
3. **Order Service** - Mengelola pesanan pengguna, termasuk pembuatan dan pengambilan informasi pesanan.
4. **API Gateway** - Bertindak sebagai titik masuk utama untuk semua request dari klien, meneruskan permintaan ke layanan backend yang sesuai, dan menangani autentikasi JWT.

## Teknologi yang Digunakan

- **Go** - Bahasa pemrograman utama untuk membangun microservices ini.
- **gRPC** - Digunakan untuk komunikasi antar layanan.
- **PostgreSQL** - Digunakan sebagai basis data untuk menyimpan data pengguna, produk, dan pesanan.
- **GORM** - ORM untuk Go yang digunakan untuk interaksi dengan database.
- **JWT** - Digunakan untuk mengelola autentikasi dan otorisasi pengguna.

## Struktur Direktori

```plaintext
go-microservices/
│
├── api-gateway/        # Kode untuk API Gateway
│
├── auth/               # Kode untuk Auth Service
│
├── order/              # Kode untuk Order Service
│
├── product/            # Kode untuk Product Service
│
└── docker-compose.yml  # Konfigurasi Docker Compose untuk menjalankan seluruh stack
```

## Cara Menjalankan

1. **Kloning Repository**:
   ```bash
   git clone https://github.com/heriipurnama/go-microservices.git
   cd go-microservices
   ```

2. **Jalankan Layanan dengan Docker Compose**:
   Pastikan Anda memiliki Docker dan Docker Compose terinstal, kemudian jalankan:
   ```bash
   docker-compose up --build
   ```

   Ini akan membangun dan menjalankan semua layanan yang didefinisikan dalam `docker-compose.yml`.

3. **Akses API Gateway**:
   API Gateway akan berjalan di `http://localhost:8080`. Anda dapat menggunakan Postman atau alat lain untuk mengakses endpoint yang tersedia.

## Author

- **Nama:** Heri Purnama
- **GitHub:** [heriipurnama](https://github.com/heriipurnama)
- **Email:** [herii.purnama@outlook.com](mailto:herii.purnama@outlook.com)
- **Web:** [heriipurnama](https://heriipurnama.rumahinformatika.com)

Proyek ini dikembangkan sebagai contoh implementasi arsitektur microservices menggunakan Go. Semua kontribusi dan saran untuk peningkatan sangat diterima.