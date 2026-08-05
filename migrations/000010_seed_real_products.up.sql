-- ============================================================
-- Ganti 24 produk placeholder dari 000008 dengan katalog asli.
--
-- Sumber: "update list produk.pdf" (24 produk).
--
-- Tiga produk tidak masuk 8 kategori yang ada (Modular Belt, V-Belt,
-- Open Mesh) sehingga kategori barunya dibuat di sini, group 'lainnya'.
--
-- `detail` diisi HTML — kolom ini dirender lewat v-html di
-- pages/produk/[slug].vue dan disanitasi oleh utils/richtext.ts, jadi
-- hanya tag yang ada di ALLOWED_TAGS (p, h3, ul, li, strong) yang dipakai.
-- Judul "Spesifikasi Teknis" sudah dirender halaman dari kolom `specs`,
-- jadi tidak diulang di dalam `detail`.
--
-- `image` sengaja NULL di blok VALUES — foto diunggah lewat panel admin
-- (Cloudinary). Kolom itu TIDAK ikut di-overwrite pada ON CONFLICT, jadi foto
-- yang sudah diunggah ke prod tetap aman saat migrasi ini dijalankan.
-- ============================================================

-- ── 1. KATEGORI BARU ──

INSERT INTO categories (slug, label, "group", sort_order) VALUES
    ('modular-belt', 'Modular Belt', 'lainnya',  9),
    ('v-belt',       'V-Belt',       'lainnya', 10),
    ('open-mesh',    'Open Mesh',    'lainnya', 11)
ON CONFLICT (slug) DO NOTHING;

-- ── 2. SELAMATKAN FOTO YANG SUDAH DIUNGGAH ──
-- Prod sudah punya foto Cloudinary di tiga baris. Dua di antaranya memakai slug
-- placeholder yang di katalog asli namanya berubah; slug-nya di-rename lebih
-- dulu supaya baris (beserta kolom image) bertahan dan cocok dengan ON CONFLICT
-- di bawah. FK industry_products.product_slug sudah ON UPDATE CASCADE sejak
-- 000009, jadi rename ini aman.
--
-- Produk ketiga, 'conveyor-belt-big-square', dibuat manual lewat panel admin
-- dan slug-nya sudah sama dengan katalog asli — tidak perlu di-rename.

UPDATE products SET slug = 'conveyor-belt-pvc-hijau' WHERE slug = 'pvc-belt-hijau-food-grade';
UPDATE products SET slug = 'conveyor-belt-pvc-putih' WHERE slug = 'pvc-belt-putih-packaging';

-- ── 3. HAPUS SISA DATA PLACEHOLDER ──
-- Hanya baris yang bukan bagian dari katalog asli. Baris yang slug-nya cocok
-- dibiarkan hidup supaya kolom image-nya tidak hilang — isinya di-refresh oleh
-- ON CONFLICT DO UPDATE di bawah.

DELETE FROM industry_products;

DELETE FROM products WHERE slug NOT IN (
    'conveyor-belt-big-square', 'conveyor-belt-dotted', 'conveyor-belt-hitam-bandara',
    'conveyor-belt-kapsul', 'conveyor-belt-pvc-hijau', 'conveyor-belt-pvc-putih',
    'pvk-belt', 'roughtop-belt-hitam-hijau', 'pu-belt-biru', 'pu-belt-putih',
    'pu-belt-hijau', 'flat-belt-hijau-kuning', 'flat-belt-biru', 'rubber-belt-polos',
    'rubber-belt-sersan', 'treadmill-belt', 'sidewall-cleated-conveyor-belt',
    'guide-conveyor-belt-profile', 'modular-belt', 'timing-belt', 'v-belt',
    'fastener-kuku-macan', 'gravity-roller', 'open-mesh-belt-dryer'
);

-- ── 4. PRODUK ASLI (24 items) ──

INSERT INTO products (slug, name, "group", kategori, category, description, detail, image, specs) VALUES

-- ── PVC Belt (8) ──

('conveyor-belt-big-square', 'Conveyor Belt Big Square', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC conveyor belt permukaan motif kotak untuk mesin sending/sanding industri woodworking. Daya cengkeram kuat menjaga material kayu tidak bergeser saat proses berlangsung.',
 '<p>PVC Conveyor Belt dengan permukaan motif kotak ini dirancang khusus untuk kebutuhan industri woodworking, terutama pada mesin sending atau sanding yang memerlukan daya cengkeram kuat dan stabilitas tinggi selama proses kerja. Permukaan kotaknya berfungsi menjaga material kayu agar tidak bergeser saat proses pengamplasan atau pengiriman berlangsung, memastikan hasil kerja yang presisi dan efisien.</p><h3>Manfaat dan Fungsi Produk</h3><ul><li>Menjaga stabilitas bahan kayu selama proses pengamplasan di mesin sending/sanding.</li><li>Mengurangi risiko slip atau pergeseran material pada conveyor.</li><li>Mendukung kinerja mesin agar lebih efisien dan aman.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Permukaan kotak stabil, ideal untuk material kayu dan permukaan halus.</li><li>Tahan aus dan memiliki umur pakai panjang meski di bawah beban kerja berat.</li><li>Dapat disesuaikan dengan berbagai jenis mesin woodworking dan sistem conveyor.</li></ul><h3>Rekomendasi Penggunaan</h3><ul><li>Pabrik woodworking dan furniture</li><li>Mesin sending/sanding kayu otomatis</li><li>Lini produksi kayu dan panel industri</li></ul>',
 NULL,
 '{"Bahan": "PVC industri berkualitas tinggi", "Permukaan": "Motif kotak (square pattern)", "Ketebalan": "±9 mm", "Lebar": "Hingga 1600 mm (custom)", "Warna": "Hitam"}'),

