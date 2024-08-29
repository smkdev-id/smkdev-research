# Submission 1 - Arithmetic Formatter

Submission ini tersinpirasi dengan [Scientific Computing with Python Projects - Arithmetic Formatter](https://www.freecodecamp.org/learn/scientific-computing-with-python/scientific-computing-with-python-projects/arithmetic-formatter). Kamu boleh melihat contoh implementasi dan penyelesaiannya dan mencoba menyelesaikan Submission ini.

## Storyboard

Siswa di sekolah dasar sering menyusun soal-soal aritmatika secara vertikal agar lebih mudah diselesaikan. Misalnya, "235 + 52" menjadi:

```
  235
+  52
-------
```

Buat fungsi yang menerima daftar string yang merupakan soal aritmatika dan mengembalikan soal yang disusun secara vertikal dan berdampingan. Fungsi tersebut secara opsional harus mengambil argumen kedua. Jika argumen kedua disetel ke `True`, jawabannya akan ditampilkan.

### Contoh

Pemanggilan Fungsi dengan Input Data:

```
arithmetic_arranger(["32 + 698", "3801 - 2", "45 + 43", "123 + 49"])
```

Output:

```
   32      3801      45      123
+ 698    -    2    + 43    +  49
-----    ------    ----    -----
```

Pemanggilan Fungsi dengan Input Data dan Kondisi True:

```
arithmetic_arranger(["32 + 8", "1 - 3801", "9999 + 9999", "523 - 49"], True)
```

Output:

```
  32         1      9999      523
+  8    - 3801    + 9999    -  49
----    ------    ------    -----
  40     -3800     19998      474
```

## Aturan

Fungsi ini akan mengembalikan konversi yang benar jika masalah yang diberikan diformat dengan benar, jika tidak, fungsi tersebut akan mengembalikan string yang menjelaskan kesalahan yang berarti bagi pengguna.

- Situasi yang akan menghasilkan kesalahan:
  - Jika ada terlalu banyak masalah yang diberikan pada fungsi tersebut. Batasnya lima, lebih dari itu akan kembali: `Error: Too many problems.`
  - Operator yang sesuai yang akan diterima fungsi tersebut adalah penjumlahan dan pengurangan. Perkalian dan pembagian akan menghasilkan kesalahan. Operator lain yang tidak disebutkan dalam poin-poin ini tidak perlu diuji. Kesalahan yang dikembalikan adalah: `Error: Operator must be '+' or '-'.`
  - Setiap angka (operan) hanya boleh berisi angka. Jika tidak, fungsi akan mengembalikan: `Error: Numbers must only contain digits.`
  - Setiap operan (alias angka di setiap sisi operator) memiliki lebar maksimal empat digit. Jika tidak, string kesalahan yang dikembalikan adalah: `Error: Numbers cannot be more than four digits.`
- Jika pengguna memberikan format masalah yang benar, konversi yang kamu kembalikan akan mengikuti aturan berikut:
  - Harus ada satu spasi antara operator dan yang terpanjang dari dua operan, operator akan berada di baris yang sama dengan operan kedua, kedua operan akan berada dalam urutan yang sama seperti yang disediakan (yang pertama akan menjadi yang teratas dan yang terakhir kedua akan menjadi bagian bawah).
  - Angka harus rata kanan.
  - Harus ada empat spasi di antara setiap masalah.
  - Harus ada tanda hubung di bagian bawah setiap soal. Tanda hubung harus berada di sepanjang setiap soal satu per satu. (Contoh di atas menunjukkan seperti apa tampilannya.)

## Code Template

Tulis kode kamu di `arithmetics_calculator.go`. Untuk pengembangan, kamu dapat menggunakan `main.go` untuk menguji fungsi `ArithmeticArranger()` kamu.

## Testing

Tes unit untuk proyek ini ada di `arithmetics_calculator_test.go`. Proses pengujian dari `arithmetics_calculator_test.go` demi kenyamanan kamu. Tes akan berjalan secara otomatis setiap kali kamu menjalankan program `main.go`. Alternatifnya, kamu dapat menjalankan tes dengan memasukkan `go test arithmetics_calculator.go arithmetics_calculator_test.go -v -coverprofile=coverage.out` di terminal.
