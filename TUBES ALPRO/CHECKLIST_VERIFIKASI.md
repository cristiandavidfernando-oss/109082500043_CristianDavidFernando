# Checklist Verifikasi - MindFlow Application

## Spesifikasi Aplikasi

### A. Fungsionalitas Dasar ✓

#### 1. Manajemen Catatan Suasana Hati
- [x] **Menambahkan** data catatan suasana hati
  - [x] Input tanggal (format YYYY-MM-DD)
  - [x] Input skor emosi (1-10)
  - [x] Input deskripsi perasaan
  - [x] Validasi kapasitas penyimpanan

- [x] **Mengubah** data catatan suasana hati
  - [x] Menampilkan daftar mood yang tersedia
  - [x] Memilih mood untuk diubah
  - [x] Update semua field
  - [x] Konfirmasi perubahan

- [x] **Menghapus** data catatan suasana hati
  - [x] Menampilkan daftar mood yang tersedia
  - [x] Memilih mood untuk dihapus
  - [x] Pergeseran data otomatis
  - [x] Konfirmasi penghapusan

#### 2. Manajemen Daftar Tugas Harian
- [x] **Menambahkan** data tugas
  - [x] Input tanggal
  - [x] Input nama tugas
  - [x] Input durasi (menit)
  - [x] Input prioritas (1-3)
  - [x] Status default: belum selesai

- [x] **Mengubah** data tugas
  - [x] Menampilkan daftar task
  - [x] Memilih task untuk diubah
  - [x] Update semua field termasuk status
  - [x] Konfirmasi perubahan

- [x] **Menghapus** data tugas
  - [x] Menampilkan daftar task
  - [x] Memilih task untuk dihapus
  - [x] Pergeseran data otomatis
  - [x] Konfirmasi penghapusan

#### 3. Pencatatan Data
- [x] Skor emosi (skala 1-10)
- [x] Deskripsi perasaan
- [x] Nama tugas
- [x] Durasi pengerjaan (menit)
- [x] Status tugas (selesai/belum selesai)
- [x] Tanggal untuk semua data

### B. Algoritma Pencarian ✓

#### 1. Sequential Search - Pencarian Mood
- [x] Implementasi pencarian berdasarkan kata kunci
- [x] Scanning dari awal hingga akhir array
- [x] Menggunakan string.Contains untuk matching
- [x] Case-insensitive search
- [x] Menampilkan semua hasil yang cocok
- [x] Notifikasi jika tidak ditemukan

**Lokasi**: Function `cariMoodSequential()`
**Time Complexity**: O(n)
**Use Case**: Mencari mood dengan deskripsi tertentu

#### 2. Binary Search - Pencarian Task
- [x] Implementasi pencarian berdasarkan tanggal
- [x] Pre-processing: sorting dengan Bubble Sort
- [x] Eksekusi Binary Search pada array terurut
- [x] Membagi pencarian menjadi setengah setiap iterasi
- [x] Menampilkan hasil yang ditemukan
- [x] Notifikasi jika tidak ditemukan

**Lokasi**: Function `cariTaskBinary()`
**Time Complexity**: O(n log n) termasuk sorting
**Use Case**: Mencari task dengan tanggal spesifik

### C. Algoritma Pengurutan ✓

#### 1. Selection Sort - Urutkan Task by Prioritas
- [x] Mengurutkan dari prioritas tertinggi (1) ke terendah (3)
- [x] Mencari elemen minimum di setiap iterasi
- [x] Melakukan swap element
- [x] Menampilkan hasil terurut
- [x] Time Complexity: O(n²)

**Lokasi**: Function `urutTaskSelection()`
**Kriteria Sortir**: arrTask[j].Prioritas
**Output Format**: "- [Nama Tugas] (Prioritas: [X])"

#### 2. Insertion Sort - Urutkan Task by Durasi
- [x] Mengurutkan dari durasi terkecil (tercepat)
- [x] Menggeser elemen yang lebih besar
- [x] Menempatkan key pada posisi yang tepat
- [x] Menampilkan hasil terurut
- [x] Time Complexity: O(n²) worst case

**Lokasi**: Function `urutTaskInsertion()`
**Kriteria Sortir**: arrTask[j].Durasi
**Output Format**: "- [Nama Tugas] (Durasi: [X] menit)"

### D. Statistik dan Analitik ✓

#### 1. Tren Suasana Hati Mingguan
- [x] Menghitung rata-rata skor emosi
- [x] Formula: sum(semua skor) / jumlah mood
- [x] Output format: "X.XX / 10"
- [x] Handling: Jika tidak ada data = notifikasi

#### 2. Tingkat Penyelesaian Task Harian
- [x] Menghitung jumlah task selesai
- [x] Formula: (jumlah selesai / total task) × 100%
- [x] Output format: "X.XX%"
- [x] Handling: Jika tidak ada task = notifikasi

#### 3. Format Tampilan Statistik
- [x] Dekorasi dengan "+++" di sekitar nama aplikasi
- [x] Format: "+++ MindFlow +++"
- [x] Judul statistik: "++++ Statistik Tren Suasana Hati & Produktivitas ++++"
- [x] Penutup: "+++ MindFlow +++"

**Lokasi**: Function `tampilStatistik()`

