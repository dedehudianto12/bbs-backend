-- Seed semua data: produk, artikel, jasa, galeri, industri
-- Categories & admin sudah seed di migration sebelumnya
-- ============================================================

-- ── 1. PRODUK (24 items) ──

INSERT INTO products (slug, name, "group", kategori, category, description, detail, image, specs) VALUES

-- PVC Belt (3)
('pvc-belt-hijau-food-grade', 'PVC Belt Hijau Food Grade', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC belt hijau food grade untuk industri makanan dan pengemasan. Permukaan halus, non-toxic, tahan minyak, dan mudah dibersihkan.',
 '<p>PVC belt hijau food grade untuk industri makanan dan pengemasan. Permukaan halus, non-toxic, tahan minyak, dan mudah dibersihkan.</p>',
 NULL,
 '{"Material": "PVC", "Warna": "Hijau", "Ketebalan": "2–5 mm", "Lebar": "300–1200 mm", "Suhu Kerja": "-10°C s/d 80°C"}'),

('pvc-belt-hitam-general', 'PVC Belt Hitam General Purpose', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC belt hitam serbaguna untuk material handling umum di gudang, pabrik, dan jalur produksi.',
 '<p>PVC belt hitam serbaguna untuk material handling umum di gudang, pabrik, dan jalur produksi.</p>',
 NULL,
 '{"Material": "PVC", "Warna": "Hitam", "Ketebalan": "2–6 mm", "Lebar": "300–1500 mm", "Suhu Kerja": "-10°C s/d 70°C"}'),

('pvc-belt-putih-packaging', 'PVC Belt Putih Packaging', 'belt-conveyor', 'pvc-belt', 'PVC Belt',
 'PVC belt putih khusus untuk industri pengemasan dan farmasi ringan. Tampilan bersih dan profesional.',
 '<p>PVC belt putih khusus untuk industri pengemasan dan farmasi ringan. Tampilan bersih dan profesional.</p>',
 NULL,
 '{"Material": "PVC", "Warna": "Putih", "Ketebalan": "1.5–4 mm", "Lebar": "200–1000 mm", "Suhu Kerja": "-5°C s/d 70°C"}'),

-- PU (3)
('pu-belt-putih-food-grade', 'PU Belt Putih Food Grade', 'belt-conveyor', 'pu', 'PU',
 'PU belt putih food grade untuk industri makanan dan farmasi. Non-toxic, tahan abrasi, dan mudah disanitasi.',
 '<p>PU belt putih food grade untuk industri makanan dan farmasi. Non-toxic, tahan abrasi, dan mudah disanitasi.</p>',
 NULL,
 '{"Material": "PU (Polyurethane)", "Warna": "Putih", "Ketebalan": "1.5–4 mm", "Lebar": "300–1500 mm", "Suhu Kerja": "-20°C s/d 100°C"}'),

('pu-belt-anti-statis', 'PU Belt Anti-Statis', 'belt-conveyor', 'pu', 'PU',
 'PU belt dengan sifat anti-statis untuk industri elektronik dan komponen sensitif terhadap listrik statis.',
 '<p>PU belt dengan sifat anti-statis untuk industri elektronik dan komponen sensitif terhadap listrik statis.</p>',
 NULL,
 '{"Material": "PU Anti-Statis", "Warna": "Abu-abu", "Ketebalan": "2–5 mm", "Lebar": "300–1200 mm", "Resistansi": "10⁶–10⁸ Ω"}'),

('pu-belt-tahan-minyak', 'PU Belt Tahan Minyak', 'belt-conveyor', 'pu', 'PU',
 'PU belt tahan minyak untuk industri pengolahan makanan berminyak dan manufaktur dengan paparan oli.',
 '<p>PU belt tahan minyak untuk industri pengolahan makanan berminyak dan manufaktur dengan paparan oli.</p>',
 NULL,
 '{"Material": "PU Oil-Resistant", "Warna": "Krem", "Ketebalan": "2–5 mm", "Lebar": "300–1500 mm", "Suhu Kerja": "-15°C s/d 90°C"}'),