('conveyor-belt-dotted', 'Conveyor Belt Dotted', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC belt dengan pola tonjolan bulat (round stud) untuk daya cengkeram tinggi. Mencegah produk tergelincir selama proses produksi maupun transportasi material.',
 '<p>PVC Dotted Belt (Round Stud Pattern) merupakan jenis conveyor belt yang dirancang khusus untuk aplikasi industri yang membutuhkan daya cengkeram tinggi, ketahanan aus, dan performa stabil. Permukaannya memiliki pola tonjolan bulat (round stud) yang membantu meningkatkan gesekan dan mencegah produk tergelincir selama proses produksi atau transportasi material.</p><h3>Manfaat &amp; Fungsi Produk</h3><ul><li>Mempermudah dan menstabilkan perpindahan barang di jalur produksi maupun pergudangan.</li><li>Mengurangi risiko slip dan kerusakan produk saat proses transportasi internal.</li><li>Mendukung sistem logistik dan produksi yang cepat, aman, dan efisien.</li></ul><h3>Keunggulan Produk</h3><ul><li>Daya cengkeram tinggi pada permukaan barang atau bahan kerja.</li><li>Tahan lama dan minim perawatan.</li><li>Stabil digunakan pada conveyor panjang dengan beban konstan.</li><li>Cocok untuk berbagai sistem transportasi industri.</li></ul><h3>Aplikasi Industri</h3><ul><li>Industri logistik dan pergudangan</li><li>Lini produksi manufaktur</li><li>Mesin sorting dan packaging</li><li>Sistem conveyor pengangkutan material ringan hingga sedang</li></ul>',
 NULL,
 '{"Bahan": "PVC (Polyvinyl Chloride)", "Permukaan": "Pola tonjolan bulat (round stud)", "Lebar": "Custom sesuai kebutuhan industri", "Ketebalan": "Menyesuaikan aplikasi dan kapasitas beban"}'),

('conveyor-belt-hitam-bandara', 'Conveyor Belt Hitam Bandara', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'Conveyor belt hitam untuk pemindahan barang bervolume tinggi. Material tahan lama dengan performa stabil di kondisi pemakaian intens.',
 '<p>Belt Hitam Bandara adalah conveyor belt khusus yang dirancang untuk pemindahan barang secara efisien. Terbuat dari material berkualitas tinggi yang tahan lama, belt ini memberikan performa stabil dan awet dalam berbagai kondisi operasional.</p><h3>Manfaat / Fungsi Produk</h3><ul><li>Memastikan pergerakan barang lebih lancar dan konsisten.</li><li>Tahan terhadap gesekan dan beban berat.</li><li>Meminimalkan kebutuhan perawatan rutin.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Daya tahan lebih tinggi terhadap beban berat.</li><li>Permukaan yang stabil untuk pergerakan barang.</li><li>Mudah dipasang dan diganti jika diperlukan.</li></ul><h3>Rekomendasi Industri atau Penggunaan</h3><ul><li>Industri manufaktur dan logistik</li><li>Sistem conveyor di gudang dan pusat distribusi</li><li>Unit conveyor industri umum</li></ul>',
 NULL,
 '{"Material": "Rubber atau PVC berkualitas tinggi", "Warna": "Hitam", "Lebar": "Sesuai permintaan industri", "Ketahanan": "Stabil di kondisi pemakaian intens"}'),

('conveyor-belt-kapsul', 'Conveyor Belt Kapsul', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC belt bertekstur motif kapsul anti-slip untuk mesin sanding dan lini produksi ringan hingga menengah. Tahan aus, fleksibel, dan mudah dibersihkan.',
 '<p>PVC Conveyor Belt Motif Kapsul adalah jenis conveyor belt dengan permukaan bertekstur kapsul yang dirancang untuk memberikan daya cengkeram optimal dan stabilitas tinggi pada proses produksi. Terbuat dari material PVC berkualitas tinggi, belt ini memiliki karakteristik tahan aus, fleksibel, dan mudah dibersihkan, menjadikannya ideal untuk digunakan pada mesin sanding, sistem transportasi internal, dan lini produksi industri ringan hingga menengah.</p><h3>Manfaat &amp; Fungsi</h3><ul><li>Meningkatkan efisiensi perpindahan barang pada jalur produksi.</li><li>Mengurangi risiko slip atau kerusakan produk.</li><li>Mendukung operasional mesin dengan kecepatan dan kestabilan optimal.</li></ul><h3>Keunggulan</h3><ul><li>Tahan lama dan tidak mudah terkelupas.</li><li>Stabil di kecepatan tinggi dan jarak conveyor panjang.</li><li>Perawatan mudah dan ekonomis.</li><li>Bisa custom ukuran dan warna.</li></ul><h3>Aplikasi Industri</h3><ul><li>Mesin sanding (woodworking)</li><li>Pergudangan &amp; logistik</li><li>Industri pengemasan (packaging line)</li><li>Sistem transportasi internal pabrik</li></ul>',
 NULL,
 '{"Bahan": "PVC (Polyvinyl Chloride)", "Permukaan": "Motif kapsul (anti-slip)", "Lebar & Panjang": "Custom sesuai kebutuhan industri", "Ketebalan": "Berbagai opsi sesuai aplikasi"}'),

('conveyor-belt-pvc-hijau', 'Conveyor Belt PVC Hijau', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'Belt conveyor PVC hijau untuk aplikasi umum industri. Daya tahan tinggi, fleksibilitas baik, dan cocok untuk beban ringan hingga menengah.',
 '<p>PVC Hijau Belt Conveyor adalah jenis belt conveyor yang dirancang dengan material PVC berkualitas. Produk ini memiliki daya tahan tinggi, fleksibilitas baik, serta permukaan yang mendukung proses produksi di berbagai sektor industri. Warna hijau pada belt memberikan identitas visual yang umum digunakan dalam sistem transportasi material ringan hingga menengah.</p><h3>Manfaat / Fungsi Produk</h3><ul><li>Mempermudah proses pemindahan barang secara cepat dan efisien.</li><li>Mengurangi beban tenaga kerja dengan sistem transportasi otomatis.</li><li>Mendukung kelancaran proses produksi dan distribusi.</li><li>Memberikan hasil produksi yang lebih konsisten.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Permukaan belt rata dan stabil untuk menjaga kualitas produk yang dibawa.</li><li>Mudah dipasang dan dirawat.</li><li>Tahan terhadap gesekan sehingga memiliki umur pakai lebih panjang.</li><li>Biaya perawatan relatif rendah.</li></ul><h3>Rekomendasi Industri / Penggunaan</h3><ul><li>Industri makanan dan minuman</li><li>Industri farmasi</li><li>Industri tekstil</li><li>Sektor logistik dan distribusi</li><li>Industri elektronik</li></ul>',
 NULL,
 '{"Material Utama": "PVC (Polyvinyl Chloride)", "Warna": "Hijau", "Kekuatan Tarik": "Baik, dengan ketahanan aus tinggi", "Kapasitas Beban": "Ringan hingga menengah"}'),

