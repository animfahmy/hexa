# 🗺️ Hexa Project Documentation & Cheat Sheet

Dokumen ini adalah panduan cepat untuk mengingat kembali konsep dasar Golang, konfigurasi, dan arsitektur Hexagonal yang diterapkan pada proyek ini.
[Link](https://gemini.google.com/u/3/app/e43d6b09a2b0fdbe?pageId=none)

---

## 💡 1. Golang Crucial Basics (Fase 1 & 2)

### Defer (Menunda Eksekusi)
`defer` menjamin sebuah perintah **pasti dijalankan di akhir fungsi**, bahkan jika fungsi tersebut keluar lebih cepat (*early return* karena error) atau terjadi *panic* (crash).
*   **⚠️ Aturan Emas:** Selalu gunakan `defer` tepat setelah berhasil membuka resource (Database, File, Network Connection).
```go
file.Open("data.txt")
defer file.Close() // Pasti aman, mencegah Memory Leak!

```

### Nil vs Null

Go tidak mengenal `null`, melainkan **`nil`**. Tipe data dasar (`string`, `int`) tidak bisa bernilai `nil` (default: `""` atau `0`). `nil` hanya berlaku untuk:

* Pointer (`*struct`)
* Interface
* Slice, Map, dan Channel

### Map & Slices Iteration

Cara tercepat membaca data dari Slice atau Map adalah menggunakan `for range`:

```go
// Map penanda format: %s untuk string, %d untuk integer, %v untuk universal
ports := map[string]int{"http": 80}
for key, value := range ports {
    fmt.Printf("Protokol: %s, Port: %d\n", key, value)
}

```

---

## 🛡️ 2. Advance Concurrency & Context

### Race Condition & Mutex (`sync.Mutex`)

Terjadi ketika banyak *Goroutines* (proses asinkronus) mengubah satu variabel yang sama secara bersamaan. Gunakan `mutex` untuk mengunci akses.

```go
mutex.Lock()
counter++ // Proses dijamin aman dari bentrokan data
mutex.Unlock()

```

### Context Timeout

Backend tidak boleh membiarkan proses/query menggantung selamanya. Gunakan `context.WithTimeout` untuk membatasi waktu eksekusi.

* `context.Background()`: Akar context kosong sebagai pondasi awal.
* `context.WithTimeout()`: Memasang bom waktu otomatis pada proses (misal maksimal 2 detik).

---

## 🛠️ 3. Environment & Configuration (Fase 3)

* **Aturan:** *Credential* sensitif (Password, DSN, Port) dilarang keras ditulis langsung di kode (*hardcoded*).
* **Mekanisme:** Ditulis di file `.env`, di-load oleh library `godotenv`, kemudian dipetakan ke dalam sebuat objek struct di `internal/config/config.go` agar terpusat.

---

## 📐 4. Hexagonal Architecture Blueprint (Fase 4)

Tujuan utama arsitektur ini adalah memisahkan **Logika Bisnis (Core)** dari **Detail Teknis (Database, HTTP Framework, Library Pihak Ketiga)**.

### Struktur Folder & Aliran Data

```text
internal/
├── core/
│   ├── domain/     # Entitas / Struct murni tanpa dependensi eksternal.
│   └── ports/      # INTERFACE (Kontrak). Jembatan komunikasi inbound & outbound.
│   └── service/    # LOGIKA BISNIS (Menerapkan Driving Port, memanggil Driven Port).
└── adapters/
    ├── handler/    # PRIMARY ADAPTER: Menerima input luar (REST API Echo, CLI, gRPC).
    └── repository/ # SECONDARY ADAPTER: Eksekusi teknis luar (GORM, MySQL, Postgres).

```

### Alur Kerja Komunikasi (Kunci Pemahaman)

1. **Request Masuk:** Masuk via HTTP (Echo) di `adapters/handler`.
2. **Panggil Core:** Handler memanggil method milik `core/service` melalui interface **Driving Port**.
3. **Eksekusi Bisnis:** Service mengolah logika bisnis. Jika butuh data, Service memanggil interface **Driven Port** (`ProductRepository`).
4. **Akses Database:** Interface Driven Port tersebut digerakkan oleh `adapters/repository` yang berisi query asli GORM ke MySQL.

### 🧩 Rumus Dependency Injection Manual (`cmd/main.go`)

Aplikasi dirakit dari luar ke dalam seperti Lego:

```go
repo := repository.NewGormRepository(db)   // 1. Buat Adapter DB
svc  := service.NewProductService(repo)    // 2. Suntikkan Repo ke Core Service
handler.NewHttpProductHandler(e, svc)      // 3. Suntikkan Service ke HTTP Handler

```

---

## 🔍 5. Cara Cepat Membaca Error Go (Stack Trace)

Jika terminal memunculkan dinding teks *panic / crash*, abaikan baris library luar, lalu terapkan **Rumus 3 Langkah**:

1. **Lihat Paling Atas:** Cari tipe errornya (misal: `runtime error: invalid memory address or nil pointer dereference`).
2. **Cari Modul Anda:** Tarik mata ke bawah, cari baris pertama yang diawali nama folder Anda (`hexa/...`). Lihat nama file dan nomor barisnya.
3. **Analisis Pointer:** Jika tertulis *nil pointer*, periksa apakah ada variabel object/koneksi database yang belum terinisialisasi tetapi sudah buru-buru dipanggil fungsi atau method-nya.

