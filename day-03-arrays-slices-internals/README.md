# DAY 03: Arrays, Slices, and Memory Internals

> **Target:** Membongkar arsitektur memori di balik struktur data koleksi Go: perbedaan mutlak antara **Array** (fixed value) dan **Slice** (dynamic view), anatomi 3 komponen **Slice Header** (`pointer`, `len`, `cap`), jebakan modifikasi sub-slice pada *backing array*, serta rahasia perulangan `for range`.

---

## 1. Mental Model: Array di JS/Dart vs Slice di Go

Di dunia JavaScript (`Array`) atau Dart (`List`):
- `const list = [1, 2, 3]` adalah objek dinamis di heap yang otomatis membesar saat kamu `.push()` atau `.add()`.
- Variabel `list` adalah referensi objek.

Di Golang, ada **Dua Tingkatan Koleksi**:
1. **Array**: Ukurannya **statis (permanen)** dan ukuran adalah **bagian dari tipe data**. Array di Go adalah *value type* murni. Mengoper `[1000]int` ke fungsi berarti menduplikasi 1.000 angka di memori!
2. **Slice**: Struktur data pembungkus (*wrapper*) tipis di atas array. Ukurannya dinamis, sangat murah dioper (hanya 24 byte di mesin 64-bit), dan merupakan tipe koleksi yang akan kamu gunakan di **99% kasus di Go**.

---

## 2. Array vs Slice (Perbedaan Sintaks)

Perhatikan tanda kurung sikunya:
```go
// 1. ARRAY: Jumlah elemen ditulis di dalam kurung siku
var a [3]int = [3]int{10, 20, 30}
b := [...]string{"Go", "Rust", "Dart"} // Kompiler menghitung ukuran otomatis: [3]string

// 2. SLICE: Kurung siku DIBIARKAN KOSONG
var s []int = []int{10, 20, 30}
frameworks := []string{"React", "Next.js", "Flutter"}
```

> [!IMPORTANT]
> Tipe `[3]int` dan `[5]int` adalah **dua tipe data yang sama sekali berbeda**. Kamu tidak bisa mengoper `[5]int` ke fungsi yang menerima `[3]int`. Itulah kenapa Array murni jarang dipakai sebagai parameter fungsi.

---

## 3. Anatomi Jeroan Slice: The 24-Byte Slice Header

Ketika kamu membuat sebuah slice, Go **tidak** menyimpan elemen langsung di variabel tersebut. Di balik layar, Go membuat struct kecil berukuran 24 byte (*Slice Header*):

```text
+-----------------------+
|  Pointer (*T)         | ---> Menunjuk ke Backing Array di memori
+-----------------------+
|  Length (int)         | ---> Jumlah elemen aktif yang terbaca (len)
+-----------------------+
|  Capacity (int)       | ---> Kapasitas maksimum dari titik pointer (cap)
+-----------------------+
```

```go
s := make([]int, 3, 5) // len = 3, cap = 5
```
- **`len(s)`**: Berapa banyak elemen yang saat ini bisa kamu akses (`s[0]`, `s[1]`, `s[2]`).
- **`cap(s)`**: Berapa banyak ruang yang tersedia di *backing array* sebelum Go terpaksa memindahkan data ke memori baru.

---

## 4. Cara Kerja `append()` & Relokasi Memori

Kenapa kita selalu menulis:
```go
numbers = append(numbers, 10)
```
dan bukan `numbers.append(10)`?

Karena:
1. **Jika `len < cap`:** Go langsung menulis elemen baru ke *backing array* yang sudah ada, lalu menaikkan nilai `len` pada header.
2. **Jika `len == cap` (Kapasitas Penuh):** 
   - Go mengalokasikan *backing array* baru yang lebih besar (biasanya **2x lipat** dari ukuran lama).
   - Go menyalin (*copy*) semua data lama ke array baru tersebut.
   - Mengembalikan **Slice Header baru** dengan pointer yang menunjuk ke alamat array baru.
   - Array lama yang ditinggalkan akan dibersihkan oleh Garbage Collector.

Itulah kenapa hasil dari `append()` **harus ditangkap kembali** ke variabel slice!

---

## 5. Jebakan Sakral: Sub-Slicing & Backing Array Bersama

Ketika kamu memotong slice: `sub := original[1:4]`, Go **TIDAK MENDUPLIKASI DATA**. 
`sub` dan `original` menggunakan **BACKING ARRAY YANG SAMA DI RAM**.

```go
original := []int{10, 20, 30, 40, 50}
sub := original[1:3] // [20, 30]

sub[0] = 999 // Mengubah elemen pertama sub

fmt.Println(original) // [10, 999, 30, 40, 50] -> IKUT BERUBAH!
```

### Solusi Pemisahan Data: Fungsi `copy()`
Jika kamu ingin potongan slice yang independen dan tidak saling mempengaruhi:
```go
subIndependent := make([]int, len(original[1:3]))
copy(subIndependent, original[1:3]) // Menduplikasi byte ke array baru
```

---

## 6. Jebakan Perulangan `for range`

Ketika meloop slice dengan `range`:
```go
numbers := []int{10, 20, 30}

for i, v := range numbers {
    v = v * 2 // ❌ TIDAK AKAN MENGUBAH ISI SLICE 'numbers'!
}
```

> [!WARNING]
> Variabel `v` di dalam `for range` adalah **COPY / DUPLIKASI** dari elemen tersebut. Mengubah `v` tidak ada efeknya pada slice asli.
> 
> **Cara Benar memodifikasi data via index:**
> ```go
> for i := range numbers {
>     numbers[i] *= 2 // ✅ Mengubah memori asli lewat index!
> }
> ```

---

