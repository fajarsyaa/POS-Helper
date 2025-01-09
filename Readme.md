# Project Helper API POS

## Deskripsi
Aplikasi ini dirancang sebagai alat bantu bagi para helper yang bertugas mendukung pengelolaan transaksi dalam acara offline. Aplikasi ini terintegrasi dengan sistem POS utama untuk memastikan sinkronisasi data yang akurat dan real-time.

**Use Case** 

Masalah yang sering dihadapi oleh brand clothing saat mengadakan event offline adalah terbatasnya jumlah admin yang mengelola POS. Biasanya, hanya ada satu admin dengan laptop yang bertugas di booth utama. Namun, karena tingginya jumlah pengunjung di acara yang ramai, proses transaksi sering terhambat jika hanya mengandalkan satu kasir, hal ini membuat antrean menjadi panjang.
Dengan aplikasi ini, para helper di booth dapat membantu mempercepat proses pemesanan menggunakan perangkat mobile. Helper dapat langsung membuat dan menginput pesanan melalui aplikasi tanpa harus bergantian di satu titik kasir. Hal ini memungkinkan pelayanan yang lebih cepat dan efisien selama acara berlangsung.
	
---

## Prasyarat
Untuk menjalankan proyek ini, pastikan bahwa syarat-syarat berikut telah terpenuhi:

1. **Bahasa Pemrograman Go**:
   - Instal Go.
   - Pastikan perintah `go` dapat diakses melalui PATH sistem Anda.

   Panduan instalasi: [Go Installation](https://go.dev/doc/install)

2. **Basis Data**:
   - Sebuah basis data relasional (misalnya, MySQL) harus sudah diatur.
   - Buat basis data dan kemudian import file yang tersedia. **sqlTable.sql**
   - Konfigurasikan parameter koneksi (host, port, username, password, nama basis data) di file **main.go**
   - Atur variabel *dsn* yang diperlukan untuk koneksi basis data. Contoh:
     ```
     main.go 

     // menggunakan mysql configuration
     dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

     ```    

---

## Memulai

### Langkah-langkah Menjalankan Aplikasi

1. **Clone Repository**:
   Clone repository ini ke komputer lokal Anda menggunakan perintah berikut:
   ```sh
   git clone https://github.com/fajarsyaa/POS-Helper.git
   cd POS-Helper
   ```

2. **Install Dependencies**:
   Jalankan perintah berikut untuk menginstal semua modul Go yang diperlukan:
   ```sh
   go mod tidy
   ```

3. **Jalankan Aplikasi**:
   Mulai aplikasi dengan perintah:
   ```sh
   go run main.go
   ```

---

## Catatan Tambahan
Gunakan POSTMAN atau tools serupa untuk memcoba


## API Endpoint

Berikut adalah daftar endpoint yang tersedia dalam aplikasi ini:

### User Management
- **REGISTER USER**
  - Method: POST
  - URL: `localhost:8080/api/slash/registration`
  - Request Body:
    ```json
    {
      "name": "MrCrab",
      "email": "MrCrab@fjr.fjr",
      "Role": "pengguna",
      "password": "password"
    }
    ```

- **LOGIN**
  - Method: POST
  - URL: `localhost:8080/api/slash/login`
  - Request Body:
    ```json
    {
      "email": "patrick@fjr.fjr",
      "password": "password"
    }
    ```

### Product Management
- **GET ALL PRODUCTS**
  - Method: GET
  - URL: `localhost:8080/api/slash/products`

- **FIND PRODUCT BY NAME**
  - Method: POST
  - URL: `localhost:8080/api/slash/products/name`
  - Request Body:
    ```json
    {
      "keyword": "formal"
    }
    ```

- **FIND PRODUCT BY ID**
  - Method: POST
  - URL: `localhost:8080/api/slash/products/id`
  - Request Body:
    ```json
    {
      "id": 1
    }
    ```

### Order Management
- **CREATE ORDER**
  - Method: POST
  - URL: `localhost:8080/api/slash/order`
  - Request Body:
    ```json
    {
      "products": [
        { "products_id": 1, "quantity": 1 },
        { "products_id": 2, "quantity": 3 }
      ],
      "orders_total": 448000,
      "customer_name": "John DC",
      "customer_phone": "1234567890",
      "customer_address": "123 Main St, Anytown, USA"
    }
    ```

- **GET ALL ORDERS**
  - Method: GET
  - URL: `localhost:8080/api/slash/orders`

- **GET ORDER DETAIL BY ID**
  - Method: POST
  - URL: `localhost:8080/api/slash/order/detail`
  - Request Body:
    ```json
    {
      "order_id": "ORD-8578316670019669-PWV"
    }
    ```

- **PAYMENT**
  - Method: POST
  - URL: `localhost:8080/api/slash/order/payment`
  - Request Body:
    ```json
    {
      "order_id": "ORD-061ad95f31374be5-CRW"
    }
    ```

- **UPDATE ORDER**
  - Method: PUT
  - URL: `localhost:8080/api/slash/order/edit`
  - Request Body:
    ```json
    {
      "order_id": "ORD-061ad95f31374be5-CRW",
      "product_id": 1,
      "quantity": 1
    }
    ```

- **DELETE ORDER BY ID**
  - Method: DELETE
  - URL: `localhost:8080/api/slash/order/delete`
  - Request Body:
    ```json
    {
      "order_id": "ORD-2071880178101609-YBA"
    }
    ```

---

