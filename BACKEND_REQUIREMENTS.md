# Backend Requirements — BBS Conveyor

Panduan untuk backend Go (chi + pgx/v5 + PostgreSQL).

---

## 1. Response Shape

Semua endpoint return JSON dengan format:

```json
// Sukses
{ "data": { ... }, "error": null }

// Error
{ "data": null, "error": "pesan error" }
```

HTTP status code tetap sesuai (200/201 sukses, 400/401/404/500 error).

---

## 2. Auth

JWT via HttpOnly cookie. Flow:

```
POST /api/admin/login  →  set cookie "token" (HttpOnly, Secure, SameSite=Lax)
POST /api/admin/logout →  clear cookie
GET  /api/admin/me     →  return user info dari JWT
```

Semua path `/api/admin/*` (kecuali login) wajib validasi JWT middleware.

---

## 3. Endpoint Admin (yang perlu dibangun — urut prioritas)

### 3.1 Produk

| Method | Path | Keterangan |
|--------|------|------------|
| `GET`    | `/api/admin/produk`      | List semua. Query params opsional: `?group=belt-conveyor&kategori=pvc-belt` |
| `GET`    | `/api/admin/produk/:id`  | Detail satu produk |
| `POST`   | `/api/admin/produk`      | Create. Body JSON, field `image` berupa string URL |
| `PUT`    | `/api/admin/produk/:id`  | Update |
| `DELETE` | `/api/admin/produk/:id`  | Hapus |

Shape (mengikuti `types/product.ts`):

```json
{
  "id": "uuid",
  "slug": "pvc-belt-hijau-2ply",
  "name": "PVC Belt Hijau 2-Ply",
  "group": "belt-conveyor",
  "kategori": "pvc-belt",
  "category": "PVC Belt",
  "description": "Deskripsi singkat (untuk card)",
  "detail": "Deskripsi lengkap (untuk halaman detail)",
  "image": "/images/product/pvc-belt-hijau.jpg",
  "specs": {
    "Ketebalan": "2 mm",
    "Warna": "Hijau",
    "Ply": "2"
  }
}
```

- `group`: enum — `"belt-conveyor"` atau `"lainnya"`
- `kategori`: slug kategori (contoh: `"pvc-belt"`, `"timing-belt"`)
- `specs`: JSON object bebas, key-value specs produk
- `image`: nullable, string path/URL

### 3.2 Artikel

| Method | Path |
|--------|------|
| `GET`    | `/api/admin/artikel`     |
| `GET`    | `/api/admin/artikel/:id` |
| `POST`   | `/api/admin/artikel`     |
| `PUT`    | `/api/admin/artikel/:id` |
| `DELETE` | `/api/admin/artikel/:id` |

Shape (`types/article.ts`):

```json
{
  "id": "uuid",
  "slug": "tips-merawat-pvc-belt",
  "title": "Tips Merawat PVC Belt Conveyor",
  "excerpt": "Ringkasan singkat artikel...",
  "content": "<p>Konten rich text HTML dari Tiptap...</p>",
  "image": "/images/article/belt-care.jpg",
  "tag": "Perawatan",
  "publishedAt": "2025-07-15",
  "author": "Admin BBS"
}
```

- `content`: HTML dari rich text editor (Tiptap)
- `tag`: string bebas, bukan enum — admin bisa bikin tag baru lewat form artikel
- `publishedAt`: format `YYYY-MM-DD`

### 3.3 Jasa

| Method | Path |
|--------|------|
| `GET`    | `/api/admin/jasa`        |
| `GET`    | `/api/admin/jasa/:id`    |
| `POST`   | `/api/admin/jasa`        |
| `PUT`    | `/api/admin/jasa/:id`    |
| `DELETE` | `/api/admin/jasa/:id`    |

Shape (`types/service.ts`):

```json
{
  "id": "uuid",
  "slug": "incline-cleated-belt",
  "name": "Incline Cleated Belt",
  "shortDescription": "Deskripsi singkat untuk card/listing",
  "fullDescription": "Deskripsi lengkap untuk halaman detail",
  "images": ["/images/service/cleated-1.jpg", "/images/service/cleated-2.jpg"]
}
```

- `images`: array string URL

### 3.4 Galeri

| Method | Path |
|--------|------|
| `GET`    | `/api/admin/galeri`      |
| `GET`    | `/api/admin/galeri/:id`  |
| `POST`   | `/api/admin/galeri`      |
| `PUT`    | `/api/admin/galeri/:id`  |
| `DELETE` | `/api/admin/galeri/:id`  |

Shape (`types/gallery.ts`):

