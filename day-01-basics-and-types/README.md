# DAY 01: Go Core Philosophy, Variables, Types, & The Go Compiler

Selamat datang di hari pertama! Di hari ini kita akan membangun pondasi yang kuat tentang cara kerja Golang, aturan sintaksnya, dan kenapa Go berbeda dari bahasa seperti JavaScript, TypeScript, atau Dart.

---

## 1. Mental Model: Kenapa Go Dibuat Berbeda?

Sebagai developer mobile atau fullstack:
- **Di TypeScript/JS:** Kamu punya `undefined`, `null`, `try-catch`, dan konversi tipe otomatis (misal: `"5" + 2` jadi `"52"`).
- **Di Dart:** Kamu punya class, inheritance, dan `sound null-safety`.
- **Di Golang:** 
  1. **Tidak ada class / inheritance** — Go memakai `struct` dan `composition`.
  2. **Tidak ada `try-catch`** — Error diperlakukan sebagai nilai biasa (*value*).
  3. **Kompiler sangat ketat (Zero Dead Code)** — Variabel atau package yang di-import tapi tidak dipakai akan langsung membuat program **gagal dikompilasi**.
  4. **Strict Typing tanpa kompromi** — Angka integer `int` **tidak bisa** otomatis ditambah dengan `float64`. Harus dikonversi manual secara eksplisit.

---

## 2. Anatomi Program Go (`package`, `import`, `main`)

Perhatikan susunan dasar file Go berikut:

```go
package main

import "fmt"

func main() {
    fmt.Println("Halo Golang!")
}
```

### Bedah Komponen:
1. **`package main`**:
   - Setiap file Go **wajib** dimulai dengan deklarasi package di baris pertama.
   - Package bernama `main` adalah penanda khusus bahwa file ini adalah **aplikasi yang bisa dieksekusi** (menghasilkan binary file), bukan sekadar library.
2. **`import "fmt"`**:
   - `fmt` adalah singkatan dari **format**. Package standar bawaan Go untuk urusan I/O teks (print ke terminal, format string, input dari user).
3. **Namespace Import (PENTING!)**:
   - Di Go, fungsi dari package yang di-import **tidak pernah masuk ke global scope**.
   - Kamu **harus** memanggilnya dengan format: `namaPackage.NamaFungsi()`.
   - Contoh: `fmt.Println()`, **bukan** `Println()`.
4. **Aturan Huruf Kapital (Visibility)**:
   - Go tidak punya kata kunci `public` atau `private`.
   - **Huruf depan Kapital (`Println`)** = Fungsi ini bersifat **Public (Exported)**, bisa dipakai dari luar package `fmt`.
   - **Huruf depan Kecil (`hitungTotal`)** = Fungsi bersifat **Private (Unexported)**, hanya bisa dipakai di dalam package yang sama.

---

## 3. Variabel: 3 Cara Deklarasi

Ada tiga cara membuat variabel di Go. Pahami kapan masing-masing digunakan:

### Cara 1: Explicit Declaration (Sebutkan tipe datanya secara gamblang)
Biasa digunakan jika kamu ingin memastikan tipe datanya jelas, atau saat belum ingin mengisi nilainya langsung.
```go
var age int = 25
var name string = "Alex"
```

### Cara 2: Type Inference (Tipe data ditebak otomatis oleh Go)
Kompiler Go pintar, dia tahu `"Alex"` adalah `string` dan `25` adalah `int`.
```go
var age = 25
var name = "Alex"
```

### Cara 3: Short Variable Declaration (`:=`) — *Paling Sering Dipakai!*
Kombinasi deklarasi + inisialisasi tanpa kata kunci `var`.
```go
age := 25
name := "Alex"
```
> ⚠️ **Aturan Penting `:=`**:
> - Operator `:=` **hanya bisa digunakan di dalam tubuh fungsi** (`func main() { ... }`).
> - Di luar fungsi (tingkat package/global), kamu **wajib** pakai `var` atau `const`.

---

## 4. Konsep "Zero Value" di Memori

Di Go, **tidak ada konsep variabel bernilai `undefined` atau garbage value**. 

