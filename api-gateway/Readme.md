# API Gateway

API Gateway bertindak sebagai titik masuk tunggal untuk berbagai layanan (Auth, Product, Order) yang dijalankan di backend. API Gateway menangani autentikasi JWT dan meneruskan permintaan ke layanan yang sesuai.

## Cara Menjalankan

1. **Pastikan semua layanan backend (Auth, Product, Order) sudah berjalan** dan database yang sesuai sudah diinisialisasi.

2. **Clone repository ini** dan navigasi ke direktori `api-gateway`:
   ```bash
   cd go-microservices/api-gateway
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Jalankan API Gateway**:
   ```bash
   go run cmd/main.go
   ```

5. **API Gateway berjalan pada port 8080**.

## Cara Mengakses Endpoint

### 1. **Register User**
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

### 2. **Login User**
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

### 3. **Buat Produk**
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
         "price": 99.99"
     }
     ```

### 4. **Buat Pesanan**
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
- **500 Internal Server Error:** Pastikan service backend (Auth, Product, Order) sudah berjalan dan API Gateway terhubung dengan benar ke service-service tersebut.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [gRPC](https://grpc.io/) - RPC framework
- [GORM](https://gorm.io/) - ORM untuk Go
- [JWT](https://github.com/golang-jwt/jwt) - Library untuk pembuatan dan verifikasi JWT