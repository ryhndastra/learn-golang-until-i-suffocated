# DAY 02: Control Flow & Memory Fundamentals (Pointers)

> **Target:** Menguasai idiom percabangan Go (`if` short statement, tagless `switch`), satu-satunya perulangan di Go (`for`), serta memahami konsep krusial **Pass-by-Value vs Pointer (`&` dan `*`)** di memori.

---

## 1. Mental Model: Kenapa Control Flow di Go Didesain Bersih?

Di JavaScript atau Dart, kamu punya:
- `for`, `while`, `do-while`, `for-in`, `for-of`, `forEach`
- `switch` yang butuh keyword `break` di setiap baris agar tidak bocor (*fallthrough*)
- Scope variabel yang sering bocor keluar blok percabangan

Di Golang, filosofinya adalah **minimalisme dan keamanan memori**:
1. **Hanya ada SATU kata kunci perulangan:** `for`. Tidak ada `while` atau `do-while`.
2. **`switch` tidak butuh `break`:** Go secara default otomatis berhenti di case yang cocok.
3. **`if` dengan *Init Statement*:** Variabel sementara hanya hidup di dalam blok `if-else`, menjaga namespace tetap bersih dari polusi variabel.
4. **Semua data di Go dioper *By Value* (duplikasi nilai):** Jika ingin mengubah data asli di memori, kamu **harus** menggunakan **Pointer**.

---

## 2. Percabangan: `if` dengan Short/Init Statement

Di bahasa lain, kamu biasanya menulis:
```javascript
// JS / Dart
let statusCode = checkStatus();
if (statusCode === 200) {
    // do something
}
// statusCode masih hidup dan mengotori scope di bawahnya
```

Di Go, ada pola sakral yang akan kamu gunakan di 90% codebase Go (terutama saat error handling):
```go
if statusCode := checkStatus(); statusCode == 200 {
    fmt.Println("Sukses!")
} else {
    fmt.Printf("Gagal dengan kode: %d\n", statusCode)
}
// statusCode HILANG dari memori di sini! Tidak bisa diakses lagi di luar if-else.
```

### Sintaks:
```text
if inisialisasi_variabel; kondisi_boolean {
    // blok kode jika true
}
```

---

## 3. Perulangan: Satu-satunya Loop di Go adalah `for`

Go membuang `while` dan `do-while`. Semua bentuk perulangan diselesaikan hanya dengan kata kunci `for`:

### A. Gaya Standar (3-Component Loop)
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### B. Gaya "While" (Hanya Kondisi)
```go
counter := 1
for counter <= 5 {
    fmt.Println(counter)
    counter++
}
```

### C. Gaya Infinite Loop (`for` tanpa henti)
Biasa dipakai untuk background worker, server listener, atau polling:
```go
for {
    // Loop selamanya sampai dipaksa berhenti dengan 'break'
    if conditionMet {
        break
    }
}
```

> [!TIP]
> Keyword `break` (menghentikan loop) dan `continue` (lompat ke iterasi berikutnya) bekerja persis seperti di JS/Dart.

---

## 4. `switch-case` Modern Tanpa `break`

### A. Auto-Break Bawaan
Di C/JS/Dart, lupa nulis `break` akan membuat eksekusi bocor ke case bawahnya. Di Go, `break` sudah otomatis!
```go
role := "admin"
switch role {
case "admin":
    fmt.Println("Full Access")
case "editor", "author": // Bisa multi-kondisi dengan koma!
    fmt.Println("Write Access")
default:
    fmt.Println("Read Only")
}
```

### B. Tagless Switch (Alternatif Pengganti `if-else if` Beruntun)
Kamu bisa membuat `switch` tanpa ekspresi di sampingnya. Setiap `case` akan dievaluasi sebagai kondisi boolean:
```go
score := 85
switch {
case score >= 90:
    fmt.Println("Grade A")
case score >= 75:
    fmt.Println("Grade B")
default:
    fmt.Println("Grade C")
}
```

---

## 5. Pondasi Memori: Apa itu Pointer (`&` dan `*`)?

Sebagai mobile/fullstack dev, konsep pointer sering dianggap menakutkan, padahal sebenarnya sangat sederhana:

### Analogi Nyata:
- **Variabel biasa** = Kamu memfotokopi dokumen dan membagikannya ke temanmu. Jika temanmu mencoret-coret salinannya, dokumen aslimu **tidak berubah**.
- **Pointer** = Kamu memberikan **koordinat GPS / link Google Docs** dari dokumen aslimu. Jika temanmu mengedit lewat link itu, dokumen aslimu **ikut berubah**!

### Dua Simbol Sakral Pointer:
1. **`&` (Address-Of Operator):** Digunakan untuk mengambil **alamat memori** suatu variabel.
2. **`*` (Dereference Operator):** Digunakan untuk **melihat atau mengubah isi nilai** di alamat memori yang ditunjuk.

