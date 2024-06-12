# Todo List App

Aplikasi sederhana untuk mengelola daftar todo menggunakan bahasa pemrograman Go, framework Fiber, dan database MySQL.

## Deskripsi

Aplikasi ini memungkinkan pengguna untuk membuat, mengubah, menghapus, menandai sebagai selesai, dan mendapatkan daftar todo. Setiap todo dapat memiliki beberapa tugas yang terkait.

## Fitur

- **CRUD Todo**: Pengguna dapat membuat, mengubah, dan menghapus todo.
- **Tandai Selesai**: Pengguna dapat menandai todo sebagai selesai.
- **Validasi**: Terdapat validasi untuk memastikan tidak ada duplikasi judul todo dan batasan pada pembaruan dan penghapusan todo yang sudah selesai atau dihapus.
- **Pengambilan Data**: Pengguna dapat mengambil daftar semua todo atau todo berdasarkan ID.

## Instalasi

1. Pastikan Go telah terinstal di komputer Anda. Jika belum, Anda dapat mengunduhnya [di sini](https://golang.org/dl/).
2. Clone repositori ini ke komputer Anda.
3. Masuk ke direktori proyek dan jalankan perintah `go mod tidy` untuk mengunduh dependensi yang diperlukan.
4. Pastikan MySQL telah terpasang di komputer Anda dan konfigurasikan koneksi ke database di `config/config.go`.
5. Buatlah database dengan nama `tododb` di MySQL.
6. Jalankan perintah `go run main.go` untuk memulai server.

## Penggunaan

Setelah menjalankan server, Anda dapat mengakses API menggunakan berbagai metode HTTP seperti POST, GET, PUT, PATCH, dan DELETE. Berikut adalah beberapa contoh penggunaan:

- **Membuat Todo**:

  ```sh
  curl -X POST http://localhost:3000/api/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Membuat Proyek",
    "tasks": [
      {"title": "Mempersiapkan Dokumentasi"},
      {"title": "Mengembangkan Fitur Utama"}
    ]
  }'
  ```

- **Mendapatkan Daftar Todo**:

  ```sh
  curl -X GET http://localhost:3000/api/todos
  ```

- **Mendapatkan Todo Berdasarkan ID**:
  ```sh
  curl -X GET http://localhost:3000/api/todos/{ID}
  ```
- **Mengubah Todo Berdasarkan ID**:

  ```sh
  curl -X PUT http://localhost:3000/api/todos/{ID} \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Membuat Proyek",
    "tasks": [
      {"title": "Mempersiapkan Dokumentasi"},
      {"title": "Mengembangkan Fitur Utama"}
    ]
  }'
  ```

- **Menghapus Todo Berdasarkan ID**:
  ```sh
  curl -X DELETE http://localhost:3000/api/todos/{ID}
  ```
- **Menandai Todo Selesai Berdasarkan ID**:
  ```sh
  curl -X PATCH http://localhost:3000/api/todos/{ID}/complete
  ```

**Created By :** Ryan Dinul Fatah.