Jika kamu membuat variabel tanpa mengisi nilai awalnya:
```go
var count int
var message string
var isReady bool
```
Go otomatis mengalokasikan nilai default (*Zero Value*) di memori:
| Tipe Data | Zero Value | Keterangan |
| :--- | :--- | :--- |
| `int`, `int32`, `int64` | `0` | Angka nol |
| `float32`, `float64` | `0` | Angka nol desimal (`0.0`) |
| `bool` | `false` | Boolean false |
| `string` | `""` | String kosong (bukan null!) |
| Pointer, Slice, Map | `nil` | Belum menunjuk ke alamat memori manapun |

---

## 5. Bedah Tuntas `fmt.Printf` (Formatting Template)

Seringkali kita bingung bedanya `Println` dan `Printf`:
- **`fmt.Println(...)`**: Mencetak nilai apa adanya, otomatis memberi spasi antar argumen, dan **otomatis ganti baris** di akhir.
- **`fmt.Printf(...)`**: Huruf **f** artinya *formatted*. Digunakan untuk membuat template kalimat dengan placeholder. `Printf` **TIDAK otomatis ganti baris**, jadi kamu harus menambahkan `\n` di akhir string.

### Anatomi `fmt.Printf`:
```go
fmt.Printf("Template dengan %placeholder di sini\n", nilaiVariabel)
```

Urutan nilai setelah tanda koma harus **cocok** dengan urutan placeholder di dalam template string:
```go
nama := "Budi"
umur := 20
fmt.Printf("Halo, nama saya %s dan umur saya %d tahun.\n", nama, umur)
// %s akan digantikan oleh nama ("Budi")
// %d akan digantikan oleh umur (20)
```

### Tabel Placeholder (Verbs) Wajib Tahu:
| Verb | Kegunaan | Contoh Output |
| :--- | :--- | :--- |
| `%s` | String (teks biasa) | `"Budi"` |
| `%d` | Decimal Integer (angka bulat) | `25` |
| `%f` | Float (desimal standar, 6 angka di belakang koma) | `3.140000` |
| `%.2f` | Float dengan pembatasan (contoh: 2 desimal) | `3.14` |
| `%t` | Boolean (`true` / `false`) | `true` |
| `%T` | **Tipe Data** dari variabel (huruf T kapital) | `string`, `int`, `bool` |
| `%q` | String diapit tanda kutip ganda (sangat berguna untuk melihat string kosong `""`) | `"teks"` atau `""` |
| `%v` | Format default (serbaguna untuk tipe apapun) | apa saja |

---

## 6. Strict Typing & Type Conversion

Go tidak memiliki *type coercion* (tidak ada konversi otomatis).
```go
var harga int = 10000
var diskon float64 = 0.5

// ❌ COMPILE ERROR:
// total := harga * diskon  (invalid operation: mismatched types int and float64)

// ✅ CARA BENAR: Konversi eksplisit dengan rumus TipeTujuan(variabelAsal)
total := float64(harga) * diskon // Hasilnya 5000 (float64)
```

---

## 7. Konstanta & `iota`

Konstanta dibuat menggunakan keyword `const`. Nilainya tidak bisa diubah setelah di-assign.
Go menyediakan fitur bernama `iota` untuk membuat enum / urutan angka otomatis:
```go
const (
    StatusPending = iota // bernilai 0
    StatusActive         // bernilai 1 (otomatis naik)
    StatusFailed         // bernilai 2 (otomatis naik)
)
```

---

## 8. Panduan Tugas Bertahap (Kerjakan di `main.go`)

Buka file `day-01-basics-and-types/main.go`. Kerjakan satu per satu dari Level 1 sampai Level 4.

---

### 🎯 LEVEL 1: Perkenalan & Eksplorasi `fmt.Printf`

**Instruksi Langkah demi Langkah:**
1. Di dalam `func main()`, buat 3 variabel berikut menggunakan gaya short declaration (`:=`):
   - `name`: isi dengan nama kamu (contoh: `"Alex"`).
   - `age`: isi dengan umur kamu dalam bentuk angka (contoh: `25`).
   - `isMobileDev`: isi dengan boolean (`true` atau `false`).
2. Cetak kalimat perkenalan menggunakan `fmt.Printf` dengan placeholder `%s`, `%d`, dan `%t`.  
   *Ingat: tambahkan `\n` di ujung string agar ganti baris!*
3. Cetak baris kedua yang menampilkan **tipe data** dari masing-masing variabel tersebut menggunakan placeholder `%T`.