('conveyor-belt-pvc-putih', 'Conveyor Belt PVC Putih', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC belt putih food grade, non-toxic dan aman kontak langsung dengan makanan. Tersedia 1–8 mm, 1–4 ply, custom lebar hingga 2000 mm.',
 '<p>PVC Conveyor Belt Putih merupakan jenis belt konveyor yang dirancang khusus untuk kebutuhan industri yang mengutamakan kebersihan, kekuatan, dan kelancaran proses produksi.</p><p>Terbuat dari material PVC Food Grade berkualitas tinggi, belt ini aman untuk kontak langsung dengan produk makanan, serta banyak digunakan pada industri makanan, minuman, roti, farmasi, dan kemasan.</p><p>Tersedia dalam berbagai ketebalan mulai dari 1 mm hingga 8 mm dan konstruksi 1 ply hingga 4 ply, sehingga dapat disesuaikan dengan kebutuhan sistem konveyor di pabrik Anda. Belt ini juga dapat dipesan secara custom hingga lebar 2000 mm, cocok untuk berbagai jenis mesin dan aplikasi produksi.</p><h3>Fitur &amp; Keunggulan</h3><ul><li>Material Food Grade, non-toxic, dan aman untuk kontak langsung dengan bahan pangan.</li><li>Permukaan putih halus, mudah dibersihkan, dan higienis.</li><li>Tahan terhadap minyak, air, dan bahan kimia ringan.</li><li>Dapat dibuat custom ukuran (lebar hingga 2000 mm).</li><li>Proses cepat — lead time 1–3 hari kerja.</li></ul>',
 NULL,
 '{"Material": "PVC Food Grade", "Warna": "Putih", "Ketebalan": "1–8 mm", "Konstruksi": "1 ply s/d 4 ply", "Lebar": "Custom hingga 2000 mm", "Lead Time": "1–3 hari kerja"}'),

('pvk-belt', 'PVK Belt', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'Conveyor belt PVK untuk aplikasi industri yang membutuhkan ketahanan, fleksibilitas, dan performa stabil pada jalur conveyor panjang.',
 '<p>PVK Belt merupakan jenis conveyor belt yang dirancang untuk aplikasi industri yang membutuhkan ketahanan, fleksibilitas, dan performa stabil. Belt ini banyak digunakan dalam sistem conveyor untuk memindahkan barang secara efisien di berbagai jenis industri.</p><h3>Manfaat / Fungsi Produk</h3><ul><li>Mempermudah perpindahan barang di jalur produksi atau pergudangan.</li><li>Mengurangi risiko kerusakan produk selama proses transportasi internal.</li><li>Mendukung operasi logistik yang cepat dan efisien.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Tahan lama dan minim perawatan.</li><li>Stabilitas tinggi saat digunakan dalam jalur conveyor panjang.</li><li>Fleksibel untuk berbagai jenis aplikasi transportasi barang.</li></ul><h3>Rekomendasi Industri atau Penggunaan</h3><ul><li>Industri logistik dan pergudangan</li><li>Distribusi barang dan paket</li><li>Sistem conveyor di pabrik manufaktur</li></ul>',
 NULL,
 '{"Bahan": "PVC / Polyvinyl Chloride", "Lebar": "Sesuai kebutuhan industri", "Ketebalan": "Standar industri, menyesuaikan aplikasi", "Tipe Permukaan": "Polos atau dengan profil khusus"}'),

('roughtop-belt-hitam-hijau', 'Roughtop Hitam & Hijau', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC belt permukaan bergerigi (roughtop) anti-slip untuk conveyor incline/menanjak. Tersedia hitam dan hijau, custom lebar hingga 2000 mm.',
 '<p>PVC Roughtop Conveyor Belt adalah jenis belt konveyor dengan permukaan bergerigi (roughtop) yang dirancang untuk memberikan daya cengkeram tinggi pada material saat proses berjalan.</p><p>Cocok digunakan pada industri packaging, logistik, carton box, percetakan, dan konveyor incline/menanjak, di mana kestabilan dan daya tarik permukaan sangat dibutuhkan.</p><h3>Fitur &amp; Keunggulan</h3><ul><li>Permukaan roughtop anti-slip, ideal untuk area menanjak atau produk kemasan ringan.</li><li>Material kuat, fleksibel, dan tahan aus.</li><li>Tersedia warna hitam dan hijau.</li><li>Dapat custom lebar hingga 2000 mm.</li><li>Proses cepat dengan lead time 1–3 hari kerja.</li></ul><h3>Rekomendasi Penggunaan</h3><ul><li>Industri packaging dan carton box</li><li>Industri logistik dan percetakan</li><li>Conveyor incline/menanjak</li></ul>',
 NULL,
 '{"Material": "PVC", "Permukaan": "Roughtop (bergerigi) anti-slip", "Warna": "Hitam & Hijau", "Ketebalan": "5–7 mm", "Konstruksi": "3 ply", "Lebar": "Custom hingga 2000 mm", "Lead Time": "1–3 hari kerja"}'),

-- ── PU (3) ──

('pu-belt-biru', 'PU Biru', 'belt-conveyor', 'pu', 'PU',
 'PU conveyor belt biru untuk quality control pabrik makanan dan farmasi. Warna kontras mempermudah deteksi kontaminasi visual, tahan minyak dan lemak.',
 '<p><strong>Deteksi Kontaminasi Visual | Standar Higienis Tinggi | Tahan Minyak</strong></p><p>PU Conveyor Belt Biru dirancang khusus untuk meningkatkan kontrol kualitas (quality control) pada pabrik olahan makanan dan farmasi. Warna biru kontras memudahkan operator maupun sensor visual mendeteksi secara cepat jika terdapat serpihan belt atau bahan asing yang jatuh ke area produksi.</p><h3>Keunggulan Utama</h3><ul><li><strong>Kontras Tinggi:</strong> memudahkan deteksi dini kontaminasi fisik pada produk berwarna terang/putih.</li><li><strong>Sesuai Standar HACCP &amp; FDA:</strong> aman untuk kontak makanan dan mendukung manajemen keamanan pangan.</li><li><strong>Fleksibel &amp; Awet:</strong> tahan terhadap minyak, lemak hewan/nabati, dan gesekan konstan.</li></ul><h3>Aplikasi Utama</h3><ul><li>Pengolahan daging, unggas, dan seafood</li><li>Pengolahan adonan roti</li><li>Pengemasan produk obat/kesehatan</li></ul>',
 NULL,
 '{"Material": "PU (Polyurethane)", "Warna": "Biru", "Standar": "HACCP & FDA food contact", "Ketahanan": "Minyak, lemak hewan/nabati, gesekan konstan"}'),

