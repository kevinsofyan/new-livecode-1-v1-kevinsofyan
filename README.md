[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/K1deSPM5)
[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)]([https://classroom.github.com/a/R1iT3SE6)
# Live-Code-1-Phase-2
## RULES
1. **Untuk kampus remote**: **WAJIB** melakukan **share screen**(**DESKTOP/ENTIRE SCREEN**) dan **unmute microphone** ketika Live Code
berjalan (tidak melakukan share screen/salah screen atau tidak unmute microphone akan di ingatkan).
2. Kerjakan secara individu. Segala bentuk kecurangan (mencontek ataupun diskusi) akan menyebabkan skor live code ini 0.
3. Waktu pengerjaan: **90 menit**
4. **Pada text editor hanya ada file yang terdapat pada repository ini**.
5. Membuka referensi eksternal seperti Google, StackOverflow, dan MDN diperbolehkan.
6. Dilarang membuka repository di organisasi tugas, baik pada organisasi batch sendiri ataupun batch lain, baik branch sendiri maupun branch orang
lain (**setelah melakukan clone, close tab GitHub pada web browser kalian**).

## Objectives
- Menguji pemahaman mengenai konsep REST API
- Menguji pemahaman dan kemampuan membuat REST API menggunakan Golang
- Menguji pemahaman dan kemampuan integrasi REST API dengan implementasi SQL Database

## Requirements
Anda diminta untuk membuat sebuah RESTful API untuk sebuah modul Order dengan menggunakan Golang dan MySQL. API harus memiliki fitur operasi CRUD (Create, Read/Retrieve, Update dan Delete) untuk data Order. Setiap data Order terdiri dari unique ID, nama pembeli, nama toko penjual, nama item, dan jumlah quantity item yang dibeli, serta timestamp kapan order tersebut dibuat.

- Untuk memenuhi kriteria diatas, buatlah REST API yang memiliki endpoint sebagai berikut
  |Endpoint|Description|
  |---|---|
  |GET /orders| Menampilkan seluruh data orders|
  |GET /orders/:id| Menampilkan data shipment berdasarkan ID|
  |POST /orders| Membuat/Menyimpan data shipment baru|
  |PUT /orders/:id| Memperbaharui data shipment berdasarkan ID|
  |DELETE /orders/:id| Menghapus data shipment berdasarkan ID|
- Pastikan anda mengikuti best practice REST API untuk http method dan response http status.
- Pastikan request dan response setiap endpoint mengikuti kontrak dokumentasi API yang telah disediakan
- Pastikan untuk handle negative casses dan edge cases yang memungkinkan terjadi, dan response sesuai dengan dokumentasi API yang terlah disediakan

## Assignment Criteria Notes
Live code ini memiliki bobot nilai sebagai berikut:

|Criteria|Meet Expectations|
|---|---|
|Problem Solving|5 API Endpoints are implemented and working correctly|
|Database Design |MySQL database meets the required specifications|
||Database queries are efficient and appropriately indexed|
|Readability|Code is well-documented and easy to read|
||Code includes appropriate comments and documentation|

### Assignment Notes:
- Jangan terburu-buru dalam menyelesaikan masalah atau mencoba untuk menyelesaikannya sekaligus.
- Jangan menyalin kode dari sumber eksternal tanpa memahami bagaimana kode tersebut bekerja.
- Jangan menentukan nilai secara hardcode atau mengandalkan asumsi yang mungkin tidak berlaku dalam semua kasus.
- Jangan lupa untuk menangani negative case, seperti input yang tidak valid
- Jangan ragu untuk melakukan refaktor kode Anda, buatlah struktur project anda lebih mudah dibaca dan dikembangkan kedepannya, pisahkanlah setiap bagian kode program pada folder sesuai dengan tugasnya masing-masing.

### Additional Notes
Total Points : 100

Deadline : Diinformasikan oleh instruktur saat briefing LC. Keterlambatan pengumpulan livecode mengakibatkan skor LC 1 menjadi 0.

Informasi yang tidak dicantumkan pada file ini harap dipastikan/ditanyakan kembali kepada instruktur. Kesalahan asumsi dari peserta mungkin akan menyebabkan kesalahan pemahaman requirement dan mengakibatkan pengurangan nilai.