**Contoh Tampilan Output Terminal yang Diharapkan:**
```text
Halo! Nama saya Alex, umur 25 tahun. Mobile Developer: true
Tipe data -> name: string | age: int | isMobileDev: bool
```

---

### 🎯 LEVEL 2: Eksperimen Zero Value Audit

**Tujuan:** Membuktikan bahwa Go tidak punya nilai `undefined` di memori.

**Instruksi Langkah demi Langkah:**
1. Buat 4 variabel menggunakan kata kunci `var` **tanpa mengisi nilainya sama sekali** (biarkan kosong):
   - `userLevel` dengan tipe `int`
   - `accountBalance` dengan tipe `float64`
   - `isVerified` dengan tipe `bool`
   - `userBio` dengan tipe `string`
2. Cetak nilai default keempat variabel tersebut menggunakan `fmt.Printf`.
   - Untuk `userLevel`, `accountBalance`, dan `isVerified`, gunakan placeholder `%v`.
   - Khusus untuk `userBio`, gunakan placeholder `%q` (agar kamu bisa melihat tanda kutip `""` yang membuktikan isinya adalah empty string, bukan null).

**Contoh Tampilan Output Terminal yang Diharapkan:**
```text
Zero Values -> Level: 0 | Balance: 0 | Verified: false | Bio: ""
```

---

### 🎯 LEVEL 3: Strict Conversion & Perhitungan Pajak

**Tujuan:** Melatih penanganan konversi tipe data yang ketat tanpa error kompilasi.

**Instruksi Langkah demi Langkah:**
1. Buat variabel `itemPrice` bertipe `int` dengan nilai `150000`.
2. Buat variabel `taxRate` bertipe `float64` dengan nilai `0.11` (pajak 11%).
3. Hitung besaran pajak (`taxAmount`) dengan rumus: `itemPrice * taxRate`.  
   *(Perhatian: `itemPrice` harus dikonversi ke `float64` dulu: `float64(itemPrice) * taxRate`)*.
4. Hitung harga akhir (`totalPrice`) dalam bentuk bilangan bulat `int` dengan rumus: `itemPrice + int(taxAmount)`.
5. Cetak rincian tersebut menggunakan `fmt.Printf`. Tampilkan pajak dengan format `%.2f` (dua desimal di belakang koma).

**Contoh Tampilan Output Terminal yang Diharapkan:**
```text
Harga Barang : Rp 150000
Pajak (11%)  : Rp 16500.00
Total Akhir  : Rp 166500
```

---

### 🎯 LEVEL 4 (Tantangan Akhir Day 01): Mini CLI Stats Card

**Tujuan:** Menggabungkan variabel, `const`, `iota`, dan formatting teks rapi.

**Instruksi Langkah demi Langkah:**
1. Di luar fungsi `main()` (di tingkat package), buat enum level menggunakan `const` dan `iota`:
   ```go
   const (
       RoleJunior = iota // 0
       RoleMid           // 1
       RoleSenior         // 2
   )
   ```
2. Di dalam `func main()`, buat beberapa variabel profil dengan mencampur cara deklarasi:
   - Gunakan `var` eksplisit untuk `experienceYears int = 3`
   - Gunakan short declaration `:=` untuk `dailyRate := 45.5` (tipe float)
   - Tentukan level kamu menggunakan salah satu konstanta di atas, misal: `currentRole := RoleMid`
3. Cetak kartu profil rapi di terminal menggunakan `fmt.Printf` dan pemisah garis.

**Contoh Tampilan Output Terminal yang Diharapkan:**
```text
========================================
DEVELOPER STATS CARD
========================================
Name        : Alex
Role Level  : 1 (0: Junior, 1: Mid, 2: Senior)
Experience  : 3 Tahun (Tipe: int)
Hourly Rate : $45.50 (Tipe: float64)
Status 90D  : Semangat Belajar Golang!
========================================
```

---

## 9. Cara Menjalankan Kode Kamu

Setiap kali kamu selesai mengetik kode di `day-01-basics-and-types/main.go`, buka terminal dan jalankan:
```bash
go run ./day-01-basics-and-types/main.go
```
Jika ada error, baca pesan kompilernya baik-baik, atau tanyakan langsung ke mentor!