('pu-belt-putih', 'PU Putih', 'belt-conveyor', 'pu', 'PU',
 'PU conveyor belt putih food & pharma grade. Bebas toksik, tidak berbau, permukaan non-porous yang mudah disterilkan.',
 '<p><strong>Higienis Standar FDA | Anti-Bakteri | Mudah Dibersihkan</strong></p><p>PU Conveyor Belt Putih adalah standar utama sabuk berjalan untuk industri makanan dan farmasi. Terbuat dari material Polyurethane (PU) berkualitas tinggi, belt ini bebas toksik, tidak berbau, dan memiliki permukaan halus yang mencegah akumulasi sisa bahan baku atau bakteri.</p><h3>Keunggulan Utama</h3><ul><li><strong>Standar Food &amp; Pharma Grade:</strong> aman untuk kontak langsung dengan produk makanan dan obat-obatan.</li><li><strong>Permukaan Non-Porous:</strong> mudah disterilkan dan tahan terhadap pembersih kimia/alkohol.</li><li><strong>Tahan Minyak &amp; Lemak:</strong> tidak mudah mengembang atau rusak saat terkena kandungan minyak/lemak.</li></ul><h3>Aplikasi Utama</h3><ul><li>Lini pencetakan tablet farmasi</li><li>Industri roti/bakery</li><li>Pengolahan cokelat dan susu</li><li>Pengemasan makanan kering</li></ul>',
 NULL,
 '{"Material": "PU (Polyurethane)", "Warna": "Putih", "Standar": "Food & Pharma Grade (FDA)", "Permukaan": "Non-porous, mudah disterilkan", "Ketahanan": "Minyak, lemak, pembersih kimia/alkohol"}'),

('pu-belt-hijau', 'PU Hijau', 'belt-conveyor', 'pu', 'PU',
 'PU conveyor belt hijau untuk aplikasi industri umum. Ketahanan gesek ekstra, bebas mulur pada kecepatan tinggi atau beban sedang.',
 '<p><strong>Tahan Aus &amp; Gesekan | Struktur Fleksibel | Bebas Mulur</strong></p><p>PU Conveyor Belt Hijau merupakan pilihan favorit untuk aplikasi industri umum yang membutuhkan daya tahan ekstra terhadap gesekan dan mekanis. Warna hijau standar industri memberikan tampilan bersih dan profesional pada lini transmisi atau conveyor penyortiran.</p><h3>Keunggulan Utama</h3><ul><li><strong>Tahan Aus Ekstra:</strong> struktur material PU hijau memiliki ketahanan gesek (abrasion resistance) yang sangat baik.</li><li><strong>Stabil &amp; Presisi:</strong> bebas mulur saat digunakan pada kecepatan tinggi atau beban sedang.</li><li><strong>Toleransi Suhu Baik:</strong> stabil digunakan pada lingkungan kerja dengan fluktuasi suhu sedang.</li></ul><h3>Aplikasi Utama</h3><ul><li>Industri elektronik</li><li>Penyortiran logistik/paket</li><li>Pencetakan kertas/karton</li><li>Industri tekstil</li><li>Perakitan umum</li></ul>',
 NULL,
 '{"Material": "PU (Polyurethane)", "Warna": "Hijau", "Ketahanan Gesek": "Abrasion resistance sangat baik", "Elongasi": "Bebas mulur pada kecepatan tinggi", "Toleransi Suhu": "Stabil pada fluktuasi suhu sedang"}'),

-- ── Flat Belt (2) ──

('flat-belt-hijau-kuning', 'Flat Belt Hijau Kuning', 'belt-conveyor', 'flat-belt', 'Flat Belt',
 'Sabuk datar untuk mesin tekstil, percetakan, dan mesin ringan. Permukaan halus dengan pergerakan material yang stabil pada kecepatan rendah hingga menengah.',
 '<p>Flat Belt Hijau Kuning adalah jenis sabuk datar yang dirancang untuk kebutuhan industri dengan pergerakan material yang stabil dan efisien. Produk ini banyak digunakan pada industri tekstil, percetakan, dan mesin ringan, membantu proses produksi berjalan lebih lancar dan presisi.</p><h3>Manfaat dan Fungsi Produk</h3><ul><li>Memastikan pergerakan material tetap stabil pada mesin.</li><li>Mengurangi risiko slip atau selip pada belt saat operasi.</li><li>Cocok untuk aplikasi dengan kecepatan rendah hingga menengah.</li><li>Mendukung efisiensi dan produktivitas proses industri.</li></ul><h3>Keunggulan Produk</h3><ul><li>Material berkualitas tinggi yang tahan aus.</li><li>Permukaan halus sehingga tidak merusak material yang dipindahkan.</li><li>Tersedia dalam ukuran yang dapat disesuaikan dengan kebutuhan mesin.</li></ul><h3>Rekomendasi Penggunaan</h3><ul><li>Mesin tekstil, seperti pemintal dan pemroses kain</li><li>Mesin percetakan dan pemindah bahan ringan</li><li>Industri ringan yang membutuhkan conveyor belt datar</li></ul>',
 NULL,
 '{"Warna": "Hijau / Kuning", "Permukaan": "Halus", "Kecepatan Kerja": "Rendah hingga menengah", "Ukuran": "Custom sesuai kebutuhan mesin"}'),

('flat-belt-biru', 'Flat Belt Biru', 'belt-conveyor', 'flat-belt', 'Flat Belt',
 'Flat belt untuk industri dengan standar kebersihan tinggi. Permukaan halus meminimalkan risiko kontaminasi produk, tahan gesekan jangka panjang.',
 '<p>Flat Belt Biru adalah jenis conveyor belt yang dirancang untuk aplikasi industri dengan kebutuhan higienis dan tahan lama. Dengan permukaan yang halus dan material berkualitas, belt ini ideal untuk transportasi produk dalam lini produksi, terutama pada industri yang menuntut standar kebersihan tinggi. Fleksibilitas dan kekuatan materialnya membuat Flat Belt Biru dapat digunakan dalam berbagai sistem conveyor, menjaga kelancaran aliran produksi tanpa mengurangi kualitas produk yang diangkut.</p><h3>Manfaat dan Fungsi</h3><ul><li>Memastikan transportasi produk yang stabil dan efisien di lini produksi.</li><li>Cocok untuk industri makanan, farmasi, dan manufaktur dengan standar kebersihan tinggi.</li><li>Permukaan halus meminimalkan risiko kerusakan atau kontaminasi produk.</li><li>Tahan terhadap gesekan dan pemakaian jangka panjang.</li></ul><h3>Keunggulan Produk</h3><ul><li>Material berkualitas tinggi yang tahan lama.</li><li>Mudah dipasang dan dirawat.</li><li>Dirancang untuk performa optimal pada berbagai jenis conveyor.</li></ul><h3>Rekomendasi Penggunaan</h3><ul><li>Sistem conveyor industri makanan dan minuman</li><li>Proses produksi farmasi dan produk higienis</li><li>Transportasi barang ringan hingga sedang</li></ul>',
 NULL,
 '{"Warna": "Biru", "Permukaan": "Halus, higienis", "Kapasitas Beban": "Ringan hingga sedang", "Ukuran": "Custom sesuai sistem conveyor"}'),