-- Flat Belt (3)
('flat-belt-nylon', 'Flat Belt Nylon', 'belt-conveyor', 'flat-belt', 'Flat Belt',
 'Flat belt nylon dengan kekuatan tarik tinggi untuk conveyor jarak jauh dan beban sedang di sektor manufaktur.',
 '<p>Flat belt nylon dengan kekuatan tarik tinggi untuk conveyor jarak jauh dan beban sedang di sektor manufaktur.</p>',
 NULL,
 '{"Material": "Nylon Core", "Warna": "Hijau / Putih", "Ketebalan": "1–4 mm", "Lebar": "20–500 mm", "Kekuatan Tarik": "s/d 50 N/mm"}'),

('flat-belt-polyester', 'Flat Belt Polyester', 'belt-conveyor', 'flat-belt', 'Flat Belt',
 'Flat belt polyester dengan elongasi rendah, ideal untuk aplikasi yang memerlukan stabilitas dimensi tinggi.',
 '<p>Flat belt polyester dengan elongasi rendah, ideal untuk aplikasi yang memerlukan stabilitas dimensi tinggi.</p>',
 NULL,
 '{"Material": "Polyester Fabric", "Warna": "Putih", "Ketebalan": "0.8–3 mm", "Lebar": "20–400 mm", "Elongasi": "< 1%"}'),

('flat-belt-rubber-cover', 'Flat Belt Rubber Cover', 'belt-conveyor', 'flat-belt', 'Flat Belt',
 'Flat belt dengan lapisan rubber cover untuk grip ekstra pada permukaan licin atau miring.',
 '<p>Flat belt dengan lapisan rubber cover untuk grip ekstra pada permukaan licin atau miring.</p>',
 NULL,
 '{"Material": "Rubber + Fabric", "Warna": "Hitam", "Ketebalan": "2–6 mm", "Lebar": "30–500 mm", "Koefisien Gesek": "Tinggi"}'),

-- Rubber Belt (3)
('rubber-belt-nn', 'Rubber Belt NN (Nylon-Nylon)', 'belt-conveyor', 'rubber-belt', 'Rubber Belt',
 'Rubber conveyor belt dengan carcass NN untuk aplikasi tambang, pelabuhan, dan industri berat dengan beban tinggi.',
 '<p>Rubber conveyor belt dengan carcass NN untuk aplikasi tambang, pelabuhan, dan industri berat dengan beban tinggi.</p>',
 NULL,
 '{"Material": "Rubber + NN Carcass", "Warna": "Hitam", "Ketebalan": "8–25 mm", "Lebar": "500–2000 mm", "Kekuatan": "s/d 400 N/mm"}'),

('rubber-belt-ep', 'Rubber Belt EP (Polyester-Nylon)', 'belt-conveyor', 'rubber-belt', 'Rubber Belt',
 'Rubber belt EP dengan elongasi rendah dan stabilitas tinggi untuk conveyor jarak menengah hingga jauh.',
 '<p>Rubber belt EP dengan elongasi rendah dan stabilitas tinggi untuk conveyor jarak menengah hingga jauh.</p>',
 NULL,
 '{"Material": "Rubber + EP Carcass", "Warna": "Hitam", "Ketebalan": "6–20 mm", "Lebar": "500–1800 mm", "Elongasi": "< 2%"}'),

('rubber-belt-heat-resistant', 'Rubber Belt Tahan Panas', 'belt-conveyor', 'rubber-belt', 'Rubber Belt',
 'Rubber belt tahan panas untuk industri semen, baja, dan pengolahan material bersuhu tinggi hingga 200°C.',
 '<p>Rubber belt tahan panas untuk industri semen, baja, dan pengolahan material bersuhu tinggi hingga 200°C.</p>',
 NULL,
 '{"Material": "Heat-Resistant Rubber", "Warna": "Hitam", "Ketebalan": "8–25 mm", "Lebar": "500–2000 mm", "Suhu Maks": "s/d 200°C"}'),

-- Timing Belt (3)
('timing-belt-at10', 'Timing Belt AT10', 'lainnya', 'timing-belt', 'Timing Belt',
 'Timing belt AT10 dengan pitch 10 mm untuk aplikasi linear drive dan positioning presisi pada mesin industri.',
 '<p>Timing belt AT10 dengan pitch 10 mm untuk aplikasi linear drive dan positioning presisi pada mesin industri.</p>',
 NULL,
 '{"Pitch": "10 mm", "Material": "PU + Steel Cord", "Lebar": "16–100 mm", "Suhu Kerja": "-20°C s/d 80°C"}'),

