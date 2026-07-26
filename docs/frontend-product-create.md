# Frontend Guide: Create Product with Image Upload

## Overview

Endpoint `POST /api/products` akan diubah dari `application/json` menjadi `multipart/form-data` agar bisa menerima file gambar langsung dalam satu request. Tidak perlu upload gambar terpisah ke Cloudinary.

## Why

| Sebelum | Sesudah |
|---|---|
| 2 request: upload image → dapat URL → create product | 1 request: create product + upload image |
| Ada risiko orphan di Cloudinary jika user batal submit | Gambar hanya terupload jika produk jadi dibuat |

## Request Spec

```
POST /api/products
Content-Type: multipart/form-data
Authorization: Bearer <token>
```

### Form Fields

| Field | Type | Wajib | Keterangan |
|---|---|---|---|
| `name` | text | ✅ | Nama produk (slug auto-generated) |
| `group` | text | ❌ | `belt-conveyor` / `lainnya` |
| `kategori` | text | ❌ | Slag kategori (FK ke `categories.slug`) |
| `category` | text | ❌ | Nama kategori display |
| `description` | text | ❌ | Deskripsi singkat (plain text) |
| `detail` | text | ❌ | Deskripsi lengkap (HTML) |
| `specs` | text | ❌ | JSON string, contoh: `{"Warna":"Hijau","Ketebalan":"3mm"}` |
| `file` | file | ❌ | Gambar produk, max 10 MB |

### Contoh fetch

```javascript
const formData = new FormData();
formData.append("name", "PVC Belt Hijau Food Grade");
formData.append("group", "belt-conveyor");
formData.append("kategori", "pvc-belt");
formData.append("category", "PVC Belt");
formData.append("description", "PVC belt hijau food grade untuk industri makanan.");
formData.append("detail", "<p>PVC belt hijau food grade...</p>");
formData.append("specs", JSON.stringify({ Warna: "Hijau", Ketebalan: "3mm" }));
formData.append("file", fileInput.files[0]); // File dari input type="file"

const res = await fetch("http://localhost:8080/api/products", {
  method: "POST",
  headers: {
    Authorization: `Bearer ${token}`,
    // Jangan set Content-Type, browser akan set otomatis dengan boundary
  },
  body: formData,
});

const { data, error } = await res.json();
// data: { id, slug, name, ..., image: "https://res.cloudinary.com/.../products/xxx.jpg", ... }
```

## Response

### Sukses (201)
```json
{
  "data": {
    "id": "uuid",
    "slug": "pvc-belt-hijau-food-grade",
    "name": "PVC Belt Hijau Food Grade",
    "group": "belt-conveyor",
    "kategori": "pvc-belt",
    "category": "PVC Belt",
    "description": "...",
    "detail": "<p>...</p>",
    "image": "https://res.cloudinary.com/cloud-name/image/upload/v123/products/xxx.jpg",
    "specs": { "Warna": "Hijau", "Ketebalan": "3mm" },
    "createdAt": "2025-07-26T...",
    "updatedAt": "2025-07-26T..."
  },
  "error": null
}
```

### Gagal (400) — nama kosong
```json
{
  "data": null,
  "error": "nama produk wajib diisi"
}
```

### Gagal (401) — no auth
```json
{
  "data": null,
  "error": "token tidak valid"
}
```

## Update Product (`PUT /api/products/{id}`)

Endpoint update juga akan diubah ke `multipart/form-data`. Semua field **optional** — hanya field yang diisi akan diupdate. Jika `file` dikirim, gambar lama akan diganti.

```
PUT /api/products/{id}
Content-Type: multipart/form-data
Authorization: Bearer <token>
```

Field sama seperti Create, tapi semuanya optional.

## Checklist Frontend

- [ ] Form create: semua field text + file input
- [ ] Form update: semua field optional + file input (gambar bisa diganti)
- [ ] Hapus `Content-Type` header manual, biarkan browser set `multipart/form-data` dengan boundary
- [ ] Preview gambar sebelum submit tetap bisa dilakukan dengan `URL.createObjectURL(file)`
- [ ] Validasi nama tidak boleh kosong
- [ ] `specs` harus dikirim sebagai **JSON string**, bukan object
- [ ] Handle error 400 (validasi) dan 401 (token expired)