-- ── Rubber Belt (3) ──

('rubber-belt-polos', 'Rubber Belt Polos', 'belt-conveyor', 'rubber-belt', 'Rubber Belt',
 'Rubber conveyor belt polos untuk pengangkutan material berat. Karet berkualitas tinggi yang tahan abrasi, tekanan, dan kondisi kerja berat.',
 '<p>Rubber Conveyor Belt Polos merupakan jenis belt yang dirancang khusus untuk pengangkutan material berat. Belt ini terbuat dari karet berkualitas tinggi sehingga memiliki daya tahan lama terhadap gesekan, tekanan, dan kondisi kerja berat. Produk ini sangat cocok untuk sistem conveyor di berbagai industri yang membutuhkan ketahanan maksimal dan performa stabil.</p><h3>Manfaat / Fungsi Produk</h3><ul><li>Memastikan pengangkutan material berat berjalan lancar dan efisien.</li><li>Tahan terhadap abrasi dan gesekan dalam jangka panjang.</li><li>Menjamin stabilitas aliran material di conveyor.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Daya tahan lebih tinggi terhadap abrasi dan kerusakan.</li><li>Meminimalisir gangguan operasional akibat belt cepat aus.</li><li>Kualitas konsisten untuk penggunaan jangka panjang.</li></ul><h3>Rekomendasi Industri / Penggunaan</h3><ul><li>Pertambangan dan quarry</li><li>Industri pengolahan material berat</li><li>Sistem conveyor pabrik yang membutuhkan daya tahan tinggi</li></ul>',
 NULL,
 '{"Bahan": "Karet natural dan sintetis berkualitas tinggi", "Permukaan": "Polos", "Ketebalan": "Sesuai kebutuhan aplikasi industri", "Lebar": "Disesuaikan dengan jenis conveyor", "Ketahanan": "Tekanan tinggi dan kondisi kerja berat"}'),

('rubber-belt-sersan', 'Rubber Belt Sersan', 'belt-conveyor', 'rubber-belt', 'Rubber Belt',
 'Rubber belt bermotif sersan untuk material menanjak dan kondisi ekstrim. Daya tahan tinggi untuk sektor pertambangan, konstruksi, dan manufaktur berat.',
 '<p>Rubber Conveyor Belt Sersan dirancang khusus untuk kebutuhan industri yang menuntut daya tahan tinggi terhadap beban berat dan kondisi ekstrim. Produk ini mampu menangani material menanjak dan beragam jenis muatan tanpa mengurangi kinerja conveyor, sehingga ideal untuk operasional di sektor pertambangan, konstruksi, dan manufaktur material berat.</p><h3>Manfaat dan Fungsi Produk</h3><ul><li>Tahan terhadap beban berat dan material kasar.</li><li>Cocok untuk aplikasi menanjak dan industri berat.</li><li>Memastikan transportasi material yang stabil dan efisien.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Durabilitas tinggi untuk penggunaan jangka panjang.</li><li>Tahan terhadap aus dan abrasi material.</li><li>Mendukung kinerja conveyor yang optimal di medan berat.</li></ul><h3>Rekomendasi Industri atau Penggunaan</h3><ul><li>Industri pertambangan</li><li>Industri konstruksi</li><li>Industri pengolahan material menanjak</li></ul>',
 NULL,
 '{"Bahan": "Karet industri berkualitas tinggi", "Permukaan": "Motif sersan (chevron)", "Ketebalan": "Sesuai kebutuhan aplikasi", "Lebar": "Variasi sesuai mesin conveyor", "Aplikasi": "Conveyor menanjak / beban berat"}'),

('treadmill-belt', 'Conveyor Belt Treadmill', 'belt-conveyor', 'rubber-belt', 'Rubber Belt',
 'Treadmill belt untuk industri fitness dan peralatan olahraga. Permukaan anti-slip yang stabil dan tahan lama untuk penggunaan intensif.',
 '<p>Treadmill Belt dirancang khusus untuk industri fitness dan peralatan olahraga, memberikan performa maksimal dan daya tahan tinggi untuk penggunaan intensif pada mesin treadmill. Produk ini membantu memastikan pengalaman latihan yang lancar dan aman, serta mendukung efisiensi operasional di gym, pusat kebugaran, atau fasilitas olahraga profesional.</p><h3>Manfaat dan Fungsi Produk</h3><ul><li>Menjamin kelancaran gerakan treadmill dengan permukaan yang stabil.</li><li>Tahan lama terhadap gesekan dan tekanan dari penggunaan berulang.</li><li>Mendukung keamanan pengguna dengan grip yang baik.</li><li>Meminimalisir perawatan rutin berkat bahan berkualitas tinggi.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Daya tahan tinggi untuk penggunaan intensif.</li><li>Permukaan anti-slip untuk keamanan pengguna.</li><li>Stabilitas gerakan lebih baik dibanding treadmill belt umum.</li></ul><h3>Rekomendasi Industri atau Penggunaan</h3><ul><li>Gym dan pusat kebugaran</li><li>Studio olahraga dan rehabilitasi fisik</li><li>Treadmill rumah atau komersial yang membutuhkan kualitas profesional</li></ul>',
 NULL,
 '{"Bahan": "Rubber atau PVC kualitas industri fitness", "Permukaan": "Anti-slip", "Ketebalan": "Standar industri treadmill", "Lebar": "Menyesuaikan ukuran treadmill komersial"}'),

-- ── Cleat (2) ──

