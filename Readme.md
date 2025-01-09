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