('timing-belt-t5', 'Timing Belt T5', 'lainnya', 'timing-belt', 'Timing Belt',
 'Timing belt T5 pitch 5 mm untuk aplikasi ringan seperti printer, mesin tekstil, dan otomasi kantor.',
 '<p>Timing belt T5 pitch 5 mm untuk aplikasi ringan seperti printer, mesin tekstil, dan otomasi kantor.</p>',
 NULL,
 '{"Pitch": "5 mm", "Material": "PU + Kevlar Cord", "Lebar": "10–50 mm", "Suhu Kerja": "-15°C s/d 70°C"}'),

('timing-belt-htd', 'Timing Belt HTD 8M', 'lainnya', 'timing-belt', 'Timing Belt',
 'Timing belt HTD 8M high torque drive untuk aplikasi beban berat dengan profil gigi curvilinear.',
 '<p>Timing belt HTD 8M high torque drive untuk aplikasi beban berat dengan profil gigi curvilinear.</p>',
 NULL,
 '{"Pitch": "8 mm", "Material": "Rubber + Fiberglass", "Lebar": "20–85 mm", "Profil": "Curvilinear HTD"}'),

-- Fastener (3)
('alligator-lacing-steel', 'Alligator Lacing Baja', 'lainnya', 'fastener', 'Fastener',
 'Alligator lacing baja untuk sambungan belt conveyor mekanis. Pemasangan mudah tanpa perlu vulkanisir panas.',
 '<p>Alligator lacing baja untuk sambungan belt conveyor mekanis. Pemasangan mudah tanpa perlu vulkanisir panas.</p>',
 NULL,
 '{"Material": "Baja Galvanis", "Ketebalan": "1.5–3 mm", "Lebar Belt": "300–1500 mm", "Tipe": "Alligator RS"}'),

('clipper-wire-hook', 'Clipper Wire Hook Fastener', 'lainnya', 'fastener', 'Fastener',
 'Clipper wire hook fastener untuk sambungan cepat belt conveyor ringan. Ideal untuk PVC dan PU belt.',
 '<p>Clipper wire hook fastener untuk sambungan cepat belt conveyor ringan. Ideal untuk PVC dan PU belt.</p>',
 NULL,
 '{"Material": "Stainless Steel", "Diameter": "0.8–1.5 mm", "Lebar Belt": "200–800 mm", "Tipe": "Clipper Lacing"}'),

('bolt-solid-plate-fastener', 'Bolt Solid Plate Fastener', 'lainnya', 'fastener', 'Fastener',
 'Bolt solid plate fastener untuk sambungan rubber belt tugas berat di tambang dan pelabuhan.',
 '<p>Bolt solid plate fastener untuk sambungan rubber belt tugas berat di tambang dan pelabuhan.</p>',
 NULL,
 '{"Material": "Baja Manganese", "Ketebalan": "4–8 mm", "Lebar Belt": "600–2000 mm", "Kekuatan": "s/d 2000 N/mm"}'),

-- Cleat (3)
('cleat-pvc-tipe-c', 'Cleat PVC Tipe C', 'lainnya', 'cleat', 'Cleat',
 'Cleat PVC tipe C untuk belt conveyor miring. Mencegah material tergelincir saat incline conveying.',
 '<p>Cleat PVC tipe C untuk belt conveyor miring. Mencegah material tergelincir saat incline conveying.</p>',
 NULL,
 '{"Material": "PVC", "Tinggi": "10–50 mm", "Tipe": "C (Trapezoid)", "Warna": "Hijau"}'),

('cleat-pu-tipe-a', 'Cleat PU Tipe A', 'lainnya', 'cleat', 'Cleat',
 'Cleat PU tipe A food grade untuk conveyor makanan. Profil rendah, mudah dibersihkan, non-toxic.',
 '<p>Cleat PU tipe A food grade untuk conveyor makanan. Profil rendah, mudah dibersihkan, non-toxic.</p>',
 NULL,
 '{"Material": "PU Food Grade", "Tinggi": "8–25 mm", "Tipe": "A (Rectangle)", "Warna": "Putih"}'),

