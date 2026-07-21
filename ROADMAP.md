# ROADMAP — BBS Backend (Clean Architecture)

Pedoman pengerjaan backend dengan **clean architecture**, mengacu pada `BACKEND_REQUIREMENTS.md`.

---

## Arsitektur Folder per Domain

Setiap domain (auth, product, article, service, gallery, category, industry) mengikuti struktur ini:

```
internal/{domain}/
├── entity.go        # pure struct, zero dependency (domain layer)
├── repository.go    # interface + pgx implementation (data layer)
├── usecase.go       # business logic, depends only on repository interface
├── handler.go       # HTTP handler, parse request/response, call usecase
└── route.go         # chi route registration (admin + public)
```

> `shared/` bukan domain — isinya database, http response, middleware, utility lintas modul.

---

## Dependency Rule

```
cmd/api/main.go        (wiring everything)
       │
       ▼
handler  ─────────────► usecase ◄── repository (interface)
  │                        │
  │                        ▼
  ▼                   repository/pgx (implementasi)
 chi
 response helper
 middleware
```

- `entity.go` tidak import apapun selain stdlib
- `repository.go` hanya mendefinisikan interface
- `usecase.go` hanya import entity + repository interface
- `handler.go` import usecase + response helper + chi
- `route.go` import handler
- `main.go` import semua route, database, middleware — wiring manual, tanpa DI framework

---

## Urutan Pengerjaan

### Fase 1 — Auth

| Step | File | Apa yang dikerjakan |
|------|------|---------------------|
| 1.1 | `migrations/000001_create_admins.up.sql` | Tabel `admins` (id, email, password_hash, name, created_at) |
| 1.2 | `internal/auth/entity.go` | Struct `Admin`, `Claims` |
| 1.3 | `internal/auth/repository.go` | Interface `FindByEmail`, implementasi pgx |
| 1.4 | `internal/auth/usecase.go` | `Login(email, password)`, `Me(token)` — bcrypt compare, JWT sign/parse |
| 1.5 | `internal/auth/handler.go` | Handler `POST /login`, `POST /logout`, `GET /me` |
| 1.6 | `internal/auth/route.go` | Register route di chi |
| 1.7 | `internal/shared/middleware/jwt.go` | Middleware validasi JWT cookie |
| 1.8 | `cmd/api/main.go` | Wiring auth module + start server |

**Cek:** Bisa login, logout, `/me` return user, `/admin/produk` return 401 tanpa cookie.

---

### Fase 2 — Kategori

| Step | File | Apa yang dikerjakan |
|------|------|---------------------|
| 2.1 | `migrations/000002_create_categories.up.sql` | Tabel `categories` (slug PK, label, group, sort_order) + seed 8 kategori |
| 2.2 | `internal/category/entity.go` | Struct `Category` |
| 2.3 | `internal/category/repository.go` | Interface `FindAll(filter)`, `UpdateLabel` |
| 2.4 | `internal/category/usecase.go` | `GetAll(group?)`, `UpdateLabel(slug, label)` |
| 2.5 | `internal/category/handler.go` | Handler `GET /admin/kategori`, `PUT /admin/kategori/:slug` |
| 2.6 | `internal/category/route.go` | Admin route + public route `GET /api/kategori` |

---

### Fase 3 — Produk

| Step | File | Apa yang dikerjakan |
|------|------|---------------------|
| 3.1 | `migrations/000003_create_products.up.sql` | Tabel `products` (id, slug, name, group, kategori FK, category, description, detail, image, specs JSONB, timestamps) |
| 3.2 | `internal/product/entity.go` | Struct `Product` |
| 3.3 | `internal/product/repository.go` | Interface `FindAll(filter)`, `FindByID`, `FindBySlug`, `Create`, `Update`, `Delete` |
| 3.4 | `internal/product/usecase.go` | Business logic CRUD + auto-generate slug dari name |
| 3.5 | `internal/product/handler.go` | Handler admin CRUD |
| 3.6 | `internal/product/route.go` | Admin route + public route `GET /api/produk`, `GET /api/produk/:slug` |

---

### Fase 4 — Artikel

| Step | File | Apa yang dikerjakan |
|------|------|---------------------|
| 4.1 | `migrations/000004_create_articles.up.sql` | Tabel `articles` |
| 4.2 | `internal/article/entity.go` | Struct `Article` |
| 4.3 | `internal/article/repository.go` | Interface CRUD |
| 4.4 | `internal/article/usecase.go` | Business logic |
| 4.5 | `internal/article/handler.go` | Handler admin CRUD |
| 4.6 | `internal/article/route.go` | Admin + public route |

---

### Fase 5 — Jasa

| Step | File | Apa yang dikerjakan |
|------|------|---------------------|
| 5.1 | `migrations/000005_create_services.up.sql` | Tabel `services`, `images` sebagai `text[]` |
| 5.2 | `internal/service/entity.go` | Struct `Service` |
| 5.3 | `internal/service/repository.go` | Interface CRUD |
| 5.4 | `internal/service/usecase.go` | Business logic |
| 5.5 | `internal/service/handler.go` | Handler admin CRUD |
| 5.6 | `internal/service/route.go` | Admin + public route |