## 7. Panduan Tugas Bertahap (Kerjakan di `day-03-arrays-slices-internals/day3.go`)

Buat folder `day-03-arrays-slices-internals` dan file `day3.go`. Kerjakan tantangan bertahap berikut:

---

### 🎯 LEVEL 1: Array vs Slice & The `append` Return
**Instruksi:**
1. Buat sebuah fixed array `mobileTech` dengan tipe `[3]string` berisi `"Flutter"`, `"React Native"`, `"Swift"`.
   - Cetak isi array dan cetak tipe datanya menggunakan `%T` (`[3]string`).
2. Buat sebuah slice `backendTech` dengan nilai awal `[]string{"Go", "Node.js"}`.
   - Tambahkan `"PostgreSQL"` menggunakan fungsi `append()`.
   - Cetak isi slice dan tipe datanya menggunakan `%T` (`[]string`).
   - Tulis komentar di kodemu: apa beda tipe data `[3]string` vs `[]string`?

**Expected Output:**
```text
Array: [Flutter React Native Swift] | Tipe: [3]string
Slice: [Go Node.js PostgreSQL] | Tipe: []string
```

---

### 🎯 LEVEL 2: Eksperimen Detektif Memori (`len`, `cap`, & Relokasi Pointer)
**Tujuan:** Melihat dengan mata kepala sendiri momen tepat saat Go memindahkan array ke alamat memori baru karena kehabisan kapasitas.

**Instruksi:**
1. Inisialisasi slice angka menggunakan `make`:
   ```go
   nums := make([]int, 0, 2) // len = 0, cap = 2
   ```
2. Lakukan perulangan `for i := 1; i <= 5; i++`:
   - Di setiap putaran, `append` nilai `i * 10` ke `nums`.
   - Cetak nilai `len(nums)`, `cap(nums)`, dan alamat pointer backing array-nya menggunakan verb `%p`:
     ```go
     fmt.Printf("Iterasi %d -> len: %d | cap: %d | ptr: %p\n", i, len(nums), cap(nums), nums)
     ```
3. Perhatikan output terminal di iterasi ke-3: kapasitas melompat dari 2 ke 4, dan alamat memori pointer `%p` **berubah**!

**Expected Output:**
```text
Iterasi 1 -> len: 1 | cap: 2 | ptr: 0xc000...
Iterasi 2 -> len: 2 | cap: 2 | ptr: 0xc000... (ptr sama)
Iterasi 3 -> len: 3 | cap: 4 | ptr: 0xc000... (ALAMAT MEMORI BERUBAH / RELOKASI!)
Iterasi 4 -> len: 4 | cap: 4 | ptr: 0xc000...
Iterasi 5 -> len: 5 | cap: 8 | ptr: 0xc000... (RELOKASI LAGI!)
```

---

### 🎯 LEVEL 3: The Sub-Slice Trap & The `copy()` Antidote
**Tujuan:** Mengetes kebocoran mutasi data akibat *shared backing array* dan cara menetralkannya.

**Instruksi:**
1. Buat slice: `original := []int{100, 200, 300, 400, 500}`.
2. Buat sub-slice `dangerSlice := original[1:4]` (mengambil `200, 300, 400`).
3. Ubah elemen pertama `dangerSlice[0] = 999`.
4. Cetak `original` dan `dangerSlice`. Buktikan bahwa angka `200` pada `original` ikut rusak berubah jadi `999`!
5. Sekarang buat klon independen dari `original[1:4]` bernama `safeSlice` menggunakan `make` dan fungsi `copy(safeSlice, ...)`.
6. Ubah `safeSlice[0] = 777`.
7. Cetak `original` dan buktikan bahwa `original` **TIDAK** berubah menjadi `777`!

**Expected Output:**
```text
=== UJI SHARE BACKING ARRAY ===
Original setelah dangerSlice diubah: [100 999 300 400 500] (Tercemar!)
=== UJI COPY INDEPENDEN ===
SafeSlice diubah : [777 300 400]
Original aman    : [100 999 300 400 500] (Aman tidak terpengaruh!)
```

---

### 🎯 LEVEL 4 (Boss Level): E-Commerce Cart Discounter & Filter
**Tujuan:** Mengombinasikan `for range` mutasi via indeks, dan pembuatan slice hasil filter tanpa kebocoran memori.

**Instruksi:**
1. Buat slice harga produk keranjang belanja:
   ```go
   cartPrices := []float64{250000, 50000, 120000, 30000, 750000, 15000}
   ```
2. **Misi A (In-place Mutation):** Berikan diskon 10% untuk **semua barang**. 
   - Gunakan perulangan `for index := range cartPrices`.
   - Ubah langsung harga di memori aslinya: `cartPrices[index] *= 0.9`.
   - *(Ingat: jangan pakai `for _, price := range cartPrices { price *= 0.9 }` karena itu cuma copy!)*
3. **Misi B (Filtering):** Saring produk yang harganya di atas Rp 100.000 setelah diskon ke dalam slice baru bernama `premiumItems []float64`.
   - Gunakan `append` untuk memasukkan item yang lolos kualifikasi.
4. Cetak rincian keranjang belanja setelah diskon dan daftar item premium yang tersaring.

**Expected Output:**
```text
=== KERANJANG BELANJA (DISKON 10%) ===
Harga Terbaru: [225000 45000 108000 27000 675000 13500]
Barang Premium (> Rp 100.000): [225000 108000 675000] (Total: 3 item)
```

---

## 8. Cara Menjalankan Kode
Jalankan file dari root folder:
```bash
go run ./day-03-arrays-slices-internals/day3.go
```
*(atau `go run ./day-03-arrays-slices-internals`)*