('sidewall-cleat-rubber', 'Sidewall Cleat Rubber', 'lainnya', 'cleat', 'Cleat',
 'Sidewall cleat rubber untuk conveyor sudut curam. Mencegah spillage material pada kemiringan tinggi.',
 '<p>Sidewall cleat rubber untuk conveyor sudut curam. Mencegah spillage material pada kemiringan tinggi.</p>',
 NULL,
 '{"Material": "Rubber", "Tinggi": "40–120 mm", "Tipe": "Sidewall + Cleat", "Warna": "Hitam"}'),

-- Gravity Roll (3)
('roller-conveyor-baja', 'Roller Conveyor Baja', 'lainnya', 'gravity-roll', 'Gravity Roll',
 'Roller conveyor baja untuk sistem material handling di gudang, pabrik, dan jalur distribusi.',
 '<p>Roller conveyor baja untuk sistem material handling di gudang, pabrik, dan jalur distribusi.</p>',
 NULL,
 '{"Material": "Baja Karbon", "Diameter": "50–150 mm", "Panjang": "300–1500 mm", "Kapasitas Beban": "s/d 500 kg/roller"}'),

('roller-conveyor-stainless', 'Roller Conveyor Stainless Steel', 'lainnya', 'gravity-roll', 'Gravity Roll',
 'Roller stainless steel untuk industri makanan, minuman, dan farmasi. Anti karat dan mudah disanitasi.',
 '<p>Roller stainless steel untuk industri makanan, minuman, dan farmasi. Anti karat dan mudah disanitasi.</p>',
 NULL,
 '{"Material": "Stainless Steel 304", "Diameter": "50–100 mm", "Panjang": "300–1200 mm", "Kapasitas Beban": "s/d 300 kg/roller"}'),

('gravity-roll-pvc', 'Gravity Roll PVC', 'lainnya', 'gravity-roll', 'Gravity Roll',
 'Gravity roll PVC ringan untuk conveyor manual dan aplikasi pergudangan ringan. Ekonomis dan tahan korosi.',
 '<p>Gravity roll PVC ringan untuk conveyor manual dan aplikasi pergudangan ringan. Ekonomis dan tahan korosi.</p>',
 NULL,
 '{"Material": "PVC", "Diameter": "40–76 mm", "Panjang": "200–1000 mm", "Kapasitas Beban": "s/d 50 kg/roller"}');

-- ── 2. ARTIKEL (9 items) ──

INSERT INTO articles (slug, title, excerpt, content, image, tag, published_at, author) VALUES

('tips-memilih-belt-conveyor-food-grade', 'Tips Memilih Belt Conveyor Food Grade untuk Industri Makanan',
 'Pelajari kriteria penting dalam memilih belt conveyor food-grade yang sesuai standar keamanan pangan BPOM dan HACCP.',
 '<p>Industri makanan memiliki persyaratan ketat untuk peralatan produksi, termasuk sistem conveyor. Belt conveyor yang digunakan harus memenuhi standar food grade — bebas dari kontaminasi, mudah dibersihkan, dan tahan terhadap prosedur sanitasi rutin.</p><p>Material yang paling umum digunakan adalah PU (Polyurethane) dan PVC food grade. PU lebih unggul dalam ketahanan abrasi dan suhu, sementara PVC lebih ekonomis untuk aplikasi ringan. Pastikan belt yang Anda pilih memiliki sertifikasi FDA atau EU 1935/2004 untuk kontak langsung dengan makanan.</p><p>Selain material, perhatikan juga desain conveyor: hindari sudut tajam, gunakan sealed bearing, dan pastikan sistem mudah dibongkar untuk pembersihan menyeluruh. Konsultasikan kebutuhan Anda dengan tim kami untuk rekomendasi yang tepat.</p>',
 NULL, 'Tips', '2026-06-15', 'Tim BBS Conveyor'),

('panduan-perawatan-belt-conveyor-tambang', 'Panduan Perawatan Belt Conveyor di Lingkungan Tambang',
 'Lingkungan tambang yang keras menuntut perawatan rutin agar belt conveyor tetap andal, awet, dan minim downtime.',
 '<p>Tambang adalah salah satu lingkungan paling keras untuk sistem conveyor. Debu, material abrasif, beban berat, dan kondisi cuaca ekstrem membuat belt conveyor memerlukan perawatan berkala yang ketat.</p><p>Inspeksi harian meliputi pengecekan tracking belt, kondisi permukaan belt (sobek, aus), dan suara tidak normal dari bearing atau pulley. Inspeksi mingguan mencakup pengecekan tension belt, kondisi idler, dan kebersihan area sekitar conveyor.</p><p>Untuk perawatan preventif jangka panjang, kami merekomendasikan penggantian idler yang aus, pelumasan bearing secara berkala, dan pengecekan alignment struktur rangka. Tim BBS Conveyor menyediakan layanan inspeksi dan perawatan onsite di seluruh Indonesia.</p>',
 NULL, 'Tips', '2026-05-20', 'Tim BBS Conveyor'),

