# Submission 2 - Student Routes

Ini merupakan submission kedua yang akan menguji pemahaman pembelajaran kamu selama live session berlangsung. Kamu dapat melihat live session yang telah diberikan sebelumnya untuk membantu kamu dalam menyelesaikan submission ini

## Storyboard

Disini kamu perlu untuk menyelesaikan setiap fungsi yang kosong dengan solusi yang dapat kamu implementasikan. Kamu juga bisa kembali melihat Live Session sebelumnya dimana implementasi yang ada pada submission ini sangat sering dibahas disana. Jadi kamu perlu menyelesaikan Submission ini berdasarkan Live Session atau gaya penyelesaian yang berbeda selama memenuhi aturan yang diberikan pada `main_test.go`

Untuk proses inisiasi submission ini, dimohon untuk melihat `go.mod` sehingga dapat membantu kamu mengenai package apa saja yang perlu kamu persiapkan agar proses pemberian solusi dapat berjalan dengan baik

## Aturan

- Kamu hanya menuliskan solusi yang ada hanya pada file `main.go`
- Jangan mengganti/menghapus `main_test.go` atau kamu tidak dapat melihat hasil testing dari solusi yang kamu berikan pada `main.go`
- Jikalau kamu telah menyelesaikan submission ini, kamu perlu update penyelesain tersebut pada Submission Completion Sheet yang telah diberikan pada Halaman Pertama.

## Code Template

Tulis kode kamu di `main.go`. Untuk pengembangan, kamu dapat menggunakan Postman atau Thunder Client untuk menguji setiap fungsi/routes yang telah kamu selesaikan.

## Testing

Tes unit untuk proyek ini ada di `main_test.go`. Proses pengujian dari `main_test.go` demi kenyamanan kamu. Tes akan berjalan secara otomatis setiap kali kamu menjalankan program `main_test.go`.

### Cara Testing

- Buka 2 Terminal berbeda
- Pada terminal pertama, kamu perlu menjalankan testing pertama dengan command `go test -v`. Ini tidak akan memunculkan hasil testing dan hanya memperlihatkan bahwa web server telah berjalan
- Pada terminal kedua, kamu perlu menjalankan testing kedua dengan command `go test -v -coverprofile=coverage.out` untuk memperlihatkan apakah kamu telah melewati setiap testing yang ada berdasarkan fungsi yang telah kamu selesaikan.
- `PASS` pada setiap fungsi (5 fungsi)mengartikan bahwa kamu berhasil menyelesaikan submission
- Jikalau kamu menghiraukan hasil coverage yang ada, kamu bisa improve lebih baik lagi. Namun tidak menjadi nilai tambah untuk submission ini.
