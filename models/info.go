package models

type DataInfo struct {
	NamaS string
	NPSN int
	NomorSS int
	Jalan string
	Kelurahan string
	Kecamatan string
	Kota string
}

var IsiInfo = DataInfo {
	NamaS: "SMK Negeri 4 Kota Tasikmalaya",
	NPSN: 20276063,
	NomorSS: 401327810004,
	Jalan: "Jl. Depok RT 02 RW 05",
	Kelurahan: "Sukamenak",
	Kecamatan: "Purbaratu",
	Kota: "Tasikmalaya",
}