('cara-menghitung-kapasitas-belt-conveyor', 'Cara Menghitung Kapasitas Belt Conveyor untuk Proyek Anda',
 'Panduan praktis menghitung kapasitas belt conveyor berdasarkan lebar belt, kecepatan, dan karakteristik material.',
 '<p>Menghitung kapasitas belt conveyor yang tepat sangat penting untuk memastikan sistem material handling Anda berjalan efisien. Kapasitas dihitung berdasarkan tiga faktor utama: lebar belt, kecepatan belt, dan densitas material yang diangkut.</p><p>Rumus dasar: Kapasitas (ton/jam) = Lebar Belt (m) × Kecepatan (m/s) × Densitas Material (ton/m³) × 3600 × Faktor Pengisian. Faktor pengisian bervariasi antara 0,5 hingga 0,9 tergantung sudut kemiringan dan karakteristik material.</p><p>Untuk belt conveyor dengan cleat atau sidewall, perhitungan menjadi lebih kompleks dan memerlukan konsultasi dengan engineer berpengalaman. Hubungi tim BBS Conveyor untuk mendapatkan kalkulasi yang akurat sesuai spesifikasi proyek Anda.</p>',
 NULL, 'Panduan', '2026-04-10', 'Tim BBS Conveyor'),

('supplier-belt-conveyor-semarang', 'Supplier Belt Conveyor Terpercaya di Semarang',
 'BBS Conveyor melayani kebutuhan belt conveyor, roller, dan komponen industri di Semarang dan sekitarnya.',
 '<p>Semarang sebagai pusat industri di Jawa Tengah memiliki kebutuhan tinggi akan sistem material handling yang andal. BBS Conveyor hadir melayani pabrik, gudang, dan fasilitas produksi di Semarang dengan produk belt conveyor berkualitas.</p><p>Kami menyediakan berbagai jenis belt conveyor: PVC belt, PU belt, rubber belt, dan flat belt dengan dukungan teknis lengkap dari konsultasi hingga instalasi. Layanan onsite joint dan reparasi juga tersedia untuk pelanggan di area Semarang dan sekitarnya — termasuk Kendal, Demak, dan Ungaran.</p><p>Dengan tim teknisi yang siap mobilisasi cepat, kami memastikan downtime minimal untuk operasional Anda. Hubungi kami untuk konsultasi dan penawaran khusus wilayah Semarang.</p>',
 NULL, 'Proyek', '2026-03-25', 'Tim BBS Conveyor'),

('supplier-belt-conveyor-surabaya', 'Supplier Belt Conveyor dan Jasa Instalasi di Surabaya',
 'Layanan lengkap belt conveyor untuk industri di Surabaya dan kawasan industri Jawa Timur.',
 '<p>Surabaya dan kawasan industri sekitarnya (Gresik, Sidoarjo, Mojokerto) merupakan salah satu pusat manufaktur terbesar di Indonesia. BBS Conveyor menyediakan solusi belt conveyor lengkap untuk berbagai sektor industri di wilayah ini.</p><p>Portofolio proyek kami di Surabaya mencakup instalasi conveyor untuk pabrik makanan, gudang logistik, dan fasilitas pengolahan material. Kami juga melayani kebutuhan rubber belt untuk industri berat di sekitar Pelabuhan Tanjung Perak.</p><p>Tim teknisi kami siap melakukan survey lokasi, memberikan rekomendasi teknis, dan melaksanakan instalasi dengan standar kualitas tinggi. Dapatkan konsultasi gratis untuk proyek conveyor Anda di Surabaya.</p>',
 NULL, 'Proyek', '2026-03-10', 'Tim BBS Conveyor'),