---

### Fase 6 — Galeri

| Step | File | Apa yang dikerjakan |
|------|------|---------------------|
| 6.1 | `migrations/000006_create_galleries.up.sql` | Tabel `galleries` |
| 6.2 | `internal/gallery/entity.go` | Struct `Gallery` |
| 6.3 | `internal/gallery/repository.go` | Interface CRUD |
| 6.4 | `internal/gallery/usecase.go` | Business logic |
| 6.5 | `internal/gallery/handler.go` | Handler admin CRUD |
| 6.6 | `internal/gallery/route.go` | Admin + public route |

---

### Fase 7 — Industri

| Step | File | Apa yang dikerjakan |
|------|------|---------------------|
| 7.1 | `migrations/000007_create_industries.up.sql` | Tabel `industries` + `industry_products` pivot |
| 7.2 | `internal/industry/entity.go` | Struct `Industry` |
| 7.3 | `internal/industry/repository.go` | Interface CRUD |
| 7.4 | `internal/industry/usecase.go` | Business logic |
| 7.5 | `internal/industry/handler.go` | Handler admin CRUD |
| 7.6 | `internal/industry/route.go` | Admin + public route |

---

## Aturan Wajib per Fase

1. **Migration dulu** — tabel harus ada sebelum kode.
2. **Entity tanpa dependency** — struct murni, tidak import `pgx`, `chi`, atau library eksternal.
3. **Repository return entity** — bukan `pgx.Rows`, bukan `sql.Result`. Layer atas tidak tahu PostgreSQL.
4. **Handler hanya parsing + panggil usecase** — tidak ada query DB atau business logic di handler.
5. **Response shape konsisten** — selalu `{ data, error }` via `http.Success()` / `http.Error()`.
6. **Satu fase selesai dan teruji** sebelum lanjut ke fase berikutnya.

---

## Konvensi

- **Error handling:** domain error didefinisikan di `internal/shared/errors/`, misal `ErrNotFound`, `ErrUnauthorized`, `ErrValidation`. Usecase return error domain, handler map ke HTTP status.
- **Validasi input:** di handler (basic: required, format) + di usecase (business rule: unique slug, dll).
- **Slug generation:** usecase, bukan handler. Dari `name` → lowercase, replace spasi dengan `-`, hapus karakter khusus.
- **Nama kolom DB:** `snake_case`, Go struct: `PascalCase` + json tag `snake_case`.
- **Endpoint publik vs admin:** dibedakan di `route.go`, admin pakai JWT middleware, publik tidak.

---

## Tech Stack

| Kebutuhan | Library |
|-----------|---------|
| Router | `go-chi/chi/v5` |
| Database | `pgx/v5` (pgxpool) |
| JWT | `golang-jwt/jwt/v5` |
| Password hash | `golang.org/x/crypto/bcrypt` |
| UUID | `google/uuid` |
| Migration | `golang-migrate/migrate` |
| Live reload | `air` |

---

## Status Saat Ini

- [x] Config loader (`.env`)
- [x] Database connection (pgxpool)
- [x] HTTP response helper (`{ data, error }`)
- [x] Chi server scaffold
- [x] Direktori domain sudah dibuat
- [x] **Fase 1: Auth** ✅
  - [x] Migration `000001_create_admins`
  - [x] Entity, Repository, Usecase, Handler, Route, Middleware
  - [x] Wiring di `main.go`
  - [x] Build sukses
- [x] **Fase 2: Kategori** ✅
  - [x] Migration `000002_create_categories` + seed 8 kategori
  - [x] Entity, Repository, Usecase, Handler, Route
  - [x] Admin + public routes
- [x] **Fase 3: Produk** ✅
  - [x] Migration `000003_create_products` (FK ke categories)
  - [x] Entity, Repository, Usecase, Handler, Route
  - [x] Admin CRUD + public read-only
  - [x] Slug auto-generate dari name
- [x] **Fase 4: Artikel** ✅
  - [x] Migration `000004_create_articles`
  - [x] Entity, Repository, Usecase, Handler, Route
  - [x] Admin CRUD + public read-only
  - [x] Slug auto-generate dari title
- [x] **Fase 5: Jasa** ✅
  - [x] Migration `000005_create_services` (images as text[])
  - [x] Entity, Repository, Usecase, Handler, Route
  - [x] Admin CRUD + public read-only
- [x] **Fase 6: Galeri** ✅
  - [x] Migration `000006_create_galleries`
  - [x] Entity, Repository, Usecase, Handler, Route
  - [x] Admin CRUD + public read-only
- [x] **Fase 7: Industri** ✅
  - [x] Migration `000007_create_industries` + pivot `industry_products`
  - [x] Entity, Repository, Usecase, Handler, Route
  - [x] Admin CRUD + public read-only
  - [x] Transactional productSlugs sync (delete old, insert new)
