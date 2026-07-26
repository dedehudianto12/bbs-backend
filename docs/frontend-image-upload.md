# Image Upload — Frontend Guide

## Endpoint

```
POST /api/upload/image
```

**Auth:** Bearer token required (include in admin-protected routes group).

## Request

### Headers

```
Authorization: Bearer <token>
Content-Type: multipart/form-data
```

### Body (multipart/form-data)

| Field    | Type   | Required | Description                          |
| -------- | ------ | -------- | ------------------------------------ |
| `file`   | File   | Yes      | Image file (jpg, png, webp, gif)     |
| `folder` | String | No       | Cloudinary folder, defaults to `general` |

**Constraints:**
- Max file size: **10 MB**
- Accepted formats: jpg, jpeg, png, webp, gif

## Response

### Success (200)

```json
{
  "data": {
    "url": "https://res.cloudinary.com/<cloud>/image/upload/v1234567/folder/filename.jpg"
  },
  "error": null
}
```

### Error (4xx / 5xx)

```json
{
  "data": null,
  "error": "file too large" 
}
```

## Example — Fetch / Axios

```js
// JavaScript / React
async function uploadImage(file, folder = "general") {
  const formData = new FormData();
  formData.append("file", file);
  formData.append("folder", folder);

  const res = await fetch("/api/upload/image", {
    method: "POST",
    headers: {
      Authorization: `Bearer ${token}`,
    },
    body: formData,
  });

  const json = await res.json();
  if (json.error) throw new Error(json.error);
  return json.data.url; // → Cloudinary secure URL
}
```

```ts
// TypeScript
interface UploadResponse {
  data: { url: string } | null;
  error: string | null;
}

async function uploadImage(file: File, folder?: string): Promise<string> {
  const formData = new FormData();
  formData.append("file", file);
  if (folder) formData.append("folder", folder);

  const res = await fetch("/api/upload/image", {
    method: "POST",
    headers: { Authorization: `Bearer ${token}` },
    body: formData,
  });

  const json: UploadResponse = await res.json();
  if (json.error) throw new Error(json.error);
  return json.data!.url;
}
```

## Flow

```
Frontend                    Backend                   Cloudinary
   │                           │                          │
   │  POST /api/upload/image   │                          │
   │  (multipart: file)        │                          │
   │──────────────────────────►│                          │
   │                           │  Upload(file, folder)    │
   │                           │─────────────────────────►│
   │                           │                          │
   │                           │  { secure_url }          │
   │                           │◄─────────────────────────│
   │                           │                          │
   │  200 { data: { url } }   │                          │
   │◄──────────────────────────│                          │
   │                           │                          │
   │  Use url as image source  │                          │
   │  or send to create/update │                          │
   │  endpoint                 │                          │
```

## After Upload — Save to Entity

The returned URL should be sent as the `image` field in create/update payloads:

```js
// 1. Upload image
const imageUrl = await uploadImage(file, "articles");

// 2. Create article with the URL
await fetch("/api/admin/articles", {
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
  },
  body: JSON.stringify({
    title: "...",
    slug: "...",
    excerpt: "...",
    content: "<p>...</p>",
    image: imageUrl,  // ← Cloudinary URL
    tag: "Tips",
    publishedAt: "2026-07-24",
    author: "Tim BBS",
  }),
});
```