('jasa-onsite-joint-cikarang', 'Jasa Onsite Joint Belt Conveyor di Cikarang dan Kawasan Industri Bekasi',
 'Layanan penyambungan belt conveyor di lokasi untuk pabrik di Cikarang, Jababeka, EJIP, dan MM2100.',
 '<p>Kawasan industri Cikarang dan Bekasi menampung ribuan pabrik manufaktur yang mengandalkan belt conveyor untuk proses produksi harian. BBS Conveyor menyediakan jasa onsite joint untuk instalasi belt baru maupun penggantian belt yang aus.</p><p>Kami melayani hot splicing untuk rubber belt di aplikasi berat serta cold splicing untuk PVC dan PU belt di industri makanan dan farmasi. Setiap pekerjaan dilakukan oleh teknisi bersertifikat dengan peralatan lengkap.</p><p>Keunggulan layanan kami: respon cepat dalam 24 jam, garansi hasil sambungan, dan dukungan teknis pasca instalasi. Area layanan mencakup Cikarang, Jababeka, EJIP, MM2100, dan seluruh kawasan industri Bekasi.</p>',
 NULL, 'Proyek', '2026-02-18', 'Tim BBS Conveyor'),

('jenis-cleat-belt-conveyor', 'Mengenal 3 Jenis Cleat untuk Belt Conveyor dan Fungsinya',
 'Cleat tipe C, T, dan L memiliki fungsi berbeda — pilih yang tepat sesuai material dan sudut kemiringan conveyor Anda.',
 '<p>Cleat adalah sekat yang dipasang pada permukaan belt conveyor untuk mencegah material tergelincir saat conveyor beroperasi pada kemiringan. Pemilihan tipe cleat yang tepat sangat mempengaruhi efisiensi material handling Anda.</p><p><strong>Cleat Tipe C (Trapezoid):</strong> Paling umum digunakan, cocok untuk material curah seperti pasir, batu bara, dan biji-bijian. Profil trapesium memberikan kapasitas angkut optimal pada kemiringan 15°–35°.</p><p><strong>Cleat Tipe T:</strong> Profil tinggi dan tegak, ideal untuk material dengan ukuran bervariasi dan aplikasi yang membutuhkan kapasitas angkut maksimal. Cocok untuk kemiringan curam hingga 45°.</p><p><strong>Cleat Tipe L:</strong> Profil rendah dengan sudut miring, dirancang untuk produk kemasan atau material yang tidak boleh rusak. Biasa digunakan di industri makanan dan farmasi.</p>',
 NULL, 'Panduan', '2026-01-30', 'Tim BBS Conveyor'),

('supplier-belt-conveyor-medan', 'Supplier Belt Conveyor di Medan — Melayani Industri Sumatera',
 'Distribusi belt conveyor dan komponen industri untuk pabrik dan tambang di Medan serta seluruh Sumatera Utara.',
 '<p>Medan sebagai gerbang industri Sumatera memiliki permintaan tinggi untuk peralatan material handling, termasuk belt conveyor. BBS Conveyor melayani pengiriman dan instalasi belt conveyor ke Medan dan kota-kota industri di Sumatera Utara.</p><p>Kami menyediakan rubber belt untuk industri kelapa sawit dan tambang, PVC/PU belt untuk pabrik pengolahan makanan, serta komponen conveyor seperti roller, pulley, dan fastener. Pengiriman didukung jaringan logistik yang menjangkau Medan, Binjai, Deli Serdang, dan sekitarnya.</p><p>Untuk proyek berskala besar, tim teknisi kami siap mobilisasi ke lokasi untuk instalasi dan commissioning. Hubungi sales kami untuk informasi stok, harga, dan estimasi pengiriman ke Medan.</p>',
 NULL, 'Proyek', '2026-01-12', 'Tim BBS Conveyor'),

