# Auth Service

Auth Service bertanggung jawab untuk mengelola autentikasi pengguna, termasuk pendaftaran, login, dan pembuatan token JWT. Service ini digunakan untuk memverifikasi identitas pengguna sebelum mereka dapat mengakses layanan lainnya.

## Cara Menjalankan

1. **Pastikan PostgreSQL berjalan** dan Anda telah membuat database `auth` dengan tabel `users`.

2. **Clone repository ini** dan navigasi ke direktori `auth`:
   ```bash
   cd go-microservices/auth
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Jalankan migrasi database** (opsional jika database dan tabel belum dibuat):
   ```bash
   psql -U <username> -d auth -a -f migrations/001_create_users_table.sql
   ```
   Gantilah `<username>` dengan username PostgreSQL Anda.

5. **Jalankan service**:
   ```bash
   go run cmd/main.go
   ```

6. **Service berjalan pada port 50051**.

## Cara Mengakses Endpoint

### 1. Register User

Endpoint ini digunakan untuk mendaftarkan pengguna baru.

- **URL:** `http://localhost:8080/register`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
- **Body (JSON):**
  ```json
  {
      "username": "testuser",
      "password": "password1234"
  }
  ```

- **Response (Contoh):**
  ```json
  {
      "message": "User registered successfully"
  }
  ```

### 2. Login User

Endpoint ini digunakan untuk login dan mendapatkan token JWT.

- **URL:** `http://localhost:8080/login`
- **Method:** `POST`
- **Headers:** 
  - `Content-Type: application/json`
- **Body (JSON):**
  ```json
  {
      "username": "testuser",
      "password": "password1234"
  }
  ```

- **Response (Contoh):**
  ```json
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

### 3. Menggunakan Token JWT

Setelah mendapatkan token JWT, Anda dapat menggunakannya untuk mengakses layanan lain yang memerlukan autentikasi. 

- **Gunakan header berikut pada request lain:**
  ```plaintext
  Authorization: Bearer <your_jwt_token>
  ```
  Gantilah `<your_jwt_token>` dengan token yang Anda dapatkan dari respons login.

## Troubleshooting

- **401 Unauthorized:** Pastikan username dan password yang Anda gunakan benar dan token JWT yang Anda berikan valid.
- **400 Bad Request:** Pastikan format JSON yang dikirim dalam body request sudah benar.
- **500 Internal Server Error:** Periksa apakah service `Auth` sudah berjalan dan terhubung dengan database dengan benar.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [gRPC](https://grpc.io/) - RPC framework
- [GORM](https://gorm.io/) - ORM untuk Go
- [JWT](https://github.com/golang-jwt/jwt) - Library untuk pembuatan dan verifikasi JWT
