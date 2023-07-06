#Dokumentasi Aplikasi

##Menjalankan Aplikasi
1. Pastikan Go 1.19 dan PostgreSQL sudah terinstall di komputer Anda. Jika belum, silakan menginstalnya terlebih dahulu.
2. Buat database di PostgreSQL dengan nama **library**.
3. Atur variabel DB_PASS dan DB_USER pada file .env agar sesuai dengan konfigurasi aplikasi PosgreSQL yang ada di lokal.
4. Perbarui dependensi dengan cara menjalankan perintah go mod tidy di terminal direktory aplikasi ini.
5. Setelah dependensi selesai, jalankan aplikasi dengan cara menjalankan perintah go run . di terminal direktory aplikasi ini.
6. Aplikasi secara default akan berjalan di http://localhost:8000

##Pengujian dengan Postman

1. Registrasi
    - Deskripsi : Endpoint ini digunakan untuk mendaftarkan user baru dalam database.
    - Endpoint  : method POST, http://localhost:8080/library/auth/regist
    - Parameter : Tidak ada parameter yang diperlukan. Namun, data user harus dikirim dalam body request dengan format JSON.
    - Contoh Request    : Content-Type: application/json
                        {
                            "username" : "elvina",
                            "password" : "12345678"
                        }
                        Ketentuan :  
                            - username tidak boleh sama dengan username yang sudah ada dalam database
                            - password minimal 8 karakter
                            - username dan password harus diisi
    - Contoh Respon    : 200 OK
                        {
                            "Message": "Success Regist"
                        }
    - Authorization :Endpoint ini tidak memerlukan otorisasi.

2. Login
    - Deskripsi : Endpoint ini digunakan untuk mendapatkan token untuk otorisasi.
    - Endpoint  : method POST, http://localhost:8080/library/auth/login
    - Parameter : Tidak ada parameter yang diperlukan. Namun, data user harus dikirim dalam body request dengan format JSON.
    - Contoh Request    : Content-Type: application/json
                        {
                            "username" : "elvina",
                            "password" : "12345678"
                        }
                        Ketentuan : Masukkan username dan password yang telah terdaftar di database.
    - Contoh Respon    : 200 OK
                        {
                            "Message": "Token JWT disimpan dalam cookie 'token'"
                        }
                        Ambil token pada cookie dengan nama token untuk otorisasi
    - Authorization :Endpoint ini tidak memerlukan otorisasi.

3. Create Author 
    - Deskripsi : Endpoint ini digunakan untuk menambahkan penulis ke database.
    - Endpoint  : method POST, http://localhost:8080/library/author
    - Parameter : Tidak ada parameter yang diperlukan. Namun, data user harus dikirim dalam body request dengan format JSON.
    - Contoh Request    : Content-Type: application/json
                        {
                            "name" : "elvina fitriani",
                            "country" : "indonesia",
                            "book" : ["Golang"]
                        }
                        Ketentuan : 
                            - name tidak boleh mengandung karakter "!", "@", "#", "$" dan "%"
                            - name harus terdiri dari jumlah karakter lebih dari 3 dan kurang dari 20
                            - country harus valid berdasarkan data negara internasional terbaru
                            - book tidak boleh diinput dua kali
                            - semua variabel diatas harus diisi
    - Contoh Respon    : 200 OK .
                        {
                            "Message": "Data created successfully."
                        }
    - Authorization : Masukkan token ke Authorization type Bearer Token

4. Get All Author 
    - Deskripsi : Endpoint ini digunakan untuk mengambil data semua penulis dari database.
    - Endpoint  : method GET, http://localhost:8080/library/author
    - Parameter : Tidak ada parameter yang diperlukan
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Respon    : 200 OK .
                        {
                            "Authors": [{
                                    "ID": 7,
                                    "CreatedAt": "2023-07-06T16:08:39.302483+07:00",
                                    "UpdatedAt": "2023-07-06T16:08:39.302483+07:00",
                                    "DeletedAt": null,
                                    "name": "elvina fitriani",
                                    "country": "indonesia",
                                    "book": [
                                        "Golang"
                                    ]
                                }
                            ],
                            "Response": "Data retrieved successfully."
                        }

5. Get all Books for a specific Author
    - Deskripsi : Endpoint ini digunakan untuk mengambil data semua penulis yang menulis sebuah buku yang diminta.
    - Endpoint  : method GET, http://localhost:8080/library/author/elvina fitriani
    - Parameter : Masukkan nama author yang sudah terdaftar dalam database sebagai parameter
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Respon    : 200 OK .
                        ["Golang"]

6. Delete Author
    - Deskripsi : Endpoint ini digunakan untuk menghapus data penulis yang ada dimasukkan ke paramater dari database.
    - Endpoint  : method DELETE, http://localhost:8080/library/author/elvina fitriani
    - Parameter : Masukkan isbn buku yang sudah terdaftar dalam database sebagai parameter
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Respon :   {
                            "Response": "Data deleted successfully."
                        }

