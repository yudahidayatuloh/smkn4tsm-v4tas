package models

type DataBerita struct {
	ID int
	Judul string
	Desc []string
	Foto string
}

var IsiBerita = []DataBerita {
	{
		ID:    1,
		Judul: "Kegiatan Hari Guru 2025 Di SMK Negeri 4 Tasikmalaya",
		Desc: []string{
			"Perayaan Hari Guru Nasional 2025 di SMKN 4 Tasikmalaya berlangsung meriah dan penuh kehangatan.",
			"Setelah upacara, acara dilanjutkan dengan pemberian apresiasi kepada guru-guru berprestasi dan pegawai yang telah mengabdi lebih dari sepuluh tahun.",
			"Menutup rangkaian kegiatan, diadakan lomba hiburan antar guru dan siswa, seperti voli ceria dan tebak lagu.",
		},
		Foto: "https://aqinbfwyqxlynlxybhod.supabase.co/storage/v1/object/public/img_berita/berita3.JPG",
	},
	{
		ID:    2,
		Judul: "Kegiatan SPMB SMK Negri 4 Tasikmalaya 2026",
		Desc: []string{
			"Seleksi Penerimaan Murid Baru (SPMB) 2026 di SMKN 4 Tasikmalaya resmi dimulai pada awal Januari dan disambut antusias oleh para calon peserta didik. Pendaftaran dibuka secara daring melalui portal resmi sekolah, dengan mekanisme verifikasi berkas yang telah disederhanakan agar lebih mudah diakses oleh masyarakat. Panitia juga menyediakan layanan bantuan langsung di sekolah untuk mendampingi orang tua maupun siswa yang mengalami kendala teknis saat proses pendaftaran.",
			"Pada tahap seleksi, calon siswa dapat memilih kompetensi keahlian sesuai minat seperti Rekayasa Perangkat Lunak, Teknik Kendaraan Ringan, Desain Komunikasi Visual, dan lainnya. SMKN 4 Tasikmalaya memastikan proses seleksi berlangsung objektif melalui kombinasi penilaian rapor, prestasi akademik/non-akademik, serta tes minat bakat yang dijadwalkan pada pertengahan bulan. Selain itu, sekolah bekerja sama dengan beberapa mitra industri untuk melakukan observasi terhadap potensi siswa sejak tahap awal.",
			"Pihak sekolah berharap kegiatan SPMB tahun ini dapat menjaring siswa-siswi berkualitas yang siap berkembang di dunia pendidikan vokasi. Kepala sekolah menyampaikan bahwa seluruh tahapan seleksi dirancang untuk memberikan pengalaman yang adil, transparan, dan informatif bagi peserta. Informasi lanjutan terkait hasil seleksi serta jadwal daftar ulang akan diumumkan secara resmi melalui website sekolah dan media sosial dalam waktu deka",
		},
		Foto: "https://aqinbfwyqxlynlxybhod.supabase.co/storage/v1/object/public/img_berita/berita2.jpg",
	},
}