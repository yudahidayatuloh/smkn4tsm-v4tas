package models

type DataJurusan struct {
	ID int
	NamaJ  string
	Desc string
	Foto string
	Warna string
	Detail []string
}


var IsiJurusan = []DataJurusan{
	{
		ID:         1,
		NamaJ:      "PPLG",
		Desc:       "Pengembangan Perangkat Lunak dan Gim",
		Foto:     "https://web-smkn4.vercel.app/_nuxt/pplg.DhbRSnK2.png",
		Warna: "",
		Detail: []string{
			"Jurusan ini mempelajari cara membuat, mengembangkan, dan mengelola aplikasi berbasis web, desktop, dan mobile.",
			"Siswa diajarkan membuat gim interaktif, animasi, serta desain antarmuka yang menarik dan fungsional.",
			"Lulusan dibekali keterampilan logika, pemrograman, dan kerja tim untuk siap berkarier di dunia IT modern.",
		},
	},
	{
		ID:         2,
		NamaJ:       "TJKT",
		Desc:       "Teknik Jaringan Komputer dan Telekomunikasi",
		Foto:     "https://web-smkn4.vercel.app/_nuxt/tjkt.XeO8V4_I.png",
		Warna: "",
		Detail: []string{
			"Jurusan ini mempelajari cara membangun, mengelola, dan memperbaiki jaringan komputer serta sistem komunikasi data.",
			"Siswa diajarkan instalasi jaringan LAN, konfigurasi router, server, dan perangkat jaringan lainnya.",
			"Lulusan dibekali kemampuan troubleshooting, keamanan jaringan, dan teknologi fiber optik untuk siap terjun ke industri IT dan telekomunikasi.",
		},
	},
	{
		ID:         3,
		NamaJ:       "TBSM",
		Desc:       "Teknik dan Bisnis Sepeda Motor",
		Foto:     "https://web-smkn4.vercel.app/_nuxt/tbsm.5YZNeyvV.png",
		Warna: "",
		Detail: []string{
			"Jurusan ini mempelajari cara merawat, memperbaiki, dan memodifikasi sepeda motor.",
			"Siswa dibekali pengetahuan tentang sistem injeksi, kelistrikan, dan mesin berbasis teknologi terkini.",
			"Lulusan memiliki keterampilan teknis dan jiwa bisnis untuk bekerja di bengkel profesional atau membuka usaha sendiri.",
		},
	},
	{
		ID:         4,
		NamaJ:       "DKV",
		Desc:       "Desain Komunikasi Visual",
		Foto:     "https://web-smkn4.vercel.app/_nuxt/dkv.Btdc-HP2.png",
		Warna: "",
		Detail: []string{
			"Jurusan ini mempelajari cara menyampaikan pesan melalui karya visual seperti poster, logo, dan ilustrasi.",
			"Siswa diajarkan menggunakan aplikasi seperti Photoshop, Illustrator, dan CorelDRAW untuk membuat desain profesional.",
			"Lulusan dibekali kreativitas dan kemampuan komunikasi visual untuk berkarier di bidang periklanan, media, dan desain grafis.",
		},
	},
	{
		ID:         5,
		NamaJ:       "TOI",
		Desc:       "Teknik Otomasi Industri",
		Foto:     "https://web-smkn4.vercel.app/_nuxt/toi.BFD6ZBmB.png",
		Warna: "",
		Detail: []string{
			"Jurusan ini mempelajari cara merancang, mengoperasikan, dan memelihara sistem otomatis pada industri modern.",
			"Siswa dibekali pengetahuan tentang PLC, sensor, motor listrik, serta sistem kendali berbasis komputer.",
			"Lulusan mampu bekerja sebagai teknisi, operator, atau perancang sistem otomasi di berbagai bidang industri manufaktur dan produksi.",
		},
	},
}