('tanda-belt-conveyor-harus-diganti', '5 Tanda Belt Conveyor Anda Harus Segera Diganti',
 'Kenali tanda-tanda kerusakan belt conveyor sebelum menyebabkan downtime fatal pada produksi Anda.',
 '<p>Belt conveyor yang aus atau rusak dapat menyebabkan downtime produksi yang merugikan. Mengenali tanda-tanda awal kerusakan membantu Anda menjadwalkan penggantian sebelum terjadi kegagalan total.</p><p><strong>1. Permukaan belt retak atau mengelupas:</strong> Indikasi material belt sudah mencapai batas usia pakai, terutama jika terpapar suhu tinggi atau bahan kimia secara terus-menerus.</p><p><strong>2. Belt sering melenceng (mis-tracking):</strong> Bisa disebabkan oleh belt yang sudah tidak rata atau struktur rangka yang bergeser. Jika adjustment tidak membantu, saatnya ganti belt.</p><p><strong>3. Sambungan mulai terbuka:</strong> Baik sambungan mechanical fastener maupun spliced joint yang mulai longgar adalah tanda belt harus segera diperbaiki atau diganti.</p><p><strong>4. Bunyi tidak normal:</strong> Suara berdecit atau berderak bisa menandakan masalah pada bearing, pulley, atau permukaan belt yang sudah tidak rata.</p><p><strong>5. Penurunan kapasitas angkut:</strong> Jika conveyor Anda tidak lagi mampu mengangkut beban sesuai spesifikasi awal, kemungkinan belt sudah mengalami degradasi struktural.</p>',
 NULL, 'Tips', '2025-12-05', 'Tim BBS Conveyor');

-- ── 3. JASA (4 items) ──

INSERT INTO services (slug, name, short_description, full_description, images) VALUES

('incline-cleated-belt', 'Incline Cleated Belt',
 'Pembuatan belt conveyor dengan cleat untuk jalur produksi miring, mencegah material tergelincir pada sudut kemiringan tinggi.',
 '<p>Layanan Incline Cleated Belt kami menyediakan solusi conveyor untuk jalur produksi dengan kemiringan curam. Kami merancang dan memasang cleat (sekat) pada permukaan belt conveyor sesuai kebutuhan material yang diangkut — baik material curah seperti pasir, batu bara, biji-bijian, maupun produk kemasan.</p><p>Cleat tersedia dalam berbagai tipe (T, C, L) dan material (PVC, PU, Rubber) yang disesuaikan dengan karakteristik material dan sudut kemiringan conveyor Anda. Tim kami akan melakukan assessment di lokasi untuk menentukan spesifikasi cleat yang optimal, memastikan material tidak longsor dan aliran produksi tetap lancar.</p>',
 '{}'),

('profile-guide-belt', 'Profile Guide Belt',
 'Pemasangan profile guide untuk mencegah belt conveyor melenceng dari jalur, ideal untuk beban tidak seimbang.',
 '<p>Profile guide adalah komponen yang dipasang di bagian bawah belt conveyor untuk menjaga belt tetap berada pada jalurnya (tracking). Layanan ini sangat penting untuk conveyor dengan beban tidak seimbang, jalur panjang, atau kondisi operasi yang menyebabkan belt mudah bergeser.</p><p>Kami menyediakan berbagai jenis profile guide — V-guide, rectangular guide, dan custom profile — yang disesuaikan dengan pulley dan sistem tracking conveyor Anda. Pemasangan dilakukan dengan teknik bonding presisi tinggi untuk memastikan guide menempel sempurna dan tidak lepas selama operasi.</p><p>Layanan mencakup: inspeksi kondisi belt dan tracking, rekomendasi tipe guide yang sesuai, pemasangan di workshop atau di lokasi (onsite), dan pengujian tracking setelah pemasangan.</p>',
 '{}'),

('jasa-reparasi', 'Jasa Reparasi',
 'Perbaikan belt conveyor rusak — penyambungan sobek, penggantian bagian aus, dan perbaikan struktur rangka.',
 '<p>Layanan reparasi belt conveyor kami mencakup perbaikan menyeluruh untuk mengembalikan performa conveyor Anda tanpa harus mengganti seluruh sistem. Tim teknisi kami berpengalaman menangani berbagai jenis kerusakan pada belt conveyor industri.</p><p>Layanan yang kami tawarkan meliputi:<ul><li>Penyambungan belt sobek atau putus (hot splicing & cold splicing)</li><li>Penggantian bagian belt yang aus atau terkikis</li><li>Perbaikan atau penggantian cleat dan sidewall</li><li>Perbaikan mechanical fastener yang longgar</li><li>Inspeksi dan perbaikan struktur rangka conveyor</li><li>Penggantian pulley, bearing, dan komponen pendukung</li></ul></p><p>Kami siap datang ke lokasi Anda di seluruh Indonesia untuk inspeksi dan perbaikan di tempat, meminimalkan downtime operasional Anda.</p>',
 '{}'),

