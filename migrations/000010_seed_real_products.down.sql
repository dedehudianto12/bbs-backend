-- Membalik 000010: hapus 24 produk asli dan 3 kategori baru.
--
-- CATATAN: 24 produk placeholder dari 000008 TIDAK dikembalikan — data itu
-- memang sengaja dibuang. Setelah down, tabel products kosong (kecuali produk
-- yang ditambahkan lewat panel admin). Untuk mengembalikan placeholder,
-- jalankan ulang blok INSERT INTO products di 000008_seed_all_data.up.sql.
--
-- PERINGATAN: down juga menghapus tautan foto Cloudinary yang diselamatkan
-- langkah 2 di file up (conveyor-belt-big-square, -pvc-hijau, -pvc-putih).
-- Asetnya sendiri tetap ada di Cloudinary, tapi kolom image-nya hilang dan
-- produk harus di-upload ulang lewat panel admin. Pulihkan dari backup
-- pg_dump kalau ingin utuh.

-- Hanya slug yang diinsert 000010, supaya produk yang ditambahkan lewat panel
-- admin setelah migrasi ini tidak ikut terhapus.
-- industry_products ikut terhapus lewat ON DELETE CASCADE.
DELETE FROM products WHERE slug IN (
    'conveyor-belt-big-square',
    'conveyor-belt-dotted',
    'conveyor-belt-hitam-bandara',
    'conveyor-belt-kapsul',
    'conveyor-belt-pvc-hijau',
    'conveyor-belt-pvc-putih',
    'pvk-belt',
    'roughtop-belt-hitam-hijau',
    'pu-belt-biru',
    'pu-belt-putih',
    'pu-belt-hijau',
    'flat-belt-hijau-kuning',
    'flat-belt-biru',
    'rubber-belt-polos',
    'rubber-belt-sersan',
    'treadmill-belt',
    'sidewall-cleated-conveyor-belt',
    'guide-conveyor-belt-profile',
    'modular-belt',
    'timing-belt',
    'v-belt',
    'fastener-kuku-macan',
    'gravity-roller',
    'open-mesh-belt-dryer'
);

-- Kategori baru hanya dihapus kalau tidak ada produk lain yang memakainya —
-- kalau ada, kategori dibiarkan supaya foreign key products.kategori tetap sah.
DELETE FROM categories c
WHERE c.slug IN ('modular-belt', 'v-belt', 'open-mesh')
  AND NOT EXISTS (SELECT 1 FROM products p WHERE p.kategori = c.slug);