```json
{
  "id": "uuid",
  "image": "/images/gallery/proyek-tambang.jpg",
  "caption": "Instalasi Belt Conveyor di Tambang Batu Bara Kalimantan",
  "location": "Kalimantan Timur"
}
```

- `location`: opsional

### 3.5 Kategori Produk

| Method | Path | Keterangan |
|--------|------|------------|
| `GET` | `/api/admin/kategori` | List semua. Query: `?group=belt-conveyor` |
| `PUT` | `/api/admin/kategori/:slug` | Hanya update `label`, slug readonly |

Shape:

```json
{
  "slug": "pvc-belt",
  "label": "PVC Belt",
  "group": "belt-conveyor",
  "sortOrder": 1
}
```

**Penting**:
- Tidak ada `POST` dan `DELETE` — slug kategori sudah fixed dan tidak bisa ditambah/dihapus oleh admin
- Seed 8 kategori via migration: 4 belt-conveyor (`pvc-belt`, `pu`, `flat-belt`, `rubber-belt`) + 4 lainnya (`timing-belt`, `fastener`, `cleat`, `gravity-roll`)
- Tambah `sortOrder` supaya admin bisa atur urutan tampil

### 3.6 Industri

| Method | Path |
|--------|------|
| `GET`    | `/api/admin/industri`     |
| `GET`    | `/api/admin/industri/:id` |
| `POST`   | `/api/admin/industri`     |
| `PUT`    | `/api/admin/industri/:id` |
| `DELETE` | `/api/admin/industri/:id` |

Shape (`types/industry.ts`):

```json
{
  "id": "uuid",
  "slug": "pertambangan",
  "name": "Pertambangan",
  "description": "Solusi conveyor untuk industri pertambangan...",
  "image": "/images/industry/mining.jpg",
  "productSlugs": ["rubber-belt", "gravity-roll"]
}
```

- `productSlugs`: array slug produk yang relevan untuk industri ini (relasi many-to-many sederhana via slug)

---

## 4. Endpoint Publik (untuk halaman depan)

CRUD untuk semua modul di atas (produk, artikel, jasa, galeri, industri, kategori) juga perlu versi publik **read-only** di luar `/admin`:

| Endpoint Admin | Endpoint Publik |
|----------------|-----------------|
| `GET /api/admin/produk` | `GET /api/produk` |
| `GET /api/admin/produk/:id` | `GET /api/produk/:slug` |
| `GET /api/admin/artikel` | `GET /api/artikel` |
| `GET /api/admin/artikel/:id` | `GET /api/artikel/:slug` |
| ... dst | ... dst |

Endpoint publik **tidak perlu auth**, return shape `{ data, error }` yang sama. Hanya `GET`, tidak ada write.

Prioritas: build endpoint admin dulu (yang mendesak untuk admin panel), endpoint publik bisa menyusul setelahnya.

---

## 5. Yang Tidak Perlu Dibangun (untuk MVP)

### Pengaturan Umum (hero, kontak, about us, footer)

Skip dulu. Semua konten ini hardcode di frontend. Kalau perlu ubah nomor WA atau jam operasional, developer yang edit. Bikin admin UI + backend + migration untuk sesuatu yang berubah 1-2 tahun sekali itu tidak sebanding effortnya.

Kalau suatu saat benar-benar dibutuhkan, cukup satu tabel `settings` single-row dengan 2 endpoint:
- `GET /api/admin/pengaturan`
- `PUT /api/admin/pengaturan`

### Upload File

Tidak perlu endpoint upload terpisah untuk MVP. Image disimpan sebagai string path/URL. Upload manual atau via integrasi nanti.

---

## 6. Database — Ringkasan Tabel

| Tabel | Kolom Kunci |
|-------|------------|
| `products` | id, slug, name, group, kategori, category, description, detail, image, specs (jsonb) |
| `articles` | id, slug, title, excerpt, content (text/html), image, tag, published_at (date), author |
| `services` | id, slug, name, short_description, full_description, images (jsonb/text[]) |
| `galleries` | id, image, caption, location (nullable) |
| `categories` | slug (pk), label, group, sort_order — seed via migration |
| `industries` | id, slug, name, description, image |
| `industry_products` | industry_id, product_slug — tabel pivot |
| `admins` | id, email, password_hash, name |

---

## 7. Urutan Pengerjaan

```
1. Auth (login/logout/me + middleware JWT)
2. Kategori (migration seed — karena dipakai produk)
3. Produk (CRUD admin + endpoint publik)
4. Artikel (CRUD admin + endpoint publik)
5. Jasa (CRUD admin + endpoint publik)
6. Galeri (CRUD admin + endpoint publik)
7. Industri (CRUD admin + endpoint publik)
```