('sidewall-cleated-conveyor-belt', 'Sidewall Cleated Conveyor Belt', 'lainnya', 'cleat', 'Cleat',
 'Conveyor belt dengan cleat untuk pengangkutan material menanjak. Menjaga material tetap stabil dan meminimalkan risiko tumpah atau tergelincir.',
 '<p>Conveyor Belt + Cleat dirancang untuk memudahkan pengangkutan material yang menanjak. Dengan tambahan cleat pada permukaan belt, produk ini mampu menjaga material tetap stabil saat bergerak naik, sehingga meminimalkan risiko tumpah atau tergelincir. Belt ini banyak digunakan pada industri makanan, agrikultur, dan packaging yang membutuhkan efisiensi dan keamanan material handling.</p><h3>Manfaat dan Fungsi Produk</h3><ul><li>Memastikan material tetap berada di posisi saat bergerak menanjak.</li><li>Meningkatkan efisiensi proses produksi dengan pengangkutan yang lebih stabil.</li><li>Mengurangi risiko tumpah atau kerusakan material selama transportasi.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Desain cleat yang kokoh untuk stabilitas material optimal.</li><li>Tahan lama dan mudah dibersihkan untuk industri makanan dan packaging.</li><li>Fleksibel untuk diaplikasikan pada berbagai tipe conveyor menanjak.</li></ul><h3>Rekomendasi Industri atau Penggunaan</h3><ul><li>Industri makanan dan minuman</li><li>Industri agrikultur</li><li>Industri packaging dan logistik</li></ul>',
 NULL,
 '{"Material": "PVC atau pilihan sesuai kebutuhan industri", "Konfigurasi": "Sidewall + cleat", "Ukuran": "Lebar dan panjang custom sesuai mesin conveyor", "Ketahanan": "Berbagai jenis material dan suhu operasi"}'),

('guide-conveyor-belt-profile', 'Guide Conveyor Belt (Profile)', 'lainnya', 'cleat', 'Cleat',
 'PVC belt dengan profil tambahan untuk mencegah produk tergelincir atau berpindah posisi di jalur konveyor. Umum dipakai di logistik, pengemasan, dan makanan.',
 '<p>PVC Conveyor Belt + Profile adalah sabuk konveyor yang dirancang khusus dengan profil tambahan pada permukaan belt untuk mencegah produk tergelincir atau berpindah posisi saat melewati jalur konveyor. Produk ini banyak digunakan pada industri logistik, pengemasan, dan makanan karena mampu mengoptimalkan aliran barang dengan lebih aman dan efisien. Material PVC memberikan daya tahan tinggi terhadap keausan dan fleksibilitas yang memudahkan pemasangan serta pemeliharaan.</p><h3>Manfaat dan Fungsi Produk</h3><ul><li>Memastikan produk tetap berada di jalur konveyor tanpa tergelincir.</li><li>Mendukung proses produksi yang lebih cepat dan efisien.</li><li>Meningkatkan keamanan dan konsistensi aliran barang di lini produksi.</li><li>Tahan lama terhadap gesekan dan beban produk ringan hingga sedang.</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Profil permukaan khusus mencegah tergelincirnya produk.</li><li>Fleksibel dan mudah dipasang.</li><li>Tahan lama dan minim perawatan.</li><li>Aman digunakan di lingkungan industri makanan.</li></ul><h3>Rekomendasi Industri atau Penggunaan</h3><ul><li>Industri logistik dan pergudangan</li><li>Industri pengemasan dan makanan</li><li>Jalur konveyor dengan produk ringan hingga sedang yang memerlukan kontrol posisi</li></ul>',
 NULL,
 '{"Material": "PVC berkualitas tinggi", "Profil": "Desain bervariasi sesuai aplikasi", "Lebar & Ketebalan": "Custom sesuai kebutuhan industri", "Kompatibilitas": "Berbagai jenis mesin konveyor"}'),

-- ── Modular Belt (1) ──

('modular-belt', 'Modular Belt', 'lainnya', 'modular-belt', 'Modular Belt',
 'Conveyor belt berbasis segmen plastik interlock, bebas karat dan food grade. Ideal untuk jalur lurus, berbelok, maupun incline/decline.',
 '<p>Modular Plastic Belt adalah solusi rantai berjalan (conveyor belt) modern berbasis segmen plastik interlock yang dirancang untuk efisiensi tinggi dan higienitas maksimal. Terbuat dari material termoplastik berkualitas tinggi, modular belt sangat tahan aus, bebas korosi, dan ideal untuk aplikasi straight, curved (berbelok), maupun incline/decline.</p><h3>Keunggulan Utama</h3><ul><li><strong>100% Bebas Karat &amp; Higienis:</strong> bebas korosi, tidak beracun, dan memenuhi standar food grade (FDA/GMP).</li><li><strong>Perawatan Cepat &amp; Hemat Biaya:</strong> jika terjadi kerusakan, cukup ganti segmen modul yang rusak tanpa mengganti seluruh belt.</li><li><strong>Drainase &amp; Airflow Optimal:</strong> pilihan desain terbuka (open grid) memudahkan aliran udara dan pembilasan/pencucian air.</li><li><strong>Daya Tahan Beban Tinggi:</strong> konstruksi kokoh sanggup menahan beban berat dan tahan terhadap gesekan berulang.</li></ul><h3>Pilihan Material</h3><ul><li><strong>Polypropylene (PP):</strong> tahan bahan kimia, ekonomis, ideal untuk suhu ruangan standar.</li><li><strong>Polyethylene (PE):</strong> sangat baik untuk suhu dingin / pembekuan (freezer application).</li><li><strong>Acetal (POM):</strong> sangat kuat, tahan gesek tinggi, dan cocok untuk beban berat.</li></ul><h3>Tipe Permukaan</h3><ul><li>Flush Grid / Open Grid — berlubang untuk pencucian/pendinginan</li><li>Flat Top — permukaan tertutup rapat &amp; rata</li><li>Rubber Top / Friction Top — anti-selip untuk sudut kemiringan</li><li>Radius / Side-Flexing — untuk jalur conveyor berbelok</li></ul><h3>Aplikasi Utama</h3><ul><li>Industri makanan &amp; minuman (bakery, meat/seafood processing)</li><li>Pengalengan dan lini pembotolan</li><li>Pengemasan farmasi</li><li>Sistem penyortiran logistik</li></ul>',
 NULL,
 '{"Material": "Polypropylene (PP) / Polyethylene (PE) / Acetal (POM)", "Standar": "Food grade (FDA/GMP), bebas korosi", "Tipe Permukaan": "Flush Grid, Flat Top, Friction Top, Radius", "Konfigurasi Jalur": "Straight, curved, incline/decline"}'),

-- ── Timing Belt (1) ──

