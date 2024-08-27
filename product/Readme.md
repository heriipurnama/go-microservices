# Product Service

Product Service bertanggung jawab untuk mengelola data produk, termasuk pembuatan produk baru dan pengambilan informasi produk. Semua operasi dilindungi oleh autentikasi JWT.

## Cara Menjalankan

1. **Pastikan PostgreSQL berjalan** dan Anda telah membuat database `product` dengan tabel `products`.

2. **Clone repository ini** dan navigasi ke direktori `product`:
   ```bash
   cd go-microservices/product
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Jalankan migrasi database** (opsional jika database dan tabel belum dibuat):
   ```bash
   psql -U <username> -d product -a -f migrations/001_create_products_table.sql
   ```
   Gantilah `<username>` dengan username PostgreSQL Anda.

5. **Jalankan service**:
   ```bash
   go run cmd/main.go
   ```

6. **Service berjalan pada port 50053**.

## Cara Mengakses Endpoint

### 1. Buat Produk

Endpoint ini digunakan untuk membuat produk baru. Setiap request ke endpoint ini memerlukan token JWT yang valid.

- **URL:** `http://localhost:8080/product`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: Bearer <your_jwt_token>`
- **Body (JSON):**
  ```json
  {
      "name": "Product Name",
      "description": "This is a product description",
      "price": 99.99
  }
  ```

- **Response (Contoh):**
  ```json
  {
      "id": "uuid-value",
      "name": "Product Name",
      "description": "This is a product description",
      "price": 99.99
  }
  ```

## Troubleshooting

- **401 Unauthorized:** Pastikan Anda telah menyertakan header `Authorization` dengan token JWT yang valid.
- **400 Bad Request:** Periksa kembali format JSON yang dikirim dalam body request Anda.
- **500 Internal Server Error:** Pastikan service `Product` berjalan dengan benar dan database terhubung tanpa masalah.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [gRPC](https://grpc.io/) - RPC framework
- [GORM](https://gorm.io/) - ORM untuk Go
- [JWT](https://github.com/golang-jwt/jwt) - Library untuk pembuatan dan verifikasi JWT
