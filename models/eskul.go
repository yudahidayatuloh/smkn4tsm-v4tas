package models

type DataEskul struct {
    ID int
	NamaE string
	Desc string
	Foto string
}

var IsiEskul = []DataEskul {
	{
        ID: 1,
        NamaE: "Sepak Bola",
        Desc:  "Permainan yang menggunakan kaki",
        Foto:  "https://d1vbn70lmn1nqe.cloudfront.net/prod/wp-content/uploads/2022/11/03042609/Wajib-Tahu-Ini-X-Teknik-Dasar-Sepak-Bola.jpg.webp",
    },
    {
        ID: 2,
        NamaE: "Basket",
        Desc:  "Olahraga yang dimainkan dengan tangan dan bola",
        Foto:  "https://froyonion.sgp1.cdn.digitaloceanspaces.com/images/blogdetail/dcc9bb715c36ec65cb5c7e25dad2a0371db09b83.jpg",
    },
    {
        ID: 3,
        NamaE: "Pencak Silat",
        Desc:  "Seni bela diri tradisional Indonesia",
        Foto:  "https://www.riauonline.co.id/foto/bank/images2/Ilustrasi-silat.jpg",
    },
    {
        ID: 4,
        NamaE: "Pencak Silat",
        Desc:  "Seni bela diri tradisional Indonesia",
        Foto:  "https://www.riauonline.co.id/foto/bank/images2/Ilustrasi-silat.jpg",
    },
}