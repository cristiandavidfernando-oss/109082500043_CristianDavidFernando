# MindFlow - Aplikasi Asisten Virtual Kesehatan Mental dan Produktivitas

## Deskripsi
MindFlow adalah aplikasi asisten virtual yang membantu pengguna menjaga keseimbangan emosional dan efisiensi kerja. Aplikasi ini dirancang untuk mahasiswa, pekerja kantoran, dan individu yang ingin menjaga kesejahteraan mental mereka.

## Fitur Utama

### 1. Manajemen Catatan Suasana Hati (Mood)
- **Tambah Mood**: Mencatat tanggal, skor emosi (1-10), dan deskripsi perasaan
- **Ubah Mood**: Mengupdate catatan mood yang sudah ada
- **Hapus Mood**: Menghapus catatan mood tertentu
- **Cari Mood**: Pencarian dengan Sequential Search berdasarkan kata kunci

### 2. Manajemen Daftar Tugas Harian (Task)
- **Tambah Task**: Menambahkan tugas dengan tanggal, nama, durasi (menit), dan prioritas
- **Ubah Task**: Mengupdate tugas termasuk status penyelesaian
- **Hapus Task**: Menghapus tugas dari daftar
- **Tandai Selesai**: Menandai tugas sebagai completed/incomplete

### 3. Pencarian Data
- **Sequential Search (Mood)**: Mencari mood berdasarkan kata kunci pada deskripsi
- **Binary Search (Task)**: Mencari tugas berdasarkan tanggal (data otomatis diurutkan)

### 4. Pengurutan Data
- **Selection Sort (Priority)**: Mengurutkan task berdasarkan prioritas (1-3)
- **Insertion Sort (Duration)**: Mengurutkan task berdasarkan durasi pengerjaan

### 5. Statistik dan Analitik
- **Tren Suasana Hati Mingguan**: Menampilkan rata-rata skor emosi
- **Tingkat Penyelesaian Task**: Menampilkan persentase task yang sudah selesai
- **Format Khusus**: Ditampilkan dengan dekorasi "+++" di sekitar nama aplikasi

## Struktur Data

### Tipe Data Custom
```go
type eel int           // Menggantikan tipe int (Bonus 10 poin)
type jebb_24 string    // Menggantikan tipe string (Bonus 10 poin)
```

### Struct Mood
- **Tanggal**: Format YYYY-MM-DD
- **Skor**: Skala 1-10
- **Deskripsi**: Deskripsi perasaan pengguna

### Struct Task
- **Tanggal**: Format YYYY-MM-DD
- **Nama**: Nama tugas
- **Durasi**: Waktu pengerjaan dalam menit
- **Prioritas**: 1 (Tinggi), 2 (Sedang), 3 (Rendah)
- **Selesai**: Status penyelesaian tugas (boolean)

## Menu Aplikasi
```
=== MindFlow: Asisten Kesehatan Mental & Produktivitas ===
1. Tambah Mood
2. Ubah Mood
3. Hapus Mood
4. Tambah Task
5. Ubah Task
6. Hapus Task
7. Cari Mood (Sequential Search - By Keyword)
8. Cari Task (Binary Search - By Tanggal)
9. Urutkan Task (Selection Sort - By Prioritas)
10. Urutkan Task (Insertion Sort - By Durasi)
11. Tampilkan Statistik
0. Keluar
```

## Panduan Penggunaan

### Menjalankan Aplikasi
```bash
cd "path/ke/TUBES ALPRO"
./mindflow.exe
```

### Contoh Penggunaan Fitur

#### 1. Menambah Mood
- Pilih menu 1
- Masukkan tanggal (contoh: 2024-05-28)
- Masukkan skor emosi 1-10
- Masukkan deskripsi perasaan

#### 2. Menambah Task
- Pilih menu 4
- Masukkan tanggal tugas
- Masukkan nama tugas
- Masukkan durasi dalam menit
- Masukkan prioritas (1-3)

#### 3. Mencari Mood
- Pilih menu 7
- Masukkan kata kunci (contoh: "sedih", "senang")
- Sistem akan menampilkan semua mood yang cocok dengan kata kunci

#### 4. Mencari Task
- Pilih menu 8
- Masukkan tanggal yang dicari
- Data akan otomatis diurutkan berdasarkan tanggal sebelum pencarian binary search

## Persyaratan Bonus

### 5 Poin Tambahan ✓
- Kode atau akun (@jebb_24) terdapat dalam source code
- Lokasi: Di awal file dan dalam nama tipe data

### 10 Poin Tambahan ✓
- Minimal 2 variabel diganti dengan nama custom:
  - `int` → `eel` (digunakan di seluruh aplikasi)
  - `string` → `jebb_24` (digunakan di seluruh aplikasi)

### Dokumentasi Kode ✓
- Setiap komentar dilengkapi dengan nama penulis: `// comment (NamaKamu)`
- Format standar untuk identifikasi apakah kode ditulis dengan/tanpa AI

## Kapasitas Penyimpanan
- Maximum 100 data mood
- Maximum 100 data task
- Total slots array: 200 data

## Algoritma Pencarian
- **Sequential Search**: O(n) - cocok untuk data kecil
- **Binary Search**: O(log n) - memerlukan data terurut

## Algoritma Pengurutan
- **Selection Sort**: O(n²) - selalu mencari elemen minimum
- **Insertion Sort**: O(n²) worst case - efisien untuk data sebagian terurut

## Teknologi
- **Bahasa**: Go 1.x
- **Platform**: Windows/Linux/macOS
- **Input**: Console-based dengan scanner (mendukung input dengan spasi)

## Struktur File Project
```
TUBES ALPRO/
├── mindflow.go      (File utama aplikasi)
├── mindflow.exe     (Executable file)
├── README.md        (Dokumentasi ini)
├── ais.go           (File tambahan)
├── chandra.go       (File tambahan)
├── coba.go          (File tambahan)
└── rian.go          (File tambahan)
```

## Catatan Pengembangan
- Semua komentar dalam bahasa Indonesia sesuai dengan spesifikasi
- Menggunakan array statis dengan ukuran maksimal 100
- Implementasi dilakukan tanpa library eksternal tambahan
- Output user-friendly dengan pesan konfirmasi untuk setiap operasi

## Status Implementasi
Semua spesifikasi telah diimplementasikan dan diuji:
- ✓ CRUD Mood (Create, Read, Update, Delete)
- ✓ CRUD Task (Create, Read, Update, Delete)
- ✓ Sequential Search
- ✓ Binary Search
- ✓ Selection Sort
- ✓ Insertion Sort
- ✓ Statistik dan Analitik
- ✓ Bonus points fulfilment

---
*Aplikasi MindFlow v1.0 - Untuk Menjaga Kesejahteraan Mental dan Produktivitas*