### E. Struktur Data ✓

#### 1. Tipe Data Custom (Bonus 10 Poin)
```go
type eel int          // Menggantikan int
type jebb_24 string   // Menggantikan string
```

- [x] `eel` digunakan untuk: skor, durasi, prioritas, counter, loop index
- [x] `jebb_24` digunakan untuk: tanggal, nama tugas, deskripsi
- [x] Konsisten di seluruh aplikasi
- [x] Catatan: @ dihapus karena tidak valid dalam Go identifier

#### 2. Struct Mood
```go
type Mood struct {
    Tanggal   jebb_24  // Format: YYYY-MM-DD
    Skor      eel      // Skala 1-10
    Deskripsi jebb_24
}
```
- [x] Struktur sesuai spesifikasi
- [x] Menggunakan custom type
- [x] Komentarnya terattribusi

#### 3. Struct Task
```go
type Task struct {
    Tanggal   jebb_24
    Nama      jebb_24
    Durasi    eel      // Menit
    Prioritas eel      // 1-3
    Selesai   bool
}
```
- [x] Struktur sesuai spesifikasi
- [x] Menggunakan custom type
- [x] Kolom status penyelesaian
- [x] Komentarnya terattribusi

#### 4. Penyimpanan Data
- [x] Array statis: `var arrMood [MAX]Mood`
- [x] Array statis: `var arrTask [MAX]Task`
- [x] Kapasitas maksimum: MAX = 100
- [x] Counter: `countMood eel` dan `countTask eel`
- [x] Validasi kapasitas sebelum insert

### F. Persyaratan Komentar ✓

#### 1. Atribusi Penulis
- [x] Setiap komentar memiliki atribusi penulis
- [x] Format: `// komentar (nama_penulis)`
- [x] Contoh: `// melakukan perulangan sebanyak n kali (gabriel edbert)`
- [x] Total komentar terattribusi: 45+
- [x] Atribusi tersebar di seluruh fungsi

#### 2. Bonus 5 Poin: Akun/Kode dalam Source
- [x] Akun "@jebb_24" tertera di awal file
- [x] Kode "eel" digunakan di seluruh aplikasi
- [x] Dokumentasi: "Memenuhi syarat tugas tambahan 5 poin"

#### 3. Bonus 10 Poin: Penggantian Variabel
- [x] Minimal 2 variabel diganti dengan custom name
- [x] `int` → `eel` (lebih dari 2 penggunaan)
- [x] `string` → `jebb_24` (lebih dari 2 penggunaan)
- [x] Dokumentasi: "untuk 10 poin tambahan"

### G. Menu Aplikasi ✓

- [x] Menu 1: Tambah Mood
- [x] Menu 2: Ubah Mood
- [x] Menu 3: Hapus Mood
- [x] Menu 4: Tambah Task
- [x] Menu 5: Ubah Task
- [x] Menu 6: Hapus Task
- [x] Menu 7: Cari Mood (Sequential Search)
- [x] Menu 8: Cari Task (Binary Search)
- [x] Menu 9: Urutkan Task (Selection Sort)
- [x] Menu 10: Urutkan Task (Insertion Sort)
- [x] Menu 11: Tampilkan Statistik
- [x] Menu 0: Keluar

**Total Menu Item**: 12 fitur + exit

### H. Validasi Input ✓

- [x] Kapasitas penyimpanan tercek sebelum insert
- [x] Indeks menu valid di-validasi
- [x] Format tanggal bisa diparsing
- [x] Range emosi 1-10 bisa diparsing
- [x] Priority 1-3 bisa diparsing
- [x] Buffer cleansing setelah input number

### I. User Experience ✓

- [x] Tampilan menu yang jelas
- [x] Prompt input yang informatif
- [x] Konfirmasi sukses untuk setiap operasi
- [x] Pesan error untuk input tidak valid
- [x] Daftar data ditampilkan sebelum edit/delete
- [x] Instruksi yang mudah dipahami

## Ringkasan Kompletasi

### Fungsionalitas Inti
- ✓ CRUD Mood (Create, Read, Update, Delete)
- ✓ CRUD Task (Create, Read, Update, Delete)

### Algoritma
- ✓ Sequential Search
- ✓ Binary Search
- ✓ Selection Sort
- ✓ Insertion Sort

### Analitik
- ✓ Statistik Mood Mingguan
- ✓ Statistik Penyelesaian Task
- ✓ Format dengan dekorasi "+++"

### Bonus Point
- ✓ 5 Poin: Akun @jebb_24 dalam source code
- ✓ 10 Poin: 2+ variabel dengan custom type (eel, jebb_24)

### Dokumentasi
- ✓ Setiap komentar terattribusi
- ✓ README.md lengkap
- ✓ DOKUMENTASI_KOMENTAR.md detail
- ✓ File checklist ini

## Kompilasi dan Testing

- [x] Go code compiles successfully
- [x] No compilation errors
- [x] Executable file generated: mindflow.exe
- [x] Ready for deployment

## Status: ✓ LENGKAP

Semua persyaratan telah diimplementasikan dan diverifikasi. Aplikasi MindFlow siap digunakan.

---
**Tanggal Verifikasi**: 28 Mei 2026
**Status**: Approved for Submission
**Version**: 1.0 Final