('jasa-onsite-joint', 'Jasa Onsite Joint',
 'Penyambungan belt conveyor di lokasi proyek dengan teknik hot/cold splicing untuk hasil presisi dan tahan lama.',
 '<p>Jasa penyambungan belt conveyor di lokasi (onsite joint) adalah layanan utama kami untuk proyek-proyek yang membutuhkan instalasi belt di tempat. Kami melayani penyambungan belt conveyor baru maupun penggantian belt lama di lokasi pelanggan.</p><p>Dua metode penyambungan yang kami tawarkan:</p><p><strong>1. Hot Splicing</strong> — Proses vulkanisasi panas menggunakan mesin press khusus. Menghasilkan sambungan yang sangat kuat dan tahan lama, ideal untuk aplikasi beban berat seperti tambang dan pelabuhan. Kekuatan sambungan mencapai 90% dari kekuatan asli belt.</p><p><strong>2. Cold Splicing</strong> — Proses penyambungan menggunakan adhesive khusus tanpa pemanasan. Lebih cepat pengerjaannya, cocok untuk belt PVC/PU dan aplikasi dengan downtime terbatas.</p><p>Tim kami dilengkapi peralatan splicing lengkap dan siap mobilisasi ke seluruh Indonesia. Setiap pekerjaan diakhiri dengan pengujian tracking dan tension untuk memastikan belt beroperasi optimal.</p>',
 '{}');

-- ── 4. GALERI (8 items) ──

INSERT INTO galleries (image, caption, location) VALUES
('', 'Instalasi Belt Conveyor PVC — Pabrik Makanan, Semarang', 'Semarang'),
('', 'Onsite Joint Rubber Belt — Pabrik Otomotif, Cikarang', 'Cikarang'),
('', 'Setup Roller Conveyor Baja — Gudang Logistik, Surabaya', 'Surabaya'),
('', 'Pemasangan Cleat PVC — Pabrik Pengolahan, Jakarta Timur', 'Jakarta'),
('', 'Instalasi Rubber Belt NN — Tambang Batu Bara, Medan', 'Medan'),
('', 'Pemasangan Profile Guide — Pabrik Farmasi, Jakarta Pusat', 'Jakarta'),
('', 'Penggantian Flat Belt Polyester — Pabrik Tekstil, Bekasi', 'Bekasi'),
('', 'Jasa Reparasi Belt Conveyor — Pabrik Semen, Semarang', 'Semarang');

-- ── 5. INDUSTRI (4 items) ──

INSERT INTO industries (slug, name, description, image) VALUES
('makanan-dan-minuman', 'Makanan & Minuman',
 'Sistem conveyor food-grade untuk jalur produksi, pengemasan, dan distribusi produk makanan dan minuman.',
 NULL),
('tambang-dan-semen', 'Tambang & Semen',
 'Belt conveyor heavy-duty untuk material curah seperti batu bara, pasir, batu kapur, dan semen.',
 NULL),
('manufaktur-dan-pergudangan', 'Manufaktur & Pergudangan',
 'Sistem material handling terintegrasi untuk pabrik dan pusat distribusi, dari roller hingga belt conveyor.',
 NULL),
('farmasi-dan-kimia', 'Farmasi & Kimia',
 'Peralatan conveyor anti korosi dan mudah disanitasi untuk produksi obat dan bahan kimia.',
 NULL);

-- ── 6. INDUSTRY_PRODUCTS (many-to-many) ──

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'pvc-belt-hijau-food-grade',
    'pu-belt-putih-food-grade',
    'roller-conveyor-stainless'
  ])
) ps
WHERE i.slug = 'makanan-dan-minuman';

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'rubber-belt-nn',
    'rubber-belt-ep'
  ])
) ps
WHERE i.slug = 'tambang-dan-semen';

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'roller-conveyor-baja',
    'roller-conveyor-stainless',
    'pvc-belt-hitam-general'
  ])
) ps
WHERE i.slug = 'manufaktur-dan-pergudangan';

INSERT INTO industry_products (industry_id, product_slug)
SELECT i.id, ps.slug
FROM industries i
CROSS JOIN LATERAL (
  SELECT unnest AS slug FROM unnest(ARRAY[
    'pu-belt-putih-food-grade',
    'roller-conveyor-stainless'
  ])
) ps
WHERE i.slug = 'farmasi-dan-kimia';

-- ── 7. CATEGORIES — already seeded in migration 000002, skip