('timing-belt', 'Timing Belt', 'lainnya', 'timing-belt', 'Timing Belt',
 'Timing belt industri untuk transmisi daya dan penggerak presisi tanpa slip. Bebas perawatan pelumas dan beroperasi halus tanpa bising.',
 '<p>Timing Belt Industri adalah solusi transmisi daya dan sistem penggerak presisi tinggi tanpa slip. Dirancang dengan profil gigi akurat dan kawat penguat (fiberglass/steel cord), produk ini memastikan sinkronisasi mesin yang efisien, tahan lama, dan beroperasi secara halus tanpa bising.</p><h3>Keunggulan Utama</h3><ul><li><strong>Tanpa Slip (Zero Slip):</strong> transmisi daya presisi hingga 99%.</li><li><strong>Bebas Perawatan:</strong> tidak memerlukan oli/gemuk pelumas.</li><li><strong>Tahan Lama:</strong> tahan terhadap gesekan, minyak, dan bahan kimia ringan.</li><li><strong>Operasi Halus:</strong> bebas bising dibandingkan sistem rantai/roda gigi.</li></ul><h3>Profil Gigi yang Tersedia</h3><ul><li>Klasik (Trapezoidal): XL, L, H, XH</li><li>HTD (Bulat): 3M, 5M, 8M, 14M</li><li>Metric: T5, T10, AT5, AT10, S8M</li></ul><h3>Aplikasi Utama</h3><ul><li>Mesin packaging</li><li>Industri makanan &amp; farmasi</li><li>Mesin CNC, tekstil, dan percetakan</li><li>Sistem konveyor otomatis</li></ul>',
 NULL,
 '{"Material": "Karet Neoprene / Polyurethane (PU) Food Grade", "Penguat": "Fiberglass / steel cord", "Profil Gigi": "XL, L, H, XH / 3M, 5M, 8M, 14M / T5, T10, AT5, AT10, S8M", "Bentuk": "Endless (tanpa sambungan) & Open-End (meteran)", "Presisi Transmisi": "Hingga 99%, zero slip"}'),

-- ── V-Belt (1) ──

('v-belt', 'V-Belt', 'lainnya', 'v-belt', 'V-Belt',
 'V-Belt industri untuk transmisi daya dari mesin penggerak ke poros pendukung. Karet sintetis diperkuat polyester cord, tahan shock load dan slip minimal.',
 '<p>V-Belt (Vanbelt) Industri adalah solusi transmisi daya fleksibel dan andal untuk mengalirkan tenaga dari mesin penggerak ke poros pendukung. Terbuat dari formulasi karet sintetis berkualitas tinggi yang diperkuat polyester cord, V-Belt ini memberikan daya cengkeram optimal, tahan terhadap beban kejutan (shock load), serta meminimalkan slip saat operasi tinggi.</p><h3>Keunggulan Utama</h3><ul><li><strong>Daya Cengkeram Tinggi:</strong> desain penampang trapesium memaksimalkan traksi pada pulley.</li><li><strong>Tahan Gesekan &amp; Panas:</strong> tahan terhadap suhu tinggi, kelembapan, serta cipratan oli/minyak.</li><li><strong>Meredam Getaran:</strong> menyerap kejutan beban (shock absorption) sehingga melindungi komponen mesin lainnya.</li><li><strong>Pemasangan Mudah:</strong> ekonomis, efisien, dan cocok untuk berbagai skala transmisi mekanis.</li></ul><h3>Tipe Profil Standar</h3><ul><li><strong>Standard V-Belt (Klasik):</strong> tipe A, B, C, D, E</li><li><strong>High Capacity / Narrow V-Belt:</strong> 3V, 5V, 8V / SPZ, SPA, SPB, SPC</li><li><strong>Cogged / Toothed V-Belt (Bergerigi):</strong> AX, BX, CX, XPZ, XPA, XPB — untuk putaran cepat &amp; fleksibilitas ekstra</li><li><strong>Double / Banded V-Belt:</strong> untuk beban kerja sangat berat dan mencegah belt terlepas</li></ul><h3>Aplikasi Utama</h3><ul><li>Mesin pompa industri dan kompresor udara</li><li>Blower/fan dan mesin pertanian</li><li>Konveyor dan pemecah batu (stone crusher)</li><li>Berbagai mesin manufaktur berat</li></ul>',
 NULL,
 '{"Material": "Karet Sintetis Premium (CR/EPDM)", "Penguat": "Polyester Cord (anti-mulur)", "Profil Klasik": "A, B, C, D, E", "Profil Narrow": "3V, 5V, 8V / SPZ, SPA, SPB, SPC", "Profil Cogged": "AX, BX, CX, XPZ, XPA, XPB", "Tipe Lain": "Double / Banded V-Belt"}'),

-- ── Fastener (1) ──

('fastener-kuku-macan', 'Fastener (Kuku Macan)', 'lainnya', 'fastener', 'Fastener',
 'Fastener penyambung belt (mechanical splicing) dengan gigi pencengkeram kuat. Pemasangan cepat tanpa proses vulcanizing, meminimalkan downtime mesin.',
 '<p>Fastener Conveyor Belt (Kuku Macan) adalah solusi praktis dan efisien untuk penyambungan mekanis (mechanical splicing) pada sabuk berjalan industri. Dirancang dengan konstruksi logam presisi dan gigi pencengkeram yang kuat, fastener ini memastikan sambungan belt tetap kokoh, rapi, dan tahan terhadap tegangan tinggi saat lini produksi beroperasi.</p><h3>Keunggulan Utama</h3><ul><li><strong>Pemasangan Cepat &amp; Mudah:</strong> meminimalkan waktu henti mesin (downtime) tanpa perlu proses vulcanizing yang rumit.</li><li><strong>Daya Cengkeram Maksimal:</strong> desain gigi "kuku macan" menancap kuat pada serat belt tanpa merusak struktur dasar material.</li><li><strong>Tahan Beban Berat &amp; Gesekan:</strong> terbuat dari material logam pilihan yang tahan karat, gesekan, dan gaya tarik tinggi.</li><li><strong>Fleksibel untuk Berbagai Belt:</strong> ideal untuk rubber belt, PVC belt, maupun PU belt dengan berbagai ketebalan.</li></ul><h3>Tipe &amp; Ukuran</h3><ul><li><strong>Staple Fastener / Wire Lace:</strong> cocok untuk belt tipis hingga menengah (aplikasi packaging &amp; industri ringan).</li><li><strong>Bolt Solid Plate / Heavy Duty:</strong> untuk belt tebal dengan beban kerja berat (stone crusher, tambang, semen).</li></ul><h3>Aplikasi Utama</h3><ul><li>Penyambungan conveyor di pabrik semen dan pertambangan</li><li>Stone crusher</li><li>Industri pakan ternak dan pengolahan kayu</li><li>Lini pengemasan manufaktur umum</li></ul>',
 NULL,
 '{"Material": "Galvanized Steel / Carbon Steel", "Opsi Anti-Karat": "Stainless Steel (304 / 316)", "Tipe": "Staple Fastener / Wire Lace, Bolt Solid Plate / Heavy Duty", "Kompatibilitas": "Rubber belt, PVC belt, PU belt"}'),