7. Update Author
    - Deskripsi : Endpoint ini digunakan untuk update data penulis yang ada dimasukkan ke paramater dari database.
    - Endpoint  : method PUT, http://localhost:8080/library/author/elvina fitriani
    - Parameter : Masukkan nama penulis yang sudah terdaftar dalam database sebagai parameter
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Request    : Content-Type: application/json
                        {
                            "name" : "elvina fitriani",
                            "country" : "indonesia",
                            "book" : ["Backend Golang"]
                        }
                        Ketentuan : 
                            - name tidak boleh mengandung karakter "!", "@", "#", "$" dan "%"
                            - name harus terdiri dari jumlah karakter lebih dari 3 dan kurang dari 20
                            - country harus valid berdasarkan data negara internasional terbaru
                            - book tidak boleh diinput dua kali
                            - semua variabel diatas harus diisi
    - Contoh Respon :   {
                            "Message": "Data updated successfully."
                        }

   
8. Create Book 
    - Deskripsi : Endpoint ini digunakan untuk menambahkan buku ke database.
    - Endpoint  : method POST, http://localhost:8080/library/book
    - Parameter : Tidak ada parameter yang diperlukan. Namun, data user harus dikirim dalam body request dengan format JSON.
    - Contoh Request    : Content-Type: application/json
                        {
                            "title" : "Golang",
                            "publishedYear" : 2023,
                            "isbn" : "1234567890",
                            "author" : [
                                "elvina fitriani"
                            ]
                        }
                        Ketentuan : 
                            - publishedYear tidak boleh lebih kecil dari 1900 dan tidak boleh lebih besar dari tahun sekarang
                            - isbn hanya boleh terdiri dari angka dan tanda -
                            - isbn terdiri dari 10 atau 13 angka
                            - author harus sudah terdaftar di database
                            - author tidak boleh diinput dua kali
                            - semua variabel diatas harus diisi
    - Contoh Respon    : 200 OK .
                        {
                            "Message": "Data created successfully."
                        }
    - Authorization : Masukkan token ke Authorization type Bearer Token

9. Get All Book 
    - Deskripsi : Endpoint ini digunakan untuk mengambil data semua buku dari database.
    - Endpoint  : method GET, http://localhost:8080/library/book
    - Parameter : Tidak ada parameter yang diperlukan
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Respon    : 200 OK .
                        {
                            "Books": [{
                                    "ID": 7,
                                    "CreatedAt": "2023-07-06T16:08:39.302483+07:00",
                                    "UpdatedAt": "2023-07-06T16:08:39.302483+07:00",
                                    "DeletedAt": null,
                                    "title" : "Golang",
                                    "publishedYear" : 2023,
                                    "isbn" : "1234567890",
                                    "author" : [
                                        "elvina fitriani"
                                    ]
                                }
                            ],
                            "Response": "Data retrieved successfully."
                        }

10. Get all Authors for a specific Book
    - Deskripsi : Endpoint ini digunakan untuk mengambil data semua buku yang ditulis oleh penulis yang diminta.
    - Endpoint  : method GET, http://localhost:8080/library/book/Golang
    - Parameter : Masukkan isbn buku yang sudah terdaftar dalam database sebagai parameter
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Respon    : ["elvina fitriani"]

11. Delete Book
    - Deskripsi : Endpoint ini digunakan untuk menghapus data buku yang ada dimasukkan ke paramater dari database.
    - Endpoint  : method DELETE, http://localhost:8080/library/book
    - Parameter : Masukkan isbn buku yang sudah terdaftar dalam database sebagai parameter
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Respon :   {
                            "Response": "Data deleted successfully."
                        }


12. Update Book
    - Deskripsi : Endpoint ini digunakan untuk update data buku yang ada dimasukkan ke paramater dari database.
    - Endpoint  : method PUT, http://localhost:8080/library/book/Golang
    - Parameter : Masukkan isbn buku yang sudah terdaftar dalam database sebagai parameter
    - Authorization : Masukkan token ke Authorization type Bearer Token
    - Contoh Request    : Content-Type: application/json
                        {
                            "title" : "Golang",
                            "publishedYear" : 2023,
                            "isbn" : "1234567890",
                            "author" : [
                                "elvina"
                            ]
                        }
                        Ketentuan : 
                            - publishedYear tidak boleh lebih kecil dari 1900 dan tidak boleh lebih besar dari tahun sekarang
                            - isbn hanya boleh terdiri dari angka dan tanda -
                            - isbn terdiri dari 10 atau 13 angka
                            - author harus sudah terdaftar di database
                            - author tidak boleh diinput dua kali
                            - semua variabel diatas harus diisi
    - Contoh Respon :   {
                            "Message": "Data updated successfully."
                        }

        

