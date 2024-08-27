# Order Service

Order Service bertanggung jawab untuk mengelola pesanan yang dibuat oleh pengguna. Service ini memungkinkan pengguna untuk membuat pesanan dan mengambil informasi pesanan yang telah dibuat. Semua operasi dilindungi oleh autentikasi JWT.

## Cara Menjalankan

1. **Pastikan PostgreSQL berjalan** dan Anda telah membuat database `order` dengan tabel `orders`.

2. **Clone repository ini** dan navigasi ke direktori `order`:
   ```bash
   cd go-microservices/order
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Jalankan migrasi database** (opsional jika database dan tabel belum dibuat):
   ```bash
   psql -U <username> -d order -a -f migrations/001_create_orders_table.sql
   ```
   Gantilah `<username>` dengan username PostgreSQL Anda.

5. **Jalankan service**:
   ```bash
   go run cmd/main.go
   ```

6. **Service berjalan pada port 50052**.

## Cara Mengakses Endpoint

### 1. Buat Pesanan

Endpoint ini digunakan untuk membuat pesanan baru. Setiap request ke endpoint ini memerlukan token JWT yang valid.

- **URL:** `http://localhost:8080/order`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
  - `Authorization: Bearer <your_jwt_token>`
- **Body (JSON):**
  ```json
  {
      "product_id": "uuid-product-id",
      "quantity": 2
  }
  ```

- **Response (Contoh):**
  ```json
  {
      "id": "uuid-order-id",
      "product_id": "uuid-product-id",
      "quantity": 2,
      "user_id": "uuid-user-id"
  }
  ```


## Troubleshooting

- **401 Unauthorized:** Pastikan Anda telah menyertakan header `Authorization` dengan token JWT yang valid.
- **400 Bad Request:** Periksa kembali format JSON yang dikirim dalam body request Anda.
- **500 Internal Server Error:** Pastikan service `Order` berjalan dengan benar dan database terhubung tanpa masalah.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [gRPC](https://grpc.io/) - RPC framework
- [GORM](https://gorm.io/) - ORM untuk Go
- [JWT](https://github.com/golang-jwt/jwt) - Library untuk pembuatan dan verifikasi JWT