-- ── Gravity Roll (1) ──

('gravity-roller', 'Gravity Roller', 'lainnya', 'gravity-roll', 'Gravity Roll',
 'Roller conveyor non-motorized untuk logistik, pergudangan, dan distribusi. Konstruksi kuat yang menahan beban berat tanpa memerlukan sumber daya listrik.',
 '<p>Gravity Roller adalah komponen penting dalam sistem conveyor yang berfungsi mempermudah pergerakan barang dari satu titik ke titik lain secara efisien. Produk ini dirancang untuk industri logistik, pergudangan, dan distribusi, mendukung aliran barang yang lancar tanpa memerlukan sumber daya tambahan seperti motor. Dengan konstruksi yang kuat dan material berkualitas, Gravity Roller mampu menahan beban berat serta memberikan kinerja yang stabil dan tahan lama.</p><h3>Manfaat dan Fungsi Produk</h3><ul><li>Mempercepat proses perpindahan barang di jalur conveyor.</li><li>Mengurangi tenaga kerja manual dalam pengangkutan barang.</li><li>Meningkatkan efisiensi operasional gudang dan pusat distribusi.</li><li>Mendukung sistem conveyor non-motorized (tanpa listrik).</li></ul><h3>Keunggulan Dibanding Produk Sejenis</h3><ul><li>Konstruksi kokoh dan awet untuk penggunaan jangka panjang.</li><li>Instalasi mudah dan fleksibel untuk berbagai konfigurasi conveyor.</li><li>Mengurangi risiko kerusakan barang karena pergerakan yang lancar.</li></ul><h3>Rekomendasi Industri atau Penggunaan</h3><ul><li>Pergudangan dan pusat distribusi</li><li>Industri logistik dan manufaktur</li><li>Lini produksi yang memerlukan transportasi barang secara efisien</li></ul>',
 NULL,
 '{"Material": "Baja atau stainless steel sesuai kebutuhan", "Diameter Roller": "Variatif, disesuaikan jenis conveyor", "Panjang Roller": "Standar atau custom sesuai jalur conveyor", "Sistem Bantalan": "Tahan lama, mendukung beban berat", "Penggerak": "Non-motorized (gravitasi)"}'),

-- ── Open Mesh (1) ──

('open-mesh-belt-dryer', 'Open Mesh/Net Dryer', 'lainnya', 'open-mesh', 'Open Mesh',
 'PTFE/Teflon open mesh belt untuk proses pengeringan dan pendinginan. Tahan suhu -70°C hingga +260°C dengan sirkulasi udara maksimal dan sifat anti-lengket.',
 '<p>PTFE / Teflon Open Mesh Belt adalah sabuk berjalan jaring-jaring yang dirancang khusus untuk proses produksi yang melibatkan paparan panas tinggi, pengeringan (drying), atau pendinginan (cooling). Struktur berlubang (mesh) memungkinkan sirkulasi udara atau hawa panas menembus produk secara merata.</p><h3>Keunggulan Utama</h3><ul><li><strong>Tahan Suhu Ekstrem:</strong> beroperasi stabil pada suhu -70°C hingga +260°C.</li><li><strong>Sirkulasi Udara Optimal:</strong> ukuran lubang mesh (seperti 4x4 mm, 10x10 mm) mempercepat proses pengeringan/pendinginan.</li><li><strong>Sifat Anti-Lengket (Non-Stick):</strong> produk tidak mudah menempel pada permukaan belt, sangat mudah dibersihkan.</li><li><strong>Tahan Bahan Kimia &amp; UV/Microwave:</strong> cocok untuk mesin tunnel oven, UV curing, maupun dryer.</li></ul><h3>Aplikasi Utama</h3><ul><li>Mesin pengering makanan/bumbu</li><li>Pengeringan sablon/tekstil (tunnel dryer)</li><li>Mesin shrink wrapping</li><li>Pabrik dengan proses sterilisasi panas</li></ul>',
 NULL,
 '{"Material": "PTFE / Teflon coated fiberglass", "Suhu Kerja": "-70°C s/d +260°C", "Ukuran Mesh": "4x4 mm, 10x10 mm (opsi lain tersedia)", "Sifat Permukaan": "Anti-lengket (non-stick)", "Ketahanan": "Bahan kimia, UV, microwave"}')

-- Baris yang selamat dari langkah 2 di-refresh isinya di sini. `image` sengaja
-- tidak masuk daftar SET: foto yang sudah diunggah ke Cloudinary harus bertahan.
-- Migrasi jadi idempoten — aman dijalankan ulang.
ON CONFLICT (slug) DO UPDATE SET
    name        = EXCLUDED.name,
    "group"     = EXCLUDED."group",
    kategori    = EXCLUDED.kategori,
    category    = EXCLUDED.category,
    description = EXCLUDED.description,
    detail      = EXCLUDED.detail,
    specs       = EXCLUDED.specs,
    updated_at  = now();

-- ── 5. INDUSTRY_PRODUCTS — tautkan ulang ke slug baru ──

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'conveyor-belt-pvc-putih',
    'pu-belt-putih',
    'pu-belt-biru',
    'modular-belt',
    'sidewall-cleated-conveyor-belt'
  ])
) ps
WHERE i.slug = 'makanan-dan-minuman';

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'rubber-belt-polos',
    'rubber-belt-sersan',
    'fastener-kuku-macan',
    'v-belt'
  ])
) ps
WHERE i.slug = 'tambang-dan-semen';

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'conveyor-belt-pvc-hijau',
    'roughtop-belt-hitam-hijau',
    'gravity-roller',
    'pvk-belt',
    'guide-conveyor-belt-profile'
  ])
) ps
WHERE i.slug = 'manufaktur-dan-pergudangan';

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'pu-belt-putih',
    'conveyor-belt-pvc-putih',
    'modular-belt',
    'open-mesh-belt-dryer'
  ])
) ps
WHERE i.slug = 'farmasi-dan-kimia';