```go
age := 25
var ptr *int = &age // ptr menyimpan alamat memori age (misal: 0xc000014070)

fmt.Println("Nilai age          :", age)  // 25
fmt.Println("Alamat memori age  :", &age) // 0xc000014070
fmt.Println("Isi ptr            :", ptr)  // 0xc000014070 (sama!)
fmt.Println("Nilai lewat pointer:", *ptr) // 25 (mengambil isi dari alamat memori)

// Mengubah nilai asli lewat pointer:
*ptr = 30
fmt.Println("Nilai age sekarang :", age)  // Berubah jadi 30!
```

---

## 6. Panduan Tugas Bertahap (Kerjakan di `day-02-control-flow-and-pointers/main.go`)

Buat file baru: `day-02-control-flow-and-pointers/main.go`. Kerjakan tantangan berikut secara bertahap:

---

### 🎯 LEVEL 1: Clean Scope dengan `if` Short Statement
**Instruksi:**
1. Di dalam `func main()`, buat simulasi respon HTTP request menggunakan `if` short statement:
   ```go
   if statusCode := 404; statusCode == 200 {
       fmt.Println("Status OK: Data berhasil diambil.")
   } else if statusCode == 404 {
       fmt.Printf("Not Found: Error kode %d\n", statusCode)
   } else {
       fmt.Printf("Server Error: Kode %d\n", statusCode)
   }
   ```
2. Coba letakkan baris `fmt.Println(statusCode)` di luar blok `if-else` tersebut, lalu perhatikan error kompiler apa yang muncul! *(Setelah melihat errornya, hapus/komentari baris itu untuk membuktikan prinsip scoped variable)*.

**Expected Output:**
```text
Not Found: Error kode 404
```

---

### 🎯 LEVEL 2: 3 Wajah Perulangan `for`
**Instruksi:**
Gunakan satu kata kunci `for` untuk menyelesaikan 3 misi:
1. **Misi A (Standard 3-Component):** Cetak angka `1` sampai `3` dalam satu baris.
2. **Misi B (While-style):** Mulai dari `countdown := 3`, lakukan perulangan selama `countdown > 0`, cetak nilainya, lalu kurangi `countdown--`. Di akhir cetak `"Misi Selesai!"`.
3. **Misi C (Infinite + Break/Continue):** Buat perulangan tak terbatas `for {}` dengan variabel penghitung `num := 0`:
   - Naikkan `num++` di setiap putaran.
   - Jika `num` adalah angka genap, lewati dengan `continue`.
   - Cetak `num` jika ganjil.
   - Jika `num` sudah mencapai `5`, hentikan loop dengan `break`.

**Expected Output:**
```text
Loop A: 1 2 3
Countdown: 3 2 1 ... Misi Selesai!
Odd Number: 1
Odd Number: 3
Odd Number: 5 (Break!)
```

---

### 🎯 LEVEL 3: Tagless `switch` Rule Engine
**Instruksi:**
Buat simulasi evaluasi diskon member belanja berdasarkan nominal belanja:
1. Buat variabel `totalBelanja := 250000` dan `isMember := true`.
2. Gunakan **tagless switch** (yaitu `switch { case ... : }` tanpa variabel di samping kata `switch`).
3. Kondisi case:
   - Jika `isMember` bernilai `true` DAN `totalBelanja >= 200000`: Diskon 20%.
   - Jika `isMember` bernilai `true` DAN `totalBelanja >= 100000`: Diskon 10%.
   - Jika bukan member tapi `totalBelanja >= 200000`: Diskon 5%.
   - Default: Tidak ada diskon.
4. Cetak hasil evaluasi diskon menggunakan `fmt.Println`.

**Expected Output:**
```text
Total Belanja: Rp 250000 | Status Member: true
Selamat! Anda mendapatkan diskon 20%
```

---

### 🎯 LEVEL 4 (Boss Level): Pointer Mutator Function
**Tujuan:** Membuktikan bagaimana fungsi mengubah data asli lewat memory address.

**Instruksi:**
1. Di luar fungsi `main()`, buat sebuah fungsi bernama `topUpEWallet`:
   ```go
   func topUpEWallet(balance *int, amount int) {
       // balance adalah pointer ke int (*int)
       // Gunakan *balance untuk menambah nilai di memori aslinya
       *balance += amount
   }
   ```
2. Di dalam `func main()`:
   - Buat variabel `myBalance := 50000`.
   - Cetak nilai awal dan alamat memorinya menggunakan verb `%d` dan `%p` (`&myBalance`).
   - Panggil fungsi `topUpEWallet(&myBalance, 100000)`. *(Perhatikan kita mengirimkan alamat memori dengan tanda `&`)*.
   - Cetak saldo setelah top-up. Buktikan bahwa nilai `myBalance` berubah menjadi `150000` meskipun tidak ada return value dari fungsi tersebut!

**Expected Output:**
```text
=== SIMULASI E-WALLET POINTER ===
Saldo Awal     : Rp 50000
Alamat Memori  : 0xc000...
Top up sebesar : Rp 100000...
Saldo Akhir    : Rp 150000 (Berhasil diubah lewat pointer!)
```

---

## 7. Cara Menjalankan Kode
Jalankan file dari root folder proyek:
```bash
go run ./day-02-control-flow-and-pointers/main.go
